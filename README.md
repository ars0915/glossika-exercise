# glossika-exercise

This is a backend service project built with Go, MySQL, and Redis. It provides functionality for user registration, login, email verification, and retrieving recommended products.

## Running the Project

1. **Build and Start the Services**

   Use `docker-compose` to build and start all the services:

   ```bash
   docker-compose up -d
   ```

   This will start the following services:
    - `app`: The main application service
    - `mysql`: MySQL database service
    - `redis`: Redis cache service

   The `app` service will be accessible at `http://localhost:8080`, the `mysql` service will be available on port `3306`, and the `redis` service will be available on port `6379`.

## API Documentation

### 1. User Registration

**Request:**

```bash
curl --location 'http://localhost:8080/register' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "aaa@gmail.com",
    "password": "!Qaz2wsx"
}'
```

**Description:**
- `email` is the user's email address.
- `password` is the user's password. The password requirements are: 6 to 16 characters, containing at least one uppercase letter, one lowercase letter, and one special character.

### 2. Verify Email

**Request:**

```bash
curl --location 'http://localhost:8080/verify' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "aaa@gmail.com",
    "verificationCode": "d27a81"
}'
```

**Description:**
- `email` is the user's email address.
- `verificationCode` is the verification code obtained from the database. Although the API is supposed to get this code from the email, the actual email sending functionality is not implemented. Therefore, you need to retrieve the `verification_code` from the `users` table in the database for testing.

### 3. User Login

**Request:**

```bash
curl --location 'http://localhost:8080/login' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "aaa@gmail.com",
    "password": "!Qaz2wsx"
}'
```

**Description:**
- `email` is the user's email address.
- `password` is the user's password. You must verify the email (`verify`) before logging in, otherwise, an error will be returned.




### 4. Get Recommendations

**Request:**

```bash
curl --location 'http://localhost:8080/recommendation' \
--header 'Authorization: ••••••'
```

**Description:**
- The `Authorization` header needs to contain a valid JWT token.
- Ensure the user is verified (`verify`) and logged in (`login`) before making this request, otherwise, an error will be returned.



## Notes

1. **You need to verify the email (`verify`) before logging in.** Otherwise, the login request will return an error.
2. **The verification code (`verificationCode`) should be obtained from the `users` table in the database.** The project does not implement actual email sending, so for testing, you need to retrieve the `verification_code` from the database.

## Development and Testing

- Ensure Docker and Docker Compose are installed and running.
- Use `docker-compose up -d` to start the services.
- Test the API using the provided curl commands.
