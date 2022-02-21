package customErrors

import "errors"

// GetUnableToReachBackendApisError returns the error message for the client side error from the apis.
func GetUnableToReachBackendApisError() error {
	return errors.New("error sending request or getting response from the api")
}

// GetUnableToGetCountryError returns the error message for when no valid countries are found.
func GetUnableToGetCountryError() error {
	return errors.New("unable to retrieve country")
}

// GetNoContentFoundError returns the error message for when the given country was not found.
func GetNoContentFoundError() error {
	return errors.New("no universities with the given search parameters were not found")
}

//GetParameterError returns error message for when the user has not provided the mandatory parameters.
func GetParameterError() error {
	return errors.New("missing parameters or wrong parameters")
}

//GetInvalidLimitError returns error message for when the user has not used a proper limit.
func GetInvalidLimitError() error {
	return errors.New("invalid limit")
}
