# Challenge

## Structure

-   "id": string, required
-   "name": string, required, min 2 characters
-   "signupTime": int64 (unix millisecond), required, min time year 1850

## Endpoints:

[END-POINTS DOCUMENTATION PUBLISHED](https://documenter.getpostman.com/view/16879881/2s93RRxDqH)

-   [x] Set a user object
-   [x] Get a user by id
-   [x] [BONUS] List all users

Requests must be correctly validated, and the response status code must correctly match the type of error
The errors should be shown to the user!

-   [x] [BONUS] Requires basic authentication for any operation (the credentials can be hard-coded for the sake of testing) (user:`idt`, pass:`leaf`)
-   [x] [BONUS] The listening port of the server can be configured with a command line argument (can be set in .env or passed with command)

## Notes

In general, the returned errors should contain a meaningful reference to the function call chain, so that it is easy to understand where the errors happened
Please provide reference on how to run and test the app

## WHY's

-   `Gin`: light weight, high perfomance, great documentation, fantastic middleware pattern that I'm familiar with in nodejs/express

-   `GO ORM`: Full-Featured ORM, associations, Migrations, Support community and more

-   `SQLite`: quick development without having to install specific dbms & drivers

## Installation

```
go get
```

### Run

-   create .env file from the provided sample

    ```
    cp .env.example .env
    ```

-   replace 8080 with any port number as you wish

    ```
    export PORT=8080 && go run main.go
    ```

### Testing

```
go test
```

## Extra

-   Use https://sqlitebrowser.org/dl/ to investigate the sqllite file
-   ![testcases.png](/documentation/testcases.png)
-   [messagingApp-architecture(maxahirwe)](</documentation/messagingApp-architecture(maxahirwe).pdf>)

## Author

[@maxahirwe](https://max.rw)
