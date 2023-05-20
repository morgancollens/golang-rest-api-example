# Golang REST API

A simple RESTful API built in order to learn Golang.

The API is designed to allow consumers to perform RESTful operations on resources related to the 2023 Formula 1 season, and is powered by [Ergast Developer API](https://ergast.com/mrd/)

#
## How to Run


### Install Dependencies

```shell
go mod download
```

### Start the API

```shell
go run .
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
        "driverId": "albon",
        "permanentNumber": "23",
        "code": "ALB",
        "url": "http://en.wikipedia.org/wiki/Alexander_Albon",
        "givenName": "Alexander",
        "familyName": "Albon",
        "dateOfBirth": "1996-03-23",
        "nationality": "Thai"
    },
    {
        "driverId": "alonso",
        "permanentNumber": "14",
        "code": "ALO",
        "url": "http://en.wikipedia.org/wiki/Fernando_Alonso",
        "givenName": "Fernando",
        "familyName": "Alonso",
        "dateOfBirth": "1981-07-29",
        "nationality": "Spanish"
    },
    {
        "driverId": "bottas",
        "permanentNumber": "77",
        "code": "BOT",
        "url": "http://en.wikipedia.org/wiki/Valtteri_Bottas",
        "givenName": "Valtteri",
        "familyName": "Bottas",
        "dateOfBirth": "1989-08-28",
        "nationality": "Finnish"
    },
    {
        "driverId": "de_vries",
        "permanentNumber": "21",
        "code": "DEV",
        "url": "http://en.wikipedia.org/wiki/Nyck_de_Vries",
        "givenName": "Nyck",
        "familyName": "de Vries",
        "dateOfBirth": "1995-02-06",
        "nationality": "Dutch"
    },
    ...
]
```

`GET /drivers/:driverName` - Fetch a single driver by name

```shell
curl --location 'localhost:8080/drivers/Max%20Verstappen'

# Example Response
{
    "driverId": "max_verstappen",
    "permanentNumber": "33",
    "code": "VER",
    "url": "http://en.wikipedia.org/wiki/Max_Verstappen",
    "givenName": "Max",
    "familyName": "Verstappen",
    "dateOfBirth": "1997-09-30",
    "nationality": "Dutch"
}
```

#
### Last Updated

May 19th, 2023