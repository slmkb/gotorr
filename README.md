# gotorr

gotorr is a command-line and API tool for managing custom RSS/Atom feed subscriptions and aggregating their content, with multi-user support and PostgreSQL as its backend database. This project is intended to help users register, log in, manage their feeds, and browse aggregated content efficiently.

---

## Prerequisites

To run gotorr, you **must** have:
- **Go** (version 1.18 or higher) installed. [Download Go](https://golang.org/dl/)
- **PostgreSQL** installed and running, as it is required for all persistent storage.

---

## Commands Overview

After installing and setting up PostgreSQL, you can run gotorr with the following commands:

| Command          | Description                                                                |
|------------------|----------------------------------------------------------------------------|
| `login`          | Authenticate an existing user.                                             |
| `register`       | Create a new user account.                                                 |
| `reset`          | Reset a user’s password.                                                   |
| `users`          | List all registered users.                                                 |
| `agg`            | Aggregate and display all feed items for the logged-in user.               |
| `addfeed`        | Add a new feed URL to your account (must be logged in).                    |
| `feeds`          | List all public feeds available in the system.                             |
| `follow`         | Follow a specific feed from the list (must be logged in).                  |
| `following`      | List all feeds the current user is following (must be logged in).          |
| `unfollow`       | Unfollow a feed you previously followed (must be logged in).               |
| `browse`         | Browse the latest content from feeds you follow (must be logged in).       |

> Commands marked with "must be logged in" require authentication first using the `login` command.

---

## Getting Started

1. Install required dependencies (Go, PostgreSQL).
2. Set up your PostgreSQL instance and connection details as required by the program.
3. Build and run the application using Go:
    ```bash
    go run main.go <command> [arguments]
    ```
   Or build and call the binary:
    ```bash
    go build -o gotorr
    ./gotorr <command> [arguments]
    ```

---

## License

Distributed under the MIT License.

gator
