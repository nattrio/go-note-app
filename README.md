# Go Note App API

## Description
Go Note App is a simple note-taking application built with Go. It allows users to create, read, update, and delete notes via a RESTful API.

## Installation

To install and run Go Note App locally, follow these steps:

1. Clone the repository:

```shell
git clone https://github.com/nattrio/go-note-app.git
```

2. Navigate to the project directory:

```shell
cd go-note-app
```

3. Create a `.env` file in the project root directory and configure the required environment variables. You can use the provided `.env.example` file as a template.

4. Build and run the application using Docker Compose:

```shell
docker-compose up --build
```
> This command will build the Docker image and start the application and its dependencies (PostgreSQL database) as Docker containers.

5. The Go Note App API should now be accessible at http://localhost:3000. You can send HTTP requests to this endpoint to interact with the application.

## Endpoints

The API provides the following endpoints:

- `GET /healthchecker`: Returns a JSON response with a status and welcome message.
- `POST /notes`: Creates a new note.
- `GET /notes`: Retrieves all notes.
- `PATCH /notes/:noteId`: Updates a note.
- `GET /notes/:noteId`: Retrieves a specific note.
- `DELETE /notes/:noteId`: Deletes a note.
