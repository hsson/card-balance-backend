[![Build Status](https://travis-ci.org/hsson/card-balance-backend.svg?branch=master)](https://travis-ci.org/hsson/card-balance-backend)
[![Go Report Card](https://goreportcard.com/badge/hsson/card-balance-backend)](https://goreportcard.com/report/hsson/card-balance-backend)

# card-balance-backend

The backend for the [Chalmers Card Balance](https://play.google.com/store/apps/details?id=se.creotec.chscardbalance2) application. It replaces the old closed-source implementation (which will become available when fully deprecated).

## Building
This implementation both supports Google AppEngine and native standalone runtimes. The API is implemented for Go v1.7+, with some fixes for Google AppEngine (which is currently running Go v1.6).

### Natively
To run the backend natively, simply run:
```
> go run main.go
```
And to build it:
```
> go build 
```

### Google AppEngine
To use Google AppEngine, first make sure you have the AppEngine SDK correctly installed. To start a test server, run:
```
> cd appengine
> dev_appserver.py app.yaml
```
And to deploy it to GoogleAppEngine:
```
> cd appengine
> gcloud app deploy -v <version> app.yaml
```
The backend is then available at `https://<version>-dot-<project-id>.appspot.com`.

## API Specification
### ```HTTP GET /balance/<card number> ```
Get information about a specified chalmers card. An example request could look like:
```
HTTP GET /balance/2222333344445555
```
If the request is successful, the response will look like:
``` json
{
    "success": true,
    "data": {
        "card_number": "2222333344445555",
        "full_name": "Emilia Emilsson",
        "email": "emilia1337@example.com",
        "balance": 69.42
    },
    "error": ""
}
```

If the request was not successful, the response could look like:
``` json
{
    "success": false,
    "data": null,
    "error": "Invalid card number"
}
```
There are several different errors that can occur.

### ```HTTP GET  /menu/<lang>```
Get today's menu from restraurants on Chalmer's campuses. Make sure to specify which langugage to get the menu in. The following different language parameters are valid:
- ```sv```: Swedish
- ```en```: English

An example request could look like:
```
HTTP GET /menu/en
```
If the request is successful, it could look like:
``` json
{
    "success": true,
    "data": {
        "language": "en",
        "menu": [
            {
                "name": "Kårrestaurangen",
                "image_url": "",
                "dishes": [
                    {
                        "title": "Classic Vegetarisk",
                        "desc": "Some nice food"
                    },
                    {
                        "title": "Classic Fisk",
                        "desc": "Another nice food"
                    }
                ]
            },
            {
                "name": "Ls Kitchen",
                "image_url": "",
                "dishes": [
                    {
                        "title": "Meat",
                        "desc": "Food"
                    }
                ]
            }
        ]
    },
    "error": ""
}
```
Note that there can be an arbitrary amount of restaurants returned, and each restaurant can have an arbitrary amount of dishes. It all depends on which day it is.

### ```HTTP GET /charge```
Use this endpoint to get redirected to the current website on which you can charge your Chalmers card.