# Go Backend - Learning Project

## Description

This repository contains a learning project for a Go backend.

The project relies heavily on the following libraries:

- [GQLGen](https://gqlgen.com/) - GraphQL server
  Schema first approach, let the library generate the code (models, resolvers, etc.)

- [Ent](https://entgo.io/) - Database ORM
  I played with [SQLBoiler](https://github.com/volatiletech/sqlboiler), [SQLX](https://github.com/jmoiron/sqlx) and given
  the simplicity of the project, I decided to go with an ORM that generates the code for me.
  If there is a need for more complex queries, I will probably switch to SQLX or mix the two.

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Task](https://taskfile.dev/#/installation)
- [Pre-commit](https://pre-commit.com/#install)
- [Atlas](https://atlasgo.io/)

### Setup

1. Clone the repository
2. Create a copy of the `.env.example` file and name it `.env`
3. Fill in the environment variables in the `.env` file
4. Run `task setup` to install the dependencies and setup the database
5. Run `task dev` to start the server

### Development

- `task dev` - Start the server
- `task gen` - Code generation
    - Generates GraphQL code
    - Generates ent code
    - Generates dataloaders code

## Testing

> Hey, you should write tests!

Indeed, I should. I will. I promise. Someday.
