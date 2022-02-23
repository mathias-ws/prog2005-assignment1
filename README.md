# Assignment 1
Project for the first assignment in PROG2005- Cloud Technologies 2022. This rest api is used
to fetch some information about a university together with some information about the country 
it is located in. This can be done based on the name, country or on the university
name in the neighbouring countries.

The api is using http://universities.hipolabs.com/ for retrieving the university information
and https://restcountries.com/ for retrieving the country information. Without access
to these apis this api will not work.

# Endpoints
The api has three endpoint:

    /unisearcher/v1/neighbourunis/
    /unisearcher/v1/uniinfo/
    /unisearcher/v1/diag

## Uniinfo
The uniinfo endpoint can be used to find information about a university with some country information.
You can search for all universities within a given country, just for the name or for both
the name and the country.

### Request
The uniinfo endpoint can be used either based on university name or country name or both.

#### Parameters:
`name` is the English name of the university, it can be partial or complete.

`country` is the country the university is in. The complete english name must be provided.

Example search:

    unisearcher/v1/uniinfo/?name=Molde&country=norway
    unisearcher/v1/uniinfo/?name=Norwegian%20University%20of%20Science%20and%20Technology
    unisearcher/v1/uniinfo/?country=norway

### Response
A response will have the content type set to `application/json`.

Status codes:
* 200: Everything is ok.
* 204: No university found based on the request.
* 400: Client side error, wrong limit/other.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

A response with the http code 200 will always be a list of universities even when only one is found.

Example body:

    [
    {
        "name": "Molde University College",
        "country": "Norway",
        "isocode": "NO",
        "webpages": [
            "http://www.himolde.no/"
        ],
        "languages": {
            "nno": "Norwegian Nynorsk",
            "nob": "Norwegian Bokmål",
            "smi": "Sami"
        },
        "map": "https://www.openstreetmap.org/relation/2978650"
    }
    ]

## Neighbourunis
This endpoint searches for universities with a similar name in the given country and in its
neighbouring countries. A list of json objects are returned. 

### Request
The neighbourunis endpoint can be used to search based on university name and country name.
An optional parameter `limit` can be used to limit the number of objects are returned.

#### Parameters:

`name` is the English name of the university, it can be partial or complete.

`country` is the country that the api uses as the basis for finding neighbouring countries. 
The complete english name must be provided.

`limit` is the number of results that are returned. This is an optional parameter and can be omitted.
If not set by the user all results are returned.

Example search:

    unisearcher/v1/neighbourunis/?name=science&country=norway&limit=1
    unisearcher/v1/neighbourunis/?name=Norwegian%20University%20of%20Science%20and%20Technology&country=norway


### Response
A response will have the content type set to `application/json`.

Status codes:
* 200: Everything is ok.
* 204: No university found based on the request.
* 400: Client side error, wrong limit/other.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

A response with the http code 200 will always be a list of universities even when only one is found.

Example body:

    [
	{
		"name": "Norwegian University of Science and Technology",
		"country": "Norway",
		"isocode": "NO",
		"webpages": [
			"http://www.ntnu.no/"
		],
		"languages": {
			"nno": "Norwegian Nynorsk",
			"nob": "Norwegian Bokmål",
			"smi": "Sami"
		},
		"map": "https://www.openstreetmap.org/relation/2978650"
	}
    ]

## Diag
Returns some information about the service. The information may include: information about the backend
apis, uptime and api version.

### Request
The request takes no parameters.

Example search:

    unisearcher/v1/diag

### Response

Status codes:
* 200: Everything is ok.
* 405: When using other methods than get.
* 500: Undefined server side error.
* 502: Unable to reach the backend apis.

Example body:

    {
	"universitiesapi": 200,
	"countriesapi": 200,
	"version": "v1",
	"uptime": "1983s"
    }

# Known bugs
* Because of a bug in the university info api using the country `Vietnam` can be a bit flacky because the api
registers as `Viet Nam`.
* When the country name in the university api is different from the common name in the
country api there might occur some edge cases where no results will be shown. 


# How to deploy
There are three ways to deploy this app. You can: manually build the project using `go build`, 
deploy it on heroku and use the Dockerfile to build it as an docker image.

## Go build
These steps require you to have go installed on the computer.

1: Clone the repo.

2: Enter the folder and run the command `go build main.go`.

3: Use `./main` to run the project.

## Heroku
First create an account and a project in Heroku according to their documentation.

1: Clone the repo

2: Enter the folder in the terminal and run `heroku login`

3: Then run `git remote rename heroku`

4: You can then push it to the heroku master using `git push heroku`. It will then be deployed onto heroku.

## Container (experimental)
The project has a Dockerfile that can be used to deploy the application. 
It can be built using `docker build` or using the gitlab pipeline. The only tested method is
using the gitlab pipeline. To use the pipeline you have to create you own gitlab repo with
available runners. The docker registry in the ci/cd yaml file must be changed to match a
docker registry that is available for you.

When the docker image is built it can be run with the exposed port 80.

# Extra

# Design choices
I chose to use parameters when searching instead of adding the search string(s) into the path.
This to make it clearer for the user to see what information is to be added where.
