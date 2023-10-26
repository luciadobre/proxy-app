# proxy-app

Proxy app (backend)
The objective is to create a Go application that acts as a reverse proxy to a predetermined target and modifies the HTTP response.
Specifically, when you make a request to your application (http://127.0.0.1:8080/) on the route /todos/1, it should make an HTTP request to https://jsonplaceholder.typicode.com/todos/1 and modify the JSON response by adding a random attribute (e.g., "foo": "bar") to the JSON object that will be forwarded. This rule applies to all routes on https://jsonplaceholder.typicode.com that respond with JSON (you don't need to search for them; you can simply check the content-type header in the response).
BONUS:

- Add rate limiting per IP (maximum X requests from an IP within an interval of Y seconds - you can choose X and Y).
- Save the requests and responses in a database (you can choose any database you prefer).
- Use the hexagonal architecture: https://medium.com/@matiasvarela/hexagonal-architecture-in-go-cfd4e436faa3.
- Aim for code coverage of 80% or more with unit tests.
