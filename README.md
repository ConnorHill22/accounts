# Accounts Service

## Endpoints

POST    /signup
POST    /login
POST    /password_reset

MFA Shit?

POST    /user
UPDATE  /user/<user_id>
GET     /user/<user_id>

## Development

### Env Variables
Create an .env file in the root directory with the .env.sample file in the /config directory and update the neccissary values

### Run Server

```sh
cd cmd/server
go run github.com/cosmtrek/air
```

### Database
- Install https://postgresapp.com/
    - Homebrew

```sh
psql postgres

CREATE ROLE localUser WITH LOGIN PASSWORD 'localUserpw';
ALTER ROLE localUser CREATEDB;
\q

psql postgres -U localuser

CREATE DATABASE accounts;
\q
```

### Migrations
```sh
make migrate
```

### Help
#### Fiber Framework
https://gofiber.io/