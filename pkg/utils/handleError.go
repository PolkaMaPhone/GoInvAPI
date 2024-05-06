package utils

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
)

const (
	InvalidParameter      = "InvalidParameter"
	MissingParameter      = "MissingParameter"
	NoResultsForParameter = "NoResultsForParameter"
	ServerError           = "ServerError"
	InvalidRoute          = "InvalidRoute"
	MethodNotAllowed      = "MethodNotAllowed"
)

var HTTPErrorMessages = map[string]string{
	InvalidParameter:      "invalid parameter '%s'",
	MissingParameter:      "missing parameter '%s'",
	NoResultsForParameter: "the parameter '%s' with id '%s' returned no results",
	ServerError:           "internal server error",
	InvalidRoute:          "invalid route '%s'",
	MethodNotAllowed:      "method '%s' is not allowed for route '%s'",
}

type InvalidRouteError struct {
	Route string
}

func (e *InvalidRouteError) Error() string {
	return InvalidRoute
}

type InvalidParameterError struct {
	ParameterName string
}

func (e *InvalidParameterError) Error() string {
	return InvalidParameter
}

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return MissingParameter
}

type NoResultsForParameterError struct {
	ParameterName string
	ID            string
	StatusCode    int
}

func (e *NoResultsForParameterError) Error() string {
	return NoResultsForParameter
}

type ServerErrorType struct{}

func (e *ServerErrorType) Error() string {
	return ServerError
}

type MethodNotAllowedError struct {
	Method string
	Route  string
}

func (e *MethodNotAllowedError) Error() string {
	return MethodNotAllowed
}

func HandleHTTPError(w http.ResponseWriter, err error) {
	var statusCode int
	var errMsg string

	// type assertion is ok here because the default case won't blow up and cause a panic
	//goland:noinspection GoTypeAssertionOnErrors
	switch e := err.(type) {
	case *InvalidRouteError:
		statusCode = http.StatusNotFound
		errMsg = fmt.Sprintf(HTTPErrorMessages[e.Error()], e.Route)
	case *MethodNotAllowedError:
		statusCode = http.StatusMethodNotAllowed
		errMsg = fmt.Sprintf(HTTPErrorMessages[e.Error()], e.Method, e.Route)
	case *NoResultsForParameterError:
		statusCode = http.StatusNotFound
		errMsg = fmt.Sprintf(HTTPErrorMessages[e.Error()], e.ParameterName, e.ID)
	case *InvalidParameterError:
		statusCode = http.StatusBadRequest
		errMsg = fmt.Sprintf(HTTPErrorMessages[e.Error()], e.ParameterName)
	case *ServerErrorType:
		statusCode = http.StatusInternalServerError
		errMsg = HTTPErrorMessages[ServerError]
	default:
		statusCode = http.StatusInternalServerError
		errMsg = HTTPErrorMessages[ServerError]
	}

	logging.ErrorLogger.Printf("%v", errMsg)
	http.Error(w, errMsg, statusCode)
}

func HandleGetByIDErrors(w http.ResponseWriter, err error, object interface{}, id int32, objectType string) bool {
	if err != nil {
		// Check if the error is a pgx.ErrNoRows error
		if errors.Is(err, pgx.ErrNoRows) {
			httpError := &NoResultsForParameterError{ParameterName: objectType, ID: strconv.Itoa(int(id))}
			HandleHTTPError(w, httpError)
		} else if errors.Is(err, &MethodNotAllowedError{}) {
			// If the error is a MethodNotAllowedError, return a 405 status code
			httpError := &MethodNotAllowedError{Method: "GET", Route: "/" + objectType + "/{id}"}
			HandleHTTPError(w, httpError)
		} else {
			// For all other errors, return a 500 status code and a generic server error message
			HandleHTTPError(w, &ServerErrorType{})
		}
		return true
	}

	if object == nil {
		httpError := &NoResultsForParameterError{ParameterName: objectType, ID: strconv.Itoa(int(id))}
		HandleHTTPError(w, httpError)
		return true
	}

	return false
}

func HandleGetAllErrors(w http.ResponseWriter, err error, objects interface{}, objectType string) bool {
	if err != nil {
		// Check if the error is a sql.ErrNoRows error
		if errors.Is(err, pgx.ErrNoRows) {
			httpError := &NoResultsForParameterError{ParameterName: objectType, ID: "all"}
			HandleHTTPError(w, httpError)
		} else {
			// For all other errors, return a 500 status code and a generic server error message
			HandleHTTPError(w, &ServerErrorType{})
		}
		return true
	}

	// Check if objects is a slice
	if objectsSlice, ok := objects.([]interface{}); ok {
		if objectsSlice == nil || len(objectsSlice) == 0 {
			httpError := &NoResultsForParameterError{ParameterName: objectType, ID: "all"}
			HandleHTTPError(w, httpError)
			return true
		}
	}

	return false
}
