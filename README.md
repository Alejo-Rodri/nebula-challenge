# Nebula Challenge

This is a CLI app in Go that uses the SSL labs API to check the TLS security of a given domain. 

## Requirements

- Go >= 1.25
- Cobra
- godotenv

## Installation

``` bash
go build
go install
```
## Usage

To execute the CLI application use the following command
``` bash
nebula-challenge
```

You can find information about the SSL labs server using the info command.
``` bash
nebula-challenge info
```

To check a domain or ip address use the analyze command, you can specify the domain as an argument using the flag -d or --domain
``` bash
nebula-challenge analyze -d www.ssllabs.com
nebula-challenge analyze --domain www.ssllabs.com
```

For more information of any command use the flag -h or --help
``` bash
nebula-challenge analyze -h
nebula-challenge analyze --help
```

## Configuration

The project has a `.env` file that only contains the url of the SSL labs API.

## Project Structure

- **cmd**

     Defines the CLI commands and injects the HTTP client.

- **configs**

    Loads environment variables.

- **internal/app**

    Defines the interface for the endpoints used by the CLI. This interface is used to keep the packages decoupled.

- **internal/infra/api**

    Implements the API requests, includes:

    - error definitons
    - response models
    - `http.Client` configuration
    - logic to wait between the status changes of the `analyze` endpoint

Analysis status flow
```
    DNS -> IN_PROGRESS -> READY or ERROR
```
