package errors

import "errors"

var (
	ErrorInvalidId                     = errors.New("invalid id")
	ErrorInvalidHttpMethod             = errors.New("invalid http method")
	ErrorInvalidDuelistFieldsJson      = errors.New("invalid duelist fields json")
	ErrorQueryToCreateDuelistIsInvalid = errors.New("query to create duelist is invalid")
	ErrorUnableToCreateDuelist         = errors.New("unable to create duelist")
)
