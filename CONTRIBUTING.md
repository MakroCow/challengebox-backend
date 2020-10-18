# Contributing

## Local setup

### Database

Setup [PostgreSQL](https://www.postgresql.org/) on your computer with the data from the `config.yaml`.
Then create a database named `challengedb`, connect to it and create the `challenge` table:

```postgresql
$ psql -U postgres

CREATE DATABASE challengedb;

\c challengedb;

CREATE TABLE challenge (
 id SERIAL PRIMARY KEY,
 title varchar(255),
 description varchar(255),
 sport_value int,
 intelligence_value int,
 culinary_value int,
 selfcare_value int
);
```

**Resources:**

- [Creating PostgreSQL databases and tables with raw SQL](https://www.calhoun.io/creating-postgresql-databases-and-tables-with-raw-sql/)
- [Connecting to a PostgreSQL database with Go's database/sql package](https://www.calhoun.io/connecting-to-a-postgresql-database-with-gos-database-sql-package/)