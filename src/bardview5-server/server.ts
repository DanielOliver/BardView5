import express from "express"
import helmet from "helmet"
import expressSession from 'express-session'
import {buildSchema} from 'graphql'
import {graphqlHTTP} from 'express-graphql'
import {rootValue, schema} from "./schema"
import passport from "passport"
// import OAuth2Strategy from 'passport-oauth2'
import morgan from 'morgan'
import jwt from 'jsonwebtoken'
import {IProfile, OIDCStrategy, VerifyCallback} from 'passport-azure-ad'
import bodyParser from 'body-parser'
import cookieParser from 'cookie-parser'

// https://github.com/AzureAD/microsoft-authentication-library-for-js/tree/dev/maintenance/passport-azure-ad
const config = {
    creds: {
        responseType: 'code',

        // Required
        responseMode: 'form_post',

        // Required, the reply URL registered in AAD for your app
        redirectUrl: 'http://localhost:4000/auth/openid/return',

        // Required if we use http for redirectUrl
        allowHttpForRedirectUrl: true,

        // Required to set to false if you don't want to validate issuer
        validateIssuer: false,

        // Recommended to set to true. By default we save state in express session, if this option is set to true, then
        // we encrypt state and save it in cookie instead. This option together with { session: false } allows your app
        // to be completely express session free.
        useCookieInsteadOfSession: true,

        // Required if `useCookieInsteadOfSession` is set to true. You can provide multiple set of key/iv pairs for key
        // rollover purpose. We always use the first set of key/iv pair to encrypt cookie, but we will try every set of
        // key/iv pair to decrypt cookie. Key can be any string of length 32, and iv can be any string of length 12.
        cookieEncryptionKeys: [
            {'key': '12345678901234567890123456789012', 'iv': '123456789012'},
            {'key': 'abcdefghijklmnopqrstuvwxyzabcdef', 'iv': 'abcdefghijkl'}
        ],

        scope: ['profile', 'openid', 'offline_access', 'https://graph.microsoft.com/mail.read'],
    }
}


const tokenToProfile = async (idToken: string) => {

    // !!<-- Make sure you validate the token's signature -->!!
    // And make sure you handle errors. I simplified the code for the blog post

    const profile = jwt.decode(idToken);

    return profile
};

function serve(isProduction: boolean) {
    const app = express()
    app.use(helmet({contentSecurityPolicy: (isProduction) ? undefined : false}))
    const expiryDate = new Date(Date.now() + 60 * 60 * 1000) // 1 hour
    app.use(morgan('combined'))
    app.use(cookieParser());
    app.use(expressSession({secret: 'keyboard cat', resave: true, saveUninitialized: false}));
    app.use(bodyParser.urlencoded({extended: true}));
    app.use(passport.initialize());
    app.use(passport.session());

    passport.use(new OIDCStrategy({
            identityMetadata: process.env.IDENTITY_METADATA || '',
            clientID: process.env.OAUTH_CLIENT_ID || '',
            clientSecret: process.env.OAUTH_CLIENT_SECRET || '',
            responseType: 'code',
            responseMode: 'form_post',
            redirectUrl: process.env.OAUTH_CALLBACK_URL || '',
            allowHttpForRedirectUrl: config.creds.allowHttpForRedirectUrl,
            validateIssuer: config.creds.validateIssuer,
            passReqToCallback: false,
            scope: config.creds.scope,
            useCookieInsteadOfSession: config.creds.useCookieInsteadOfSession,
            cookieEncryptionKeys: config.creds.cookieEncryptionKeys,
            loggingLevel: 'info'
        },
        function (iss: string, sub: string, profile: IProfile, access_token: string, refresh_token: string, done: VerifyCallback) {
            if (!profile.oid) {
                return done(new Error("No oid found"), null);
            }

            return done(null, profile);
        }
    ));

    passport.serializeUser(function (user, done) {
        done(null, user);
    });
    passport.deserializeUser(function (user, done) {
        // @ts-ignore
        done(null, user);
    });

    app.use('/graphql', graphqlHTTP({
        schema: schema,
        rootValue: rootValue,
        graphiql: true
    }))

    app.get(
        "/api",
        ensureAuthenticated,
        function (req, res) {
            res.status(200).json({user: req.user});
        }
    );

    // @ts-ignore
    function ensureAuthenticated(req, res, next) {
        if (req.isAuthenticated()) {
            return next();
        }
        res.redirect('/login');
    };

    app.get(
        "/",
        function (req, res) {
            var claims = req.authInfo
            console.log("User info: ", req.user)
            // @ts-ignore
            res.status(200).json(req.user)
        }
    );

    app.get('/login',
        function (req, res, next) {
            passport.authenticate('azuread-openidconnect',
                {
                    // response: res, // supposedly required, but nah
                    failureRedirect: '/'
                }
            )(req, res, next);
        },
        function (req, res) {
            console.log('Login was called in the Sample');
            res.redirect('/');
        });

// POST /auth/openid/return
//   Use passport.authenticate() as route middleware to authenticate the
//   request.  If authentication fails, the user will be redirected back to the
//   home page.  Otherwise, the primary route function function will be called,
//   which, in this example, will redirect the user to the home page.
    app.use('/auth/openid/return',
        (req, res, next) => {
            // @ts-ignore
            passport.authenticate('azuread-openidconnect', {
                failureRedirect: '/',
                response: res,
                successRedirect: '/',
            })(req, res, next)
        },
        function (req, res) {
            res.redirect('/');
        });

    app.get('/logout', function (req, res) {
        req.logout();
        res.redirect('/');
    });

    app.listen(4000)
}

export {
    serve
}
