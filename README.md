# Bookstore CRUD API

This a simple CRUD API using GO representing the backend of a bookstore. It's implemented for educational purposes.
It supports the following operations:

- Create a new Book
- Get a Book by ID
- Get all Books
- Update a Book by ID
- Delete a Book by ID


## Tools

The API is connected to a MySQL database and uses the following libraries:
- [Gorm](https://gorm.io/)
- [Gorilla Mux](https://github.com/gorilla/mux)


## How to run
1. Clone the repository
    ```bash
    git clone https://github.com/AhmadHoseiny/Bookstore-CRUD-API.git
    ```
2. Configure the database connection parameters in `DBConnector.go` (username, password and database name)
3. Run `go run main.go`
4. Open Postman and test the API using the following endpoints:
    - Create a new Book: `POST http://localhost:5000/book`
    - Get a Book by ID: `GET http://localhost:5000/book/{ID}`
    - Get all Books: `GET http://localhost:5000/book`
    - Update a Book by ID: `PUT http://localhost:5000/book/{ID}`
    - Delete a Book by ID: `DELETE http://localhost:5000/book/{ID}`
