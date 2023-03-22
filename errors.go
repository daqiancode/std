package std

type FieldError struct {
	Field  string
	Key    string
	Params []interface{}
}

var DefaultErrorStatusCode = 500

// Error i18n error
type Error struct {
	statusCode  int
	err         error
	key         string        // i18n key
	params      []interface{} // i18n params
	fieldErrors []FieldError  // request field errors
	errorCode   string
}

func NewError(key string, params ...interface{}) *Error {
	return &Error{
		key:    key,
		params: params,
	}
}
func WrapError(err error, key string, params ...interface{}) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	return &Error{
		statusCode: DefaultErrorStatusCode,
		err:        err,
		key:        key,
		params:     params,
	}
}

func (s *Error) Error() string {
	return s.key
}

func (s *Error) Key() string {
	return s.key
}

func (s *Error) Params() []interface{} {
	return s.params
}
func (s *Error) Cause() error {
	return s.err
}

func (s *Error) SetStatusCode(statusCode int) *Error {
	s.statusCode = statusCode
	return s
}
func (s *Error) StatusCode() int {
	return s.statusCode
}

func (s *Error) SetErrorCode(errorCode string) *Error {
	s.errorCode = errorCode
	return s
}
func (s *Error) ErrorCode() string {
	return s.errorCode
}

func (s *Error) FieldErrors() []FieldError {
	return s.fieldErrors
}

func (s *Error) AddFieldError(field, key string, params ...interface{}) *Error {
	s.fieldErrors = append(s.fieldErrors, FieldError{Field: field, Key: key, Params: params})
	return s
}

func (s *Error) AppendFieldError(fe FieldError) *Error {
	s.fieldErrors = append(s.fieldErrors, fe)
	return s
}

func (s *Error) ToResult(translator func(key string, params ...interface{}) string) Result {
	if s == nil {
		return Result{}
	}
	fes := make(map[string]string, len(s.fieldErrors))
	for _, v := range s.fieldErrors {
		if _, ok := fes[v.Field]; ok {
			continue
		}
		fes[v.Field] = translator(v.Key, v.Params...)
	}
	r := Result{
		State:       1,
		Error:       translator(s.key, s.params...),
		ErrorCode:   s.errorCode,
		FieldErrors: fes,
	}

	return r
}
