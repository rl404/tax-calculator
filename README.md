> This is just a sample backend application.

# Tax-Calculator

A simple user's tax calculator system.

## Features

- Save and retrieve user tax
- Add user tax
- Delete user tax

## Requirements

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [Docker compose](https://docs.docker.com/compose/)

## Quick Start

1. `git clone https://github.com/rl404/tax-calculator.git`
2. `cd tax-calculator`
3. `make`
4. Wait until `server listen at :32001`
5. Endpoints are ready to use.

#### Remove containers

```
make docker-stop
make docker-rm
```

## Endpoints

- `GET` - `/v1/list`
- `POST` - `/v1/add`
- `DELETE` - `/v1/delete`

*import the `postman_collection.json` for more details.*

## Tech Stacks

### [Go](https://golang.org/)

Using golang as main programming language. Go is good for creating a simple service such as this since the compiled binary is very small (around 10 MB) which can be used to create a small docker container. Won't be using framework since this is a simple service and pretty sure won't be using all of the framework features. Using framework may also affect to compiling time and compiled binary size.

### [PostgreSQL](https://postgresql.org/)

Database to keep the user tax data. Since we already know the column needed to keep the data, we better use relational database.

Table name: `tax`

Column Name | Type | Description
--- | --- | ---
id | int | Primary key
user_id | int | User's ID
name | varchar | Tax item name
tax_code | int | [Tax code category](#tax-code)
price | numeric | Item price
created_at | timestamp | Created date
updated_at | timestamp | Updated date

### [Docker](https://www.docker.com/) & [Docker compose](https://docs.docker.com/compose/)

Quick and easy container management and deployment without installing the stack directly to local/server.

## Calculation

### Tax Code

Code | Name
--- | ---
1 | Food & Beverage
2 | Tobacco
3 | Entertainment

### Tax Calculation

Food & Beverage

- 10% of `price`
- Refundable

Tobacco

- 10 + (2% of `price`)
- Not refundable

Entertainment

- 0 < `price` < 100: tax-free
- price >= 100: 1% of (`price` - 100)
- Not refundable
