# LetsGoPolls

LetsGoPolls is a Go and PostgreSQL-based API for managing polls with real-time updates using WebSockets. The application allows users to create and participate in polls, view live vote counts, and manage active polls with end times. It supports real-time updates to reflect votes as they are cast, providing a dynamic and interactive user experience.

## Why LetsGoPolls?
LetsGoPolls provides a powerful solution for managing polls with real-time updates using WebSockets. It allows users to instantly see voting results as they happen, making it perfect for live events or interactive decision-making processes. Built with Go and PostgreSQL, it ensures high performance and reliability, making it a great choice for anyone looking to add dynamic polling functionality to their platform.

## Features

- **Poll Creation**: Users can create new polls with a title, description, and end date.
- **Voting**: Users can vote on available options within a poll.
- **Real-Time Updates**: Instant updates are provided via WebSocket as votes are cast.
- **Poll Management**: Administrators can manage active polls, set end times, and track vote counts.

## Prerequisites

- Go 1.22
- PostgreSQL
- pressly/goose (a database migration tools)

## Getting Started

1. **Clone the Repository**
```bash
   git clone https://github.com/Raihanki/LetsGoPolls.git
   cd lets-go-polls
```
2. **Setup Environtment Variable**
```bash
   cp .env-example .env
```
Edit the .env file to configure your database and other environment-specific settings.

```bash
   goose -dir=migrations/ postgres "postgres://postgres:@localhost:5432/letsgopolls" up
```
Above is command to run migrations if you use postgres driver and letsgopolls database name, you can change the format if you have different url or database driver the format is 
goose -dir=mirgrations/ DRIVER DBSTRING up

3. **Build the Application**
To build the application, run:
```bash
   go build -o out/ ./cmd/api/*.go
```
4. **Run the Application**
Start the server
```bash
   ./out/api
```
The API will be available at http://localhost:8080 (adjust the port if necessary).

## API Documentation
You can import the provided Postman collection for API testing and exploration:
- Open Postman.
- Go to File > Import.
- Choose the LetsGoPolls.postman_collection.json file located in the root of the project.
- Import the collection to start testing the API endpoints.

## WebSocket
To connect to the WebSocket handler, use the following URL:
Start the server
```bash
   ws://localhost:8080/ws
```
This endpoint provides real-time updates on poll data and vote counts. Ensure you handle WebSocket connections appropriately in your client application.
