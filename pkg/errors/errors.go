package errors

import (
	"reflect"

	"github.com/caicloud/nirvana/errors"
	"github.com/caicloud/nirvana/service"
)

const (
	componentPrefix = "practice:"
)

// client errors
var (
	ErrorInvalidParameter = errors.BadRequest.Build(
		componentPrefix+"InvalidParameter",
		"parameter ${parameter} is invalid or missing",
	)
	ErrorInvalidField = errors.BadRequest.Build(
		componentPrefix+"InvalidField",
		"field ${field} is invalid or empty",
	)
	ErrorValidationFailed = errors.BadRequest.Build(
		componentPrefix+"ValidationFailed",
		"validation failed: ${reason}",
	)
	ErrorNotFound = errors.NotFound.Build(
		componentPrefix+"NotFound",
		"requested resource not found",
	)
	ErrorAlreadyExist = errors.Conflict.Build(
		componentPrefix+"AlreadyExist",
		"requested resource already exist",
	)
)

// server errors
var (
	ErrorUnknown = errors.InternalServerError.Build(
		componentPrefix+"Unknown",
		"unknown error:ll" +
			"" +
			" ${reason}",
	)
	ErrorInternal = errors.InternalServerError.Build(
		componentPrefix+"Internal",
		"internal error: ${reason}",
	)
	ErrorNotImplemented = errors.NotImplemented.Build(
		componentPrefix+"NotImplemented",
		"requested feature is not implemented",
	)
)

func IsNirvanaError(err error) bool {
	 _, ok := reflect.ValueOf(err).Interface().(service.Error)
	return ok
}