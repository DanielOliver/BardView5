# BardView5

A self-hosted Fifth Edition Dungeons & Dragons Campaign Management Tool. 

## Screenshots

![Landing page](e2e/screenshots/unauth_landing.png)

## Roadmap

1. [X] Start with basic user management
2. [ ] Become a compendium of worlds and monsters.
3. [ ] Become a campaign/session/encounter tracker for the game master.
4. [ ] Create a display window for tv so that DM can share information with players. 
5. [ ] Explore player data management.

### Roadmap Details

- User
    - [X] Login
    - [X] Registration
    - [X] Authenticate
    - [ ] Email verification
    - [ ] Password reset
- D&D 5e
    - World: A general setting to play in
        - [X] Read
        - [X] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Assigning user permissions
    - Monster: A single monster with statblocks and info
        - [ ] Read
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Associating with worlds
    - Inhabitant: Associating a monster with one or more worlds
        - [ ] Read
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Associating with worlds
    - Data pack: An entire collection of worlds, monsters, and associated information. Meant to be imported or exported and shared.
        - [ ] Export
        - [ ] Import
    - Player tv display:
        - [ ] Display encounter information on secondary webpage with a sharable link.
    - Campaign: A collection of players and game masters associated with sessions and encounters
        - [ ] Read
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Associating with players
    - Session: zero or more encounters
        - [ ] Read
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Associating with Campaign
    - Encounter
        - [ ] Read
        - [ ] Create
        - [ ] Update
        - [ ] Delete
        - [ ] Query
        - [ ] Associating with Session
- Instrumentation
    - [X] Prometheus
- Configuration
    - [X] Local development
    - [ ] Production

## Development

### Prerequisites

* Go 1.17
* [Mage](https://magefile.org/)
* Docker
* Docker Compose
* Node 17
* NPM 8

I use [JetBrains products](https://www.jetbrains.com/) but I'd expect VSCode or another to work as well.

Operating system: I alternate development on openSUSE Tumbleweed and Windows 10 professional. GitHub Actions runs Ubuntu-latest.

### Dependent services

* PostgreSQL
* Ory Kratos
* Nginx

### Running locally

1. [Add "proxy.local" to point to localhost on your machine.](https://linuxize.com/post/how-to-edit-your-hosts-file/)
2. Confirm docker is running.
3. Open server folder in root of this repository.
   ```powershell
   cd server
   ```
4. Start dependencies to open docker-compose and run database migrations
   ```powershell
   mage up
   ```
   ```
   Creating network "server_default" with the default driver
   Creating server_mailslurper_1    ... done
   Creating server_db_1             ... done
   Creating server_kratos-migrate_1 ... done
   Creating server_kratos_1         ... done
   Creating server_nginx_1          ... done
   {"level":"info","role":"bardview5","host":"PeanutButter","type":"Migration","migration_version":0,"migration_dirty":false,"time":"2022-01-21T18:49:14-05:00","message":"Migrating bardview5"}
   {"level":"info","role":"bardview5","host":"PeanutButter","type":"Migration","migration_dirty":false,"migration_version":0,"time":"2022-01-21T18:49:15-05:00","message":"Migrated bardview5"}
   {"level":"info","role":"bardview5","host":"PeanutButter","type":"Migration","migration_dirty":false,"migration_version":1,"time":"2022-01-21T18:49:15-05:00","message":"bardview5 database version"}
   Run: DumpSchema
   ```
5. Go mod download to install
   ```powershell
   go mod download
   ```
6. Go run server
   ```powershell
   $env:bardview5_connection = "postgresql://postgres:mysecretpassword@localhost/bardview5?sslmode=disable"
   ```
   ```powershell
   go run . serve
   ```
7. Open bv-app folder in root of this repository.   
   ```powershell
   cd bv5-app
   ```
8. Node run server
   ```powershell
   npm run start-linux
   ```
   or 
   ```powershell
   npm run start-windows
   ```
9. Open "http://proxy.local/"

## Changing OpenAPI surface.

[OpenAPI](server/bardview5.yaml) shows the API surface and schema definitions for the API.

1. Edit schema.
2. Generate Go files
   ```powershell
   cd server
   go generate ./...
   ```
3. Generate TypeScript files
   ```powershell
   cd bv5-app
   npm run generate
   ```
4. Implement API endpoints in Go.

## Changing Database

[Migrations](server/migrations) shows the database migrations.

1. Add database migration
2. Add queries to be translated to Go using [sqlc](https://sqlc.dev/) to [query.sql](server/db/query.sql)
3. Generate Go files
   ```
   cd server
   mage -v InternalMigrate
   ```

## Methodology

Development proceeds as I have time, usually a couple lines very late at night or very early in the morning. Further, I have no plans of monetization at this time. Neither am I optimizing for "web-scale" As such, the development habits here generally reflect my mood of what I want to explore at the time. Further, this isn't my job, so please don't judge me too harshly on what I write here and shortcuts taken since my free time is valuable.

The things I value here are:

* Automating as much boilerplate as possible.
* Using stable technologies in order to smooth out difficulties. The adventure here is in building a complete application instead of bleeding edge technology.
* Speed of development.


