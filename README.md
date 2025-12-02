# Nebula Challenge

A Go CLI tool that interacts with the SSL Labs API to analyze the TLS security of domains and IP addresses.  
It includes a background daemon that stores assessments, updates them and communicates with the CLI using Unix domain sockets.

## Requirements

- Go ≥ 1.25
- cobra
- godotenv
- testify

## Installation

```bash
go build
go install
```
## Usage

To execute the CLI application use the following command
``` bash
nebula-challenge
```

### Start the daemon
The daemon stores assessments and exposes an internal API over Unix domain sockets.
``` bash
nebula-challenge serve
```

### Get SSL Labs server information
You can find information about the SSL labs server using the info command.
``` bash
nebula-challenge info
```

### Analyze a domain or IP address
To check a domain or ip address use the analyze command, you can specify the domain as an argument using the flag -d or --domain
``` bash
nebula-challenge analyze -d www.example.com
nebula-challenge analyze --domain www.example.com
```

To run in the background
``` bash
nebula-challenge analyze -d example.com -p
```
Save the results specifying a key
``` bash
nebula-challenge analyze -d example.com -k my-key
```

### Print stored assessments
List all stored assessments, bear in mind that every time you execute this command it polls the assessments in status not ready, to avoid overcharging the SSL Labs server it only does so if at least 15 seconds have passed since the last poll
``` bash
nebula-challenge print
```

Print details of a specific key
``` bash
nebula-challenge print -k my-key
```

### Command help
For more information of any command use the flag -h or --help
``` bash
nebula-challenge analyze -h
nebula-challenge print --help
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

- **internal/daemon**

    Background service that stores assessments and processes requests via Unix domain sockets, includes:
    
    - error definitions
    - request and responses models
    - a client implementing the requests for the unix server
    - a server implementing the logic for storing and updating the assessments

Analysis status flow
```
DNS -> IN_PROGRESS -> READY or ERROR
```
