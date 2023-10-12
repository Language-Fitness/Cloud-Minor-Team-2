
# User Management Repository

This repository contains two applications: a Golang GIN API and a gRPC service for managing user data.

## GIN API

The GIN API REST API build with the GIN framework. It provides endpoints for retrieving user information. Currently, it supports the following features:

- Retrieve all users.
- Retrieve a user by their ID.

### Getting Started

To run the GIN API:

1. Make sure you have Golang installed on your system.
2. Clone this repository.
3. Navigate to the GIN API directory.
4. Run the API with `go run main.go`.

### Usage

- To retrieve all users, make a GET request to `/user/`.
- To retrieve a user by their ID, make a GET request to `/user/{id}`, where `{id}` is the user's ID.

## gRPC Service

The gRPC service is a remote procedure call service for managing user data. It provides methods for retrieving user information. The service interacts with a database or data source to fetch user data.

### Getting Started

To run the gRPC service:

1. Make sure you have Golang installed on your system.
2. Clone this repository.
3. Navigate to the gRPC service directory.
4. Run the service with `go run UserServiceServer.go`.

### Usage

The gRPC service exposes the following methods:

- `GetUserByID`: Retrieve a user by their ID.
- `GetAllUsers`: Retrieve all users.

You can interact with the service using gRPC client libraries in various programming languages.


## License

This repository is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

-
