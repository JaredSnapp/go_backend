


## MVP
- Clean folder structure
- CRUD for one table
- Postgres


Additional features:
- Swagger?




## Common commands 
`go mod tidy`
`go mod vendor`

`docker ps` - view docker containers
`docker network ls` - list docker networks
`docker network inspect <network name>` - inspect network

`pg_isready -p 5432` - check if postgres database is available


## Database
Compose file built from https://github.com/felipewom/docker-compose-postgres

### Setup
1. Create `postgres.env` file in `database` folder
2. Add following to `postgres.env` file
>```ENV
>POSTGRES_USER=username
>POSTGRES_PASSWORD=password
>PGADMIN_DEFAULT_EMAIL=email
>PGADMIN_DEFAULT_PASSWORD=password
>```
3. Data will be stored in `DB` folder in project folder

### Run
- `make start_db` - Run DB
- `make stop_db` - Stop DB
- `docker ps` - Check what containers are running on docker

### Maintenance
Can access DB in browser at [localhost:15433](localhost:15433)
