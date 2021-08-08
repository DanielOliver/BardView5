// Construct a schema, using GraphQL schema language
import { buildSchema } from 'graphql'
import { graphqlHTTP } from 'express-graphql'
import express from 'express'

const schema = buildSchema(`
  type Query {
    posts: [Post]
    post(id: ID): Post
    authors: [Person]
    author(id: ID): Person
  }

  type Post {
    id: ID
    author: Person
    body: String
  }

  type Person {
    id: ID
    posts: [Post]
    firstName: String
    lastName: String
  }
`)

const PEOPLE = new Map()
const POSTS = new Map()

class Post {
  constructor (data: any) { Object.assign(this, data) }
  get author () {
    // @ts-ignore
    return PEOPLE.get(this.authorId)
  }
}

class Person {
  constructor (data: any) { Object.assign(this, data) }
  get posts () {
    // @ts-ignore
    return [...POSTS.values()].filter(post => post.authorId === this.id)
  }
}

// The root provides a resolver function for each API endpoint
const rootValue = {
  posts: () => POSTS.values(),
  // @ts-ignore
  post: ({ id }) => POSTS.get(id),
  authors: () => PEOPLE.values(),
  // @ts-ignore
  author: ({ id }) => PEOPLE.get(id)
}

const initializeData = () => {
  const fakePeople = [
    { id: '1', firstName: 'John', lastName: 'Doe' },
    { id: '2', firstName: 'Jane', lastName: 'Doe' }
  ]

  fakePeople.forEach(person => PEOPLE.set(person.id, new Person(person)))

  const fakePosts = [
    { id: '1', authorId: '1', body: 'Hello world' },
    { id: '2', authorId: '2', body: 'Hi, planet!' }
  ]

  fakePosts.forEach(post => POSTS.set(post.id, new Post(post)))
}

initializeData()

const app = express()
app.use('/graphql', graphqlHTTP({
  schema: schema,
  rootValue: rootValue,
  graphiql: true
}))
app.listen(4000)
console.log('Running a GraphQL API server at http://localhost:4000/graphql')
