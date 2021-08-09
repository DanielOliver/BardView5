// Construct a schema, using GraphQL schema language
import {serve} from './server'
import dotenv from 'dotenv'

const dotenvResult = dotenv.config()
if (dotenvResult.error) {
    throw dotenvResult.error
}

const isProduction = process.env.NODE_ENV === 'production'
serve(isProduction)
console.log('Running a GraphQL API server at http://localhost:4000/graphql')
