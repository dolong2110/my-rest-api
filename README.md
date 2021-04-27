# Build a RESTFUL API With Go and MongoDB

### I. Setup Environment
##### - First download and install docker
##### - Get MongoDB image: 
$ docker pull mongo
##### - When download complete run:
$ docker run --name mongo-db -p 27017:27017 -d mongo:latest
##### - The port parameter, -p 27017:27017, is very important. It tells the Docker container to expose this port to be accessed from outside the container. Without this, you will not be able to connect to the database from the application.
##### - With the container running, we can now access the MongoDB instance with the URL:
mongodb://localhost:27017
##### - Getting libraries: We’ll be using Fiber for the web framework and MongoDB as a database.
$ go get github.com/gofiber/fiber
$ go get go.mongodb.org/mongo-driver/mongo
### II. Coding
### 1. Set up http routes
##### - The main function instantiates a Fiber instance, declares which routes will be served and its handlers, and starts listening on port 8000.
### 2. Connecting to MongoDB
##### - Mainly we need two functions — the first is responsible for the connection with the database, the second will return a Collection object in which we can execute the operations.
### 3. Building Handlers
##### The operations are the simplest ones
##### - Get: Returns one or multiple results.
##### - Post: Adds one entry to the database.
##### - Put: Updates one entry.
##### - Delete: Removes one entry.
##### a. POST
##### We will start with the POST method because of how the MongoDB works. The database is only really created when the first document is inserted, so we need to insert at least one person to create the database itself.
##### In the context variable, we have access to the body of the request through the function Body, so we need to call it and parse to our Person object . Then we call MongoDB to insert it in the database. If there’s an error, we return the status code 500, which means “internal server error,” and send the description of the error that the MongoDB returns. If the insert operations run successfully, we parse the response to JSON and send it back with the response.


### III. References
https://betterprogramming.pub/building-a-restful-api-with-go-and-mongodb-93e59cbbee88