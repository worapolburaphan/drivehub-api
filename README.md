# DriveHub Golang Interview

## Get Started
1. Clone this repository
2. Run `go build .` to start the server
3. Open your browser and go to `http://localhost:8080/api/v1` followed by the endpoint you want to test

## Endpoints
### Health check
- `/health` - Returns the health status of the server

### Entity Cars
- GET `/cars` - Returns all cars
- POST `/cars` - Creates a new car
- PUT `/cars/{id}` - Updates a car by id
- DELETE `/cars/{id}` - Deletes a car by id
