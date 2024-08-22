# Email Service

This project implements a flexible and extensible email service that supports multiple email providers.

## Approach

Our email service is designed with the following key features:

1. **Facade Pattern**: The service implements the Facade design pattern to provide a simplified interface to a complex subsystem of email providers.

2. **Failover Mechanism**: If one provider fails to send an email, the service automatically tries the next available provider.


## How to Use

Set your environment variables:

```shell
export SENDGRID_API_KEY=sendgrid_api_key
export MAILGUN_API_KEY=mailgun_api_key
export MAILGUN_DOMAIN=mailgun_domain
```

## Running the Application

To run the application, you have several options:

1. Directly using Go:

```shell
go run main.go -from "SENDER_MAILID" -to "RECEIVER_MAILID"
```

2. Using Docker:

Build the Docker image:
```shell
docker build -t email-service .
```

Run the Docker container:
```shell
docker run --rm -it email-service:latest --from="SENDER_MAILID" --to="RECEIVER_MAILID"
```

## Running Tests

To run the tests, navigate to the project directory and run:

```shell
go test -v ./...
```


