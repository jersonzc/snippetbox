# Snippetbox
A web application to create and view text snippets.

## Requirements
- Docker & Docker Compose.
- [Go 1.23.11](https://go.dev/doc/install) and [MySQL 9.3](https://www.mysql.com/downloads/) (if running locally without containers).

## Setup Instructions
### 1. Prepare your environment
Copy the `.env.example` file to `.env` to set up environment variables:
```sh
cp .env.example .env
```
### 2. Launch everything with Docker Compose
```bash
docker compose up --build
```
This will spin up:
- The application running at [https://localhost:5002](https://localhost:5002).
- A MySQL database with the initial schema setup.

## References
Some parts of this code are based on [Alex Edwards's book Let's Go](https://lets-go.alexedwards.net/).