## Description

## Installation

```bash
$ npm install
```

## Running the app

```bash
# development
$ npm run start

# watch mode
$ npm run start:dev

# production mode
$ npm run start:prod
```

## Running dependencies

```bash
docker run --name bardview5 -p 5432:5432 --rm -e POSTGRES_PASSWORD=mysecretpassword -e POSTGRES_DB=bardview5 -d postgres
```

## Test

```bash
# unit tests
$ npm run test

# e2e tests
$ npm run test:e2e

# test coverage
$ npm run test:cov
```
