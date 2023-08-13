package errors

type ErrorType struct {
	t string
}

var (
	ErrorTypeUnknown        = ErrorType{"unknown"}
	ErrorTypeAuthorization  = ErrorType{"authorization"}
	ErrorTypeIncorrectInput = ErrorType{"incorrect-input"}
	ErrorTypeNotFound       = ErrorType{"not-found"}
)

type SlugError struct {
	slug      string
	error     string
	errorType ErrorType
}

func (s SlugError) Slug() string {
	return s.slug
}

func (s SlugError) Error() string {
	return s.error
}

func (s SlugError) ErrorType() ErrorType {
	return s.errorType
}

func NewSlugError(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeUnknown,
	}
}

func NewAuthorizationError(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeAuthorization,
	}
}

func NewIncorrectInputError(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeIncorrectInput,
	}
}

func NewNotFoundError(slug, err string) SlugError {
	return SlugError{
		slug:      slug,
		error:     err,
		errorType: ErrorTypeNotFound,
	}
}
