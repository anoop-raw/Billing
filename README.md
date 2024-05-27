# Amarth Billing Service

## Features
- GetOutstanding : This returns the current outstanding on a loan, 0 if no outstanding(or closed),
- IsDelinquent : If there are more than 2 weeks of Non payment of the loan amount
- MakePayment: Make a payment of certain amount on the loan

## Technologies Used
- Go (Golang) programming language
- Fiber web framework
- Docker
- PostgreSQL

## Installation and running
1. Clone the repository:
    ```bash
    git clone git@github.com:anoop-raw/Billing.git
    ```
2. move to pismo folder
    ```bash
    cd Billing
    ```
3. Update the `.env` file with appropriate values for your environment.
4. Do docker setup on local
5. Start the Docker containers:
    ```bash
    docker-compose up
    ```
6. run init.sql file to create operation types
7. Access the application at `http://localhost:8080`.

## Run Test
1. command for running unit test
    ```bash
    make test-long
    ```
   
## Curl for postman setup
1. create loan schedule:
    ```bash
    curl --location 'http://localhost:8000/v1/loans/create' \
    --header 'Content-Type: application/json' \
    --data '{
    "amount": 100,
    "interest_rate": 10,
    "weeks": 5
    }'
    ```
2. Get Loan:
   ```bash
   curl --location --request GET 'http://localhost:8000/v1/loans/1' \
   --header 'Content-Type: application/json' \
   --data '{
   "amount": 1000000,
   "interest_rate": 10,
   "weeks": 50
   }'
   ```
3. MakePayment:
    ```bash
    curl --location 'http://localhost:8000/v1/loans/1/payments' \
    --header 'Content-Type: application/json' \
    --data '{
    "week": 1,
    "amount": 20.19
    }'
    ```

4. Get Outstanding amount:
    ```bash
    curl --location 'http://localhost:8000/v1/loans/1/outstanding'
    ```

5. Check IsDelinquent:
    ```bash
    curl --location 'http://localhost:8000/v1/loans/1/delinquent?weekNumber=3'
    ```


