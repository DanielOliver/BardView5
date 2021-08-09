import express from "express"
import helmet from "helmet"
import session from 'cookie-session'
import {buildSchema} from 'graphql'
import {graphqlHTTP} from 'express-graphql'
import {rootValue, schema} from "./schema"
import passport from "passport"
import OAuth2Strategy from 'passport-oauth2'
import morgan from 'morgan'
import jwt from 'jsonwebtoken'


const tokenToProfile = async (idToken: string) => {

    // !!<-- Make sure you validate the token's signature -->!!
    // And make sure you handle errors. I simplified the code for the blog post

    const profile = jwt.decode(idToken);

    return profile
};

function serve(isProduction: boolean) {
    //REDO IT WITH: https://github.com/AzureAD/microsoft-authentication-library-for-js/tree/dev/maintenance/passport-azure-ad#8-samples-and-documentation
    passport.use(new OAuth2Strategy({
            authorizationURL: process.env.OAUTH_AUTH_URL || 'test',
            tokenURL: process.env.OAUTH_TOKEN_URL || 'test',
            clientID: process.env.OAUTH_CLIENT_ID || 'test',
            clientSecret: process.env.OAUTH_CLIENT_SECRET || 'test',
            callbackURL: process.env.OAUTH_CALLBACK_URL || 'test',
            scope: 'profile email openid',
            state: true,
            passReqToCallback: true
        },
        // @ts-ignore
        function (accessToken, refreshToken, profile, cb) {
            console.log('profile', profile)
            return cb(null, {
                'user': 'example',
                'id': 'test@test.com',
                'accessToken': accessToken
            })
        }
    ));
    passport.serializeUser(function (user, done) {
        console.log('serializeUser', user)
        // @ts-ignore
        done(null, user);
    });

    passport.deserializeUser(function (user, done) {
        console.log('deserializeUser', user)
        // @ts-ignore
        done(null, user)
    });

    const app = express()
    app.use(helmet({contentSecurityPolicy: (isProduction) ? undefined : false}))
    const expiryDate = new Date(Date.now() + 60 * 60 * 1000) // 1 hour
    app.use(morgan('combined'))
    app.use(session({
        name: 'session2',
        keys: ['keytar'],
        secure: isProduction,
        httpOnly: isProduction,
        expires: expiryDate
    }))
    app.use(passport.initialize());
    app.use(passport.session());


    app.use('/graphql', graphqlHTTP({
        schema: schema,
        rootValue: rootValue,
        graphiql: true
    }))

    app.get('/logout', function (req, res) {
        req.logout();
        console.log('logout', req.user)
        res.redirect('/graphql');
    });

    app.get(
        "/api",
        passport.authenticate("oauth2", {session: false}),
        function (req, res) {
            var claims = req.authInfo;
            console.log("User info: ", req.user);
            console.log("Validated claims: ", claims);
            // @ts-ignore
            res.status(200).json({name: claims["name"]});
        }
    );


    app.get(
        "/",
        passport.authenticate("oauth2", {session: false}),
        function (req, res) {
            var claims = req.authInfo
            console.log("User info: ", req.user)
            console.log("Validated claims: ", claims)
            // @ts-ignore
            res.status(200).json({name: claims["name"]})
        }
    );


// Redirect the user to the OAuth 2.0 provider for authentication.  When
// complete, the provider will redirect the user back to the application at
//     /auth/provider/callback
    app.get('/auth/provider', passport.authenticate('oauth2'));

    app.get('/auth/provider/callback',
        passport.authenticate('oauth2', {failureRedirect: '/login'}),
        function (req, res) {
            // Successful authentication, redirect home.
            res.redirect('/');
        });

    app.listen(4000)
}

export {
    serve
}
