# Assignment 1
Project for the first assignment in PROG2005- Cloud Technologies 2022.

# Endpoints
The api has three endpoint:

    /unisearcher/v1/neighbourunis/
    /unisearcher/v1/uniinfo/
    /unisearcher/v1/diag

## Uniinfo
The uniinfo endpoint can be used to find information about a university with some country information.
You can search for all universities within a given country, just for the name or for both.

### Request
The uniinfo endpoint can be used either based on university name or country name or both.

####Parameters:
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
* 400: Client side error, wrong limit/other.
* 404: No university found based on the request.
* 405: When using other methods than get.

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

##Neighbourunis
This endpoint searches for universities with a similar name in the given country and in its
neighbouring countries. A list of json objects are returned. 

### Request
The neighbourunis endpoint can be used to search based on university name and country name.
An optional parameter `limit` can be used to limit the number of objects are returned.

####Parameters:

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
* 400: Client side error, wrong limit/other.
* 404: No university found based on the request.
* 405: When using other methods than get.

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

###Request
The request takes no parameters.

Example search:

    unisearcher/v1/diag

###Response

Status codes:
* 200: Everything is ok.
* 405: When using other methods than get.

Example body:

    {
	"universitiesapi": 200,
	"countriesapi": 200,
	"version": "v1",
	"uptime": "1983s"
    }

# Todo
* uniinfo: contains: https://github.com/Hipo/university-domains-list-api/blob/master/app.py#L28
* Documentation
* Correct endpoints search method?
* Error handling?
* Code clean up
* Comments?
* Heroku
* Limit uniinfo endpoint
* Error handling when the backend apis are not up