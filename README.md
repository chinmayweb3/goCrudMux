# A SIMPLE CRUD API using Go

This project demonstrates the development of a CRUD API using Go. It showcases the implementation of Create, Read, Update, and Delete operations, following RESTful principles to ensure a scalable and maintainable solution.

## Folder Structure

The codebase is organized into separate folders according to the context of the code. Here's a brief description of the key folders:

- `routes`: Contains the route handlers for the API endpoints.
- `database`: Contains the database-related code.
- `models`: Contains the struct-related code.

## API Endpoints

The API includes five key endpoints:

1. **Get All Movies:** Retrieves all movies stored in the database. Route: `/movies`, Method: `GET`,
2. **Get Single Movie by Name:** Searches for a specific movie by its name. Route: `/movie?name={movie_name}`, Method: `Get`
3. **Add Movie to Database:** Adds a new movie to the database. Route: `/addmovie`, Method: `POST`
4. **Delete Movie from Database by Name:** Removes a movie from the database based on its name. Route: `/deletemovie`, Method: `DELETE`
5. **Get Random Quotes:** Returns a random quote. Route: `/` , Method: `GET`

## Running the Project

To run this project locally, clone the repository and execute the main file with Go.
