package utils

import (
	"errors"
	"fmt"
	"github.com/PolkaMaPhone/GoInvAPI/pkg/middleware/logging"
	"github.com/jackc/pgx/v5"
	"net/http"
	"strconv"
)

type InvalidRouteError struct {
	Route string
}

func (e *InvalidRouteError) Error() string {
	return fmt.Sprintf("invalid route '%s'", e.Route)
}

type InvalidParameterError struct {
	ParameterName string
}

func (e *InvalidParameterError) Error() string {
	return fmt.Sprintf("invalid parameter '%s'", e.ParameterName)
}

type MissingParameterError struct {
	ParameterName string
}

func (e *MissingParameterError) Error() string {
	return fmt.Sprintf("missing parameter '%s'", e.ParameterName)
}

type NoResultsForParameterError struct {
	ParameterName string
	ID            string
}

func (e *NoResultsForParameterError) Error() string {
	return fmt.Sprintf("the parameter '%s' with id '%s' returned no results", e.ParameterName, e.ID)
}

type ServerErrorType struct{}

func (e *ServerErrorType) Error() string {
	return "internal server error"
}

type MethodNotAllowedError struct {
	Method string
	Route  string
}

func (e *MethodNotAllowedError) Error() string {
	return fmt.Sprintf("method '%s' is not allowed for route '%s'", e.Method, e.Route)
}

func HandleGetIDFromRequestErrors(w http.ResponseWriter, err error, id string) error {
	if err != nil {
		logging.ErrorLogger.Printf("Error getting ID from request: %v", err)

		return &InvalidParameterError{ParameterName: id}
	}
	return nil
}

type JSONResponseError struct {
	Err error
}

func (e *JSONResponseError) Error() string {
	return "error responding with JSON: " + e.Err.Error()
}

func HandleRespondWithJSONErrors(w http.ResponseWriter, err error) {
	if err != nil {
		logging.ErrorLogger.Printf("Error responding with JSON: %v", err)
		w.Header().Set("X-Error-Details", "Error responding with JSON")
	}
}

func HandleHTTPError(w http.ResponseWriter, err error) {
	var statusCode int
	var errMsg string
	var errCode string

	switch e := err.(type) {
	case *InvalidRouteError:
		statusCode = http.StatusNotFound
		errMsg = e.Error()
		errCode = "INVALID_ROUTE"
	case *MethodNotAllowedError:
		statusCode = http.StatusMethodNotAllowed
		errMsg = e.Error()
		errCode = "METHOD_NOT_ALLOWED"
	case *NoResultsForParameterError:
		statusCode = http.StatusNotFound
		errMsg = e.Error()
		errCode = "NO_RESULTS"
	case *InvalidParameterError:
		statusCode = http.StatusBadRequest
		errMsg = e.Error()
		errCode = "INVALID_PARAMETER"
	case *ServerErrorType:
		statusCode = http.StatusInternalServerError
		errMsg = e.Error()
		errCode = "SERVER_ERROR"
	case *strconv.NumError:
		statusCode = http.StatusBadRequest
		errMsg = fmt.Sprintf("invalid parameter: cannot parse '%s' as %s", e.Num, e.Func)
		errCode = "INVALID_PARAMETER"
	default:
		statusCode = http.StatusInternalServerError
		errMsg = fmt.Errorf("unexpected error type %T: %v", e, e).Error()
		errCode = "UNKNOWN_ERROR"
	}

	logging.ErrorLogger.Printf("%v", errMsg)
	w.Header().Set("X-Error-Code", errCode)
	http.Error(w, errMsg, statusCode)
}

func HandleGetByIDErrors(w http.ResponseWriter, err error, object interface{}, id int32, objectType string) {
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			HandleHTTPError(w, &NoResultsForParameterError{ParameterName: objectType, ID: strconv.Itoa(int(id))})
		} else {
			HandleHTTPError(w, &ServerErrorType{})
		}
	} else if object == nil {
		HandleHTTPError(w, &NoResultsForParameterError{ParameterName: objectType, ID: strconv.Itoa(int(id))})
	}
}

func HandleGetAllErrors(w http.ResponseWriter, err error, objects interface{}, objectType string) {
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			HandleHTTPError(w, &NoResultsForParameterError{ParameterName: objectType, ID: "all"})
		} else {
			HandleHTTPError(w, &ServerErrorType{})
		}
	} else if objectsSlice, ok := objects.([]interface{}); ok {
		if objectsSlice == nil || len(objectsSlice) == 0 {
			HandleHTTPError(w, &NoResultsForParameterError{ParameterName: objectType, ID: "all"})
		}
	}
}
