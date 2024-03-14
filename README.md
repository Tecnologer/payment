# Deuna Payment

Deuna Payment is a payment gateway that allows you to accept payments from your customers.

## Structure

The project is divided into 3 microservices:

1. [auth](./auth): It's a pretty simple microservice that generates JWT tokens and validates them.
2. [bank](./bank): It's a microservice that simulates a bank. It allows get accounts and transfers between them.
3. [gatepay](./gatepay): It's the main microservice. It allows to create, refund and get payments.

## Pre-requisites and justifications

1. [go1.22][5]: I used go1.22 to build the API because it's the last stable version of go.
2. [enumer][6]: enumer allows to simplify the creation of enums in go. It creates the methods to work with SQL and JSON.
3. [gorilla/mux][8]: A powerful URL router and dispatcher for golang.
4. [logrus][7]: logrus is a structured logger for Go (golang).
5. [postgres][9]: I decided to use postgres because it's a powerful and open-source database. It's also easy to use and has a lot of documentation.
6. [gorm][12]: Simplify the interaction with the database and add speed to the development.
7. [Docker][2]: Docker allows to pack and run the app easier.
8. [Docker Compose][3]: Docker Compose is a tool for defining and running multi-container Docker applications.
9. [Postman][4]: Postman is a collaboration platform for API development. Postman's features simplify each step of building an API and streamline collaboration so you can create better APIs—faster.
10. [golangci-lint][10]: golangci-lint is a fast linters runner for Go. It runs linters in parallel, uses caching, and works well with all version of Go.
    - I made a gist to easily install and configure golangci-lint. You can find it [here][11].

# Installation

1. Clone the repository
    - `git clone git@github.com:tecnologer/payment.git deuna-payment`
2. Access the project folder.
   - `cd deuna-payment`
3. Run composer 
   - `make docker-run-all`



# Usage

Use the postman collection for more details on how to use the API.

1. Open Postman
2. Import the collection `postman_collection.json`. More info [here][1]. 

[1]: https://learning.postman.com/docs/getting-started/importing-and-exporting/importing-data/
[2]: https://docs.docker.com/get-docker/
[3]: https://docs.docker.com/compose/install/
[4]: https://www.postman.com/
[5]: https://golang.org/doc/install
[6]: https://github.com/dmarkham/enumer
[7]: https://github.com/sirupsen/logrus
[8]: https://github.com/gorilla/mux
[9]: https://www.postgresql.org/
[10]: https://golangci-lint.run/
[11]: https://gist.github.com/Tecnologer/9051643d839913294f3570bd9920a022
[12]: https://gorm.io/
