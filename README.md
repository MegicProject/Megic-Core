# Megic-Core

## Table of Contents

- [Requirements](#requirements)
- [Getting Started](#getting-started)
  - [Installation](#installation)
  - [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [License](#license)

## Requirements

Before running the application, make sure you have the following environment variables set:

- `DB_USER`: Username for connecting to the database.
- `AGENT_API_URL`: URL of the Agent API for external communication.
- `DB_PASSWORD`: Password for the database user.
- `DB_HOST`: Hostname of the database server.
- `DB_PORT`: Port number on which the database server is running.
- `DB_NAME`: Name of the database schema.

## Getting Started

Follow these steps to set up and run the Megic-Core application:

### Installation

1. **Clone the repository:**

   ```bash
   git clone <repository_url>
   cd Megic-core
   ```

2. **Set up environment variables:**

   Create a `.env` file in the root directory of your project and add the necessary environment variables:

   ```plaintext
   DB_USER=<your_db_user>
   AGENT_API_URL=<agent_api_url>
   DB_PASSWORD=<your_db_password>
   DB_HOST=<your_db_host>
   DB_PORT=<your_db_port>
   DB_NAME=<your_db_name>
   ```

   Replace `<your_db_user>`, `<agent_api_url>`, `<your_db_password>`, `<your_db_host>`, `<your_db_port>`, and `<your_db_name>` with your actual database credentials and Agent API URL.

3. **Install dependencies:**

   ```bash
   go mod tidy
   ```

### Running the Application

To run the Megic-Core application locally:

```bash
go run main.go
```

The application will start and listen on `localhost:8081` by default.


## License

Specify the license under which your project is distributed. For example:

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

