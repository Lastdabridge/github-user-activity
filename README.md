# GitHub User Activity CLI

A simple Command Line Interface (CLI) application built with Go to fetch and display the recent public activity of any GitHub user directly in your terminal. 

This project was built to practice interacting with REST APIs, handling JSON data, and creating clean CLI tools without external dependencies.

## Features

- **Zero External Dependencies:** Built entirely using Go's standard library (`net/http`, `os`, `encoding/json`).
- **Real-time Activity Fetching:** Uses the official GitHub Events API.
- **Graceful Error Handling:** Provides clear messages for network issues, missing arguments, or non-existent usernames.

## Installation

Make sure you have [Go](https://go.dev) installed on your machine.

1. Clone the repository:
   ```bash
   git clone https://github.com
   cd github-user-activity
   ```

2. Build the application:
   ```bash
   go build -o github-activity main.go
   ```

## Usage

Run the executable from your terminal and provide a GitHub username as an argument:

```bash
./github-activity <username>
```

### Example

```bash
./github-activity kamranahmedse
```

**Output:**
```text
- Pushed  to kamranahmedse/developer-roadmap
- Opened an issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
```
URL: https://roadmap.sh/projects/github-user-activity
