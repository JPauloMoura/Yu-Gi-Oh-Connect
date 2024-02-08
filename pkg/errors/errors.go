package errors

import "errors"

var (
	ErrorInvalidId                        = errors.New("invalid id")
	ErrorInvalidHttpMethod                = errors.New("invalid http method")
	ErrorInvalidDuelistFieldsJson         = errors.New("invalid duelist fields json")
	ErrorQueryToCreateDuelistIsInvalid    = errors.New("query to create duelist is invalid")
	ErrorUnableToCreateDuelist            = errors.New("unable to create duelist")
	ErrorInvalidDate                      = errors.New("invalid date")
	ErrorDateMustBeLessThanTheCurrentDate = errors.New("date must be less than the current date")
	ErrorCepServiceIsUnavailable          = errors.New("cep service is unavailable")
	ErrorInvalidCep                       = errors.New("invalid cep")
	ErrorInvalidRequest                   = errors.New("invalid request")
	ErrorInvalidResponseBody              = errors.New("invalid response body")
	ErrorInvalidDateFormat                = errors.New("invalid date format, expected dd/mm/yyyy")
)

func Join(err error, join error) error {
	return errors.New(err.Error() + ": " + join.Error())
}
