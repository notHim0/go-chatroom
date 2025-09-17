# Chatroom API

Chatroom built using go, gin and gorilla websockets where you can join and create chatroom with user authentication.

---

## 🚀 Getting Started

Follow these steps to set up and run the project locally.

---

### 1. Clone the Repository

Clone the project from GitHub and navigate into the directory.

```bash
git clone https://github.com/notHim0/go-chatroom.git
cd kanban
```

### 2. Install Dependencies

Install all the required Go packages.

```bash
go mod tidy
```

### 3. Setup Environment Variables

Create a file named .env in the root directory and add your database URL and a JWT secret key.

```
DATABASE_URI=your_postgres_database_uri
JWT_SECRET=your_jwt_secret_key
```

### 4. Run the Application

Start the application by running the main Go file.

```bash
go run server/cmd/main.go
```

### 5. API Endpoints

The API provides the following core functionalities:

- POST : "/signup"
- POST : "/login"
- GET  : "/logout"
- POST : "/ws/createRoom"
- GET  : "/ws/joinRoom/:roomId"
- GET  : "/ws/getRooms"
- GET  : "/ws/getClients/:roomId"
