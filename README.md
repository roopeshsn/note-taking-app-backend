# Note Taking Application Backend

Backend for note taking application written in Go using Gin, GORM.

## Getting Started

### Setting up to run the server or local development
Create a `.env` file in the root of the project and copy paste the following into that file,

```
PORT=<SERVER_PORT>
DATABASE_URL="host=<DB_HOST> user=<DB_USER> password=<DB_PASSWORD> dbname=<DB_NAME> port=<DB_PORT> sslmode=disable"
SECRET=<SECRET>
```

### Docker

```
docker compose build
docker compose up
```

## Endpoints

POST /signup
POST /login
POST /notes
GET /notes
DELETE /notes