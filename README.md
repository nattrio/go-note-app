# Go Note App API

## Installation

Create `app.env`

```env
DB_HOST=${HOST}
DB_USERNAME=${USER}
DB_PASSWORD=${PASSWORD}
DB_NAME=${DBNAME}
DB_PORT=${PORT}
```

Postgres with Docker

- `docker pull postgres:alpine`
- `docker run --name note_db_test -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${PASSWORD} -p ${PORT}:5432 -d postgres:alpine`
- `docker exec -it note_db_test bash`
- `psql -U ${USER}`
- `CREATE DATABASE ${DBNAME};`

Start App

```bash
go run main.go
```

or

```bash
air -c .air.toml
```

The server will start running on http://localhost:3000.

## Endpoints

The API provides the following endpoints:

- `GET /healthchecker`: Returns a JSON response with a status and welcome message.
- `POST /notes`: Creates a new note.
- `GET /notes`: Retrieves all notes.
- `PATCH /notes/:noteId`: Updates a note.
- `GET /notes/:noteId`: Retrieves a specific note.
- `DELETE /notes/:noteId`: Deletes a note.
