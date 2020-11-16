# challengebox-backend

## Setting up Postgres database for backend: 

### With Docker:
- Run Docker Command to start Database 
`` docker run --rm --name pg-docker -e POSTGRES_PASSWORD=docker -d -p 5432:5432 postgres ``

- Call Api-Endpoint to fill database with starting data: 
`` localhost:8080/setup ``