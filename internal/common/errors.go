package common

import (
	"github.com/pkg/errors"
)

/************/
/*	Errors	*/
/************/

// ErrInternalError is returned if an unexpected error is raised.
type ErrInternalError struct{ error }

// ErrInvalidEntity is returned if an entity is not valid.
type ErrInvalidEntity struct{ error }

/********************/
/*	Constructors	*/
/********************/

// NewErrInternalError creates a new ErrInternalError.
func NewErrInternalError(err error, msg string, args ...interface{}) ErrInternalError {
	return ErrInternalError{errors.Wrapf(err, msg, args...)}
}

// NewErrInternalErrorMsg creates a new ErrInternalError without wrapping an existing error.
func NewErrInternalErrorMsg(msg string, args ...interface{}) ErrInternalError {
	return ErrInternalError{errors.Errorf(msg, args...)}
}

// NewErrInvalidEntity creates a new ErrInvalidEntity.
func NewErrInvalidEntity(structname string, reason string) ErrInvalidEntity {
	return ErrInvalidEntity{errors.Errorf("invalid %v: %v", structname, reason)}
}
