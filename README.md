# A simple commenting service using RESTful APIs

The application is designed on the CRUD system of commenting and additionally implemented with JWT authorization to post comments.

## Application Overview

Designed code in this project is following clean architecture principles.

Application contains 3 distinct layers:

* **The Transport Layer** - responsible for handling incoming HTTP requests and passing them on to the relevant service functions.
* **The Service Layer** - responsible for all the business logic in the application.
* **The Repository Layer** - responsible for all the interactions with the database.

## Technologies Used:

* **Postgres**
* **Docker + Docker-Compose**
* **Taskfile**
* **Postman**

## Frameworks + Libraries Used

* **sqlx**
* **golang-migrate**
* **dgrijalva/jwt-go**
* **satori/go.uuid**
* **sirupsen/logrus**
* **stretchr/testify**

## How to run

- Run unit tests:

```bash
task test
```

- Run server:

```bash
task run
```

- Run integration test:

```bash
task integration test
```
