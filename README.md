# Go Movies CRUD API

This is a simple CRUD (Create, Read, Update, Delete) API for managing movies, built using Go and the Gorilla Mux router.

## Features

- **Get All Movies**: Retrieve a list of all movies.
- **Get a Single Movie**: Retrieve details of a specific movie by its ID.
- **Create a Movie**: Add a new movie to the collection.
- **Update a Movie**: Update details of an existing movie.
- **Delete a Movie**: Remove a movie from the collection.

## Project Structure

- `main.go`: Contains the main application logic, including route handlers and server setup.

## Endpoints

### 1. Get All Movies

- **URL**: `/movies`
- **Method**: `GET`
- **Response**: JSON array of all movies.

### 2. Get a Single Movie

- **URL**: `/movies/{id}`
- **Method**: `GET`
- **Response**: JSON object of the requested movie.

### 3. Create a Movie

- **URL**: `/movies`
- **Method**: `POST`
- **Request Body**:
  ```json
  {
    "isbn": "string",
    "title": "string",
    "director": {
      "firstName": "string",
      "lastName": "string"
    }
  }
  ```
- **Response**: JSON object of the created movie.

### 4. Update a Movie

- **URL**: `/movies/{id}`
- **Method**: `PUT`
- **Request Body**:
  ```json
  {
    "isbn": "string",
    "title": "string",
    "director": {
      "firstName": "string",
      "lastName": "string"
    }
  }
  ```
- **Response**: JSON object of the updated movie.

### 5. Delete a Movie

- **URL**: `/movies/{id}`
- **Method**: `DELETE`
- **Response**: JSON array of remaining movies.

## How to Run

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```
2. Navigate to the project directory:
   ```bash
   cd go-movies-crud
   ```
3. Install dependencies:
   ```bash
   go mod tidy
   ```
4. Run the application:
   ```bash
   go run main.go
   ```
5. Access the API at `http://localhost:8000`.

## Dependencies

- [Gorilla Mux](https://github.com/gorilla/mux): A powerful HTTP router and URL matcher for building Go web applications.

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
