Sidra Api - Basic Auth Plugin
# Basic Auth Plugin

## Description

The Basic Auth Plugin is a simple authentication plugin that validates credentials (username and password) using the Authorization header in Basic Authentication format. This plugin ensures that each incoming request has valid credentials before further processing.

## How It Works

- Username and Password are stored as Environment Variables:
  - `BASIC_AUTH_USERNAME`
  - `BASIC_AUTH_PASSWORD`
- Each request must include an Authorization Header in the following format:
  ```
  Authorization: Basic <base64-encoded-credentials>
  ```
- The credentials are in the `username:password` format, then encoded using Base64.
- If the credentials are valid, the request will be responded to with a `200 OK` status.
- If the credentials are invalid or the header is missing, a `401 Unauthorized` status will be returned.

## Prerequisites

- Golang must be installed.
- Sidra Api and Sidra Plugin Hub are properly configured.
- The following environment variables must be set:
  - `BASIC_AUTH_USERNAME`
  - `BASIC_AUTH_PASSWORD`

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/sidra-api/plugin-basic-auth.git
   cd plugin-basic-auth
   ```
2. Set Environment Variables:
   ```sh
   export BASIC_AUTH_USERNAME=admin
   export BASIC_AUTH_PASSWORD=secret
   ```
3. Build and Run the Plugin:
   ```sh
   go build -o basic-auth-plugin
   ./basic-auth-plugin
   ```

## Testing the Plugin

1. Use a tool like Postman or curl to test the plugin.
2. Base64 encode the credentials `username:password`:
   ```sh
   echo -n "admin:secret" | base64
   ```
   Result: `YWRtaW46c2VjcmV0`
3. Send a request with the Authorization header:
   ```sh
   curl -X GET http://localhost:8080 \
     -H "Authorization: Basic YWRtaW46c2VjcmV0"
   ```

### Expected Responses

- `200 OK`: Authentication successful.
  ```
  You are authenticated
  ```
- `401 Unauthorized`: Invalid or missing credentials.
  ```
  Unauthorized - Invalid Credentials
  ```

## Environment Variables

| Variable             | Description                  | Example |
|----------------------|------------------------------|---------|
| `BASIC_AUTH_USERNAME`| Username for authentication  | admin   |
| `BASIC_AUTH_PASSWORD`| Password for authentication  | secret  |

## Notes

- Ensure the plugin is running and listening on the correct socket or port.
- The plugin is designed to integrate seamlessly with Sidra Api.

## License

This project is licensed under the MIT License.