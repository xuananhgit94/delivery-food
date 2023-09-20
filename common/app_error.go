package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
}

func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

func NewFullErrorRequest(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{statusCode, root, msg, log, key}
}

func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{StatusCode: http.StatusUnauthorized, RootErr: root, Message: msg, Key: key}
}

func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}
	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

func (e *AppError) RootError() error {
	var err *AppError
	switch {
	case errors.As(e.RootErr, &err):
		return err.RootError()
	default:
		return e.RootErr
	}
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorRequest(http.StatusInternalServerError, err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorRequest(http.StatusInternalServerError, err, "something went wrong in the server", err.Error(), "ErrInternal")
}
func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot list %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotList%s", entity))
}

func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotDelete%s", entity))
}

func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s deleted", strings.ToLower(entity)), fmt.Sprintf("Err%sDeleted", entity))
}

func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s already exists", strings.ToLower(entity)), fmt.Sprintf("Err%sAlreadyExists", entity))
}

func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("%s not found", strings.ToLower(entity)), fmt.Sprintf("Err%sNotFound", entity))
}

func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)), fmt.Sprintf("ErrCannotCreate%s", entity))
}

func ErrNoPermission(err error) *AppError {
	return NewCustomError(err, fmt.Sprintf("You have no permission"), fmt.Sprintf("ErroNoPermission"))
}

var RecordNotFound = errors.New("record not found")
