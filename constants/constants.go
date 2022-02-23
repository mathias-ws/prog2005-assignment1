package constants

import "time"

// Program info
const PORT = "80"
const PROGRAM_VERSION = "v1"

// URL paths
const UNISEARCH_LOCATION = "/unisearcher/v1/uniinfo/"
const NEIGHBOUR_UNIS_LOCATION = "/unisearcher/v1/neighbourunis/"
const DIAG_LOCATION = "/unisearcher/v1/diag/"

// API urls
const COUNTRY_API = "https://restcountries.com/v3.1/name/"
const COUNTRY_API_CONSTRAINTS = "?fields=name,maps,languages,borders"
const COUNTRY_API_ROOT_URL = "https://restcountries.com/"
const COUNTRY_API_ALPHA_CODE = "https://restcountries.com/v3.1/alpha/?codes="
const UNIVERSITY_API = "http://universities.hipolabs.com/search?"
const UNIVERSITY_API_ROOT_URL = "http://universities.hipolabs.com"

// URL params
const URL_PARAM_NAME = "name"
const URL_PARAM_NAME_CONTAINS = "name_contains"
const URL_PARAM_COUNTRY = "country"
const URL_PARAM_AND = "&"
const URL_PARAM_EQUALS = "="
const URL_PARAM_LIMIT = "limit"

// Web web_client timeout time
const CLIENT_TIMEOUT = 30 * time.Second
