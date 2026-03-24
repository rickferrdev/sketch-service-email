# 📧 Sketch Service Email

**sketch-service-email** is a microservice developed in Go, focused on asynchronous email processing and delivery.

> **Study Project:** This repository was created exclusively for learning purposes. The main goal is to explore modern framework integration, dependency management, and concurrency patterns in Go.

---

## 🛠️ Technologies & Concepts Explored

The project utilizes a robust stack to simulate a real-world production environment:

*   **Go (1.25.5):** The core language of the project.
*   **Uber fx:** Used for **Dependency Injection**, facilitating modularization and application lifecycle management.
*   **Fiber v3:** A high-performance web framework for building API routes.
*   **Gomail:** A library for constructing and sending emails via SMTP.
*   **Concurrency (Worker Pool):** Implementation of queues using `channels` for background email processing, preventing HTTP request blocking.

---

## 🏗️ System Architecture

The project follows a modular structure:

*   **`config/`**: Environment variable management using [dotenv](https://github.com/rickferrdev/dotenv).
*   **`internal/handlers/`**: The transport layer (HTTP) that handles incoming subscription requests.
*   **`internal/services/`**: Business logic layer, where subscription and email preparation reside.
*   **`pkg/mail/`**: A utility package responsible for direct SMTP communication and Worker management.

---

## 🚀 Getting Started

### Prerequisites
*   Go 1.25.5 or higher installed.
*   An SMTP server for testing (e.g., Mailtrap).

### Setup Instructions
1.  **Clone the repository:**
    ```bash
    git clone https://github.com/rickferrdev/sketch-service-email.git
    ```

2.  **Configure environment variables:**
    Create a `.env` file in the root directory based on the keys defined in the code:
    ```env
    MAIL_PORT=587
    MAIL_HOST=smtp.example.com
    MAIL_USERNAME=your_username
    MAIL_PASSWORD=your_password
    ```

3.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

4.  **Run the application:**
    ```bash
    go run main.go
    ```
    The server will start on port **8080**.

---

## 📨 Main Endpoint

### Create Subscription
`POST /api/v1/subs/`

**Request Body:**
```json
{
  "email": "user@example.com"
}
```

**Internal Workflow:**
1.  **Handler**: Validates the request body.
2.  **Service**: Generates the email content.
3.  **Mailer**: Adds the message to a **channel**.
4.  **Worker**: A background worker consumes the channel and executes the SMTP delivery.

---

## 📝 License

This project is licensed under the **MIT License**.

---
