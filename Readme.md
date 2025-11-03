# API Types Implementations in Go

This project demonstrates all major types of APIs implemented in **Go**:
- REST API  
- GraphQL API  
- gRPC API  
- Webhooks  
- WebSockets  

Each example is kept simple and uses the same base concept — **Employees and Companies** — so you can easily compare how each API type works differently.

---

## REST API
REST (Representational State Transfer) is the most common API style.  
It uses standard HTTP methods like **GET, POST, PUT, DELETE** and works great for simple CRUD operations.

### Example
**Use Case:** Fetch and add employees.

- `GET /employees` → Returns all employees  
- `POST /employees` → Adds a new employee  

## Graphql
GraphQL allows the client to ask for exactly the data it needs.
It reduces over-fetching and under-fetching issues common in REST.

### Example
**Use Case:** Query and add employees with flexible fields.

- query { employees { name age company } }
- mutation { addEmployee(name:"Aarti", age:28, company:"JK Tyres"){ name } }

## Grpc
gRPC uses Protocol Buffers (protobuf) for efficient, binary communication over HTTP/2.
It’s ideal for internal communication between microservices.

### Example
**Use Case:** Compute employee salary using another microservice.

- RPC: ComputeSalary(SalaryRequest) → SalaryResponse

## Webhooks
Webhooks are reverse APIs — instead of clients calling the server,
the server calls another endpoint when an event occurs.

### Example
**Use Case:** Notify a company system whenever a new employee is added.

- /webhook → Receives and logs the data

## WebSocket API
WebSockets create a persistent, bi-directional connection between client and server.
Perfect for real-time features like chat, notifications, or dashboards.

### Example

**Use Case:** Simple chat system where multiple clients can send and receive live messages.

- /ws → WebSocket endpoint for connecting clients

