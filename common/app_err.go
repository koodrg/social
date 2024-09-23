package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Key        string `json:"key"`
	Log        string `json:"log"`
}

func NewErrorResponse(root error, msg string, key string, log string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Key:        key,
		Log:        log,
	}
}

func NewFullErrorResponse(statusCode int, root error, msg string, key string, log string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Key:        key,
		Log:        log,
	}
}

func NewUnauthorizedErrorResponse(root error, msg, key string, log string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Key:        key,
		Log:        log,
	}
}

func NewCustomError(root error, msg string, key string, log string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, key, log)
	}

	return NewErrorResponse(errors.New(msg), msg, key, log)
}

func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

func (e *AppError) Error() string {
	return e.RootError().Error()
}

func ErrDB(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}

func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "Invalid request", err.Error(), "INVALID_REQUEST_ERROR")
}

func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Internal error", err.Error(), "INTERNAL_ERROR")
}

//todo: cannot list, cannot create, cannot delete, cannot update

func ErrPermissionDenied(err error) *AppError {
	return NewCustomError(err, "Permission denied", err.Error(), "PERMISSION_DENIED_ERROR")
}

func ErrCanNotCreateEntity(tableName string, err error) *AppError {
	return NewCustomError(err, "Can not create "+tableName, err.Error(), "CANNOT_CREATE_ENTITY_ERROR")
}

var RecordNotFound = errors.New("Record not found")
