package utils

import (
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware"
	"net/http"
)

const (
	InvalidParameter      = "InvalidParameter"
	MissingParameter      = "MissingParameter"
	NoResultsForParameter = "NoResultsForParameter"
	ServerError           = "ServerError"
)

var HTTPErrorMessages = map[string]string{
	InvalidParameter:      "invalid parameter '%s'",
	MissingParameter:      "missing parameter '%s'",
	NoResultsForParameter: "the parameter '%s' with id '%s' returned no results",
	ServerError:           "internal server error",
}

type InvalidParameterError struct {
	ParameterName string
}

func (e *InvalidParameterError) Error() string {
	return fmt.Sprintf(HTTPErrorMessages[InvalidParameter], e.ParameterName)
}

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return fmt.Sprintf(HTTPErrorMessages[MissingParameter], e.ParameterName)
}

type NoResultsForParameterError struct {
	ParameterName string
	ID            string
	StatusCode    int
}

func (e *NoResultsForParameterError) Error() string {
	return fmt.Sprintf(HTTPErrorMessages[NoResultsForParameter], e.ParameterName, e.ID)
}

type ServerErrorType struct{}

func (e *ServerErrorType) Error() string {
	return HTTPErrorMessages[ServerError]
}

func HandleHTTPError(w http.ResponseWriter, err error, defaultStatusCode int) {
	statusCode := defaultStatusCode
	if httpError, ok := err.(*NoResultsForParameterError); ok {
		statusCode = httpError.StatusCode
	}
	middleware.ErrorLogger.Printf("%v", err.Error())
	http.Error(w, err.Error(), statusCode)
}
