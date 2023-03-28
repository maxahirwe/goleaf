# Challenge

## Structure

-   "id": string, required
-   "name": string, required, min 2 characters
-   "signupTime": int64 (unix millisecond), required, min time year 1850

## Endpoints:

-   Set a user object
-   Get a user by id
-   [BONUS] List all users

Requests must be correctly validated, and the response status code must correctly match the type of error
The errors should be shown to the user!

-   [BONUS] Requires basic authentication for any operation (the credentials can be hard-coded for the sake of testing)
-   [BONUS] The listening port of the server can be configured with a command line argument

## Notes

In general, the returned errors should contain a meaningful reference to the function call chain, so that it is easy to understand where the errors happened
Please provide reference on how to run and test the app

## Installation

```
go get
```

if needed

```
go install github.com/githubnemo/CompileDaemon
```

### Development

```
CompileDaemon -command="./goleaf"
```

### Run

replace 8080 with any port number as you wish

```
export PORT=8080 && go run main.go
```

### Testing

```
go test --cover
```
