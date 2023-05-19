# Golang REST API

A simple RESTful API built in order to learn Golang.

The API is designed to allow consumers to perform RESTful operations on resources related to the 2023 Formula 1 season.

#
## How to Run


### Install Dependencies

```shell
go mod download
```

### Start the API

```shell
go run main.go
```

After running the above command, the API should now be listening at `http://localhost:8080`.
#
## Usage


`GET /drivers` - Fetch a list of all drivers

```shell
curl --location 'localhost:8080/drivers'

# Example Response
[
    {
        "name": "Lewis Hamilton",
        "number": 44,
        "team": "Mercedes AMG Petronas"
    },
    {
        "name": "George Russell",
        "number": 63,
        "team": "Mercedes AMG Petronas"
    },
    {
        "name": "Max Verstappen",
        "number": 1,
        "team": "Red Bull Racing"
    },
    {
        "name": "Sergio Perez",
        "number": 11,
        "team": "Red Bull Racing"
    },
    ...
]
```

`GET /drivers/:driverName` - Fetch a single driver by name

```shell
curl --location 'localhost:8080/drivers/Max%20Verstappen'

# Example Response
{
    "name": "Max Verstappen",
    "number": 1,
    "team": "Red Bull Racing"
}
```

#
### Last Updated

May 19th, 2023