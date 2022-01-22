# BardView5

A Fifth Edition Dungeons & Dragons Campaign Management Tool 

## Roadmap

- [ ] User
    - [X] Login
    - [X] Registration
    - [ ] Email verification
    - [ ] Password reset
- D&D 5e
    - World
        - [X] Read
        - [X] Create
        - [ ] Update
        - [ ] Delete
- Instrumentation
    - [X] Prometheus

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


