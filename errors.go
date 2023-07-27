package std

type FieldError struct {
	Field   string
	Message string
	Key     string
	Params  []interface{}
}

type Error struct {
	Message     string
	Key         string
	Args        []interface{}
	Err         error
	FieldErrors []FieldError
	ErrorCode   string
	HttpStatus  int
}

func NewError(msg, key string, args ...interface{}) *Error {
	return &Error{
		Message: msg,
		Key:     key,
		Args:    args,
	}
}
func Err(msg string) *Error {
	return &Error{
		Message: msg,
	}
}

func WrapError(err error, key string, args ...interface{}) *Error {
	if e, ok := err.(*Error); ok {
		return e
	}
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	return &Error{
		Message: msg,
		Err:     err,
		Key:     key,
		Args:    args,
	}
}

func (s *Error) Error() string {
	return s.Message
}

func (s *Error) AddFieldError(field, message, key string, params ...interface{}) *Error {
	s.FieldErrors = append(s.FieldErrors, FieldError{Field: field, Message: message, Key: key, Params: params})
	return s
}
func (s *Error) AddField(field, message string) *Error {
	s.FieldErrors = append(s.FieldErrors, FieldError{Field: field, Message: message})
	return s
}
func (s *Error) I18n(key string, args ...interface{}) *Error {
	s.Key = key
	s.Args = args
	return s
}

func (s *Error) SetError(err error) *Error {
	s.Err = err
	return s
}
func (s *Error) SetErrorCode(code string) *Error {
	s.ErrorCode = code
	return s
}
func (s *Error) SetHttpStatus(status int) *Error {
	s.HttpStatus = status
	return s
}

func (s *Error) ToResult(translator func(key string, params ...interface{}) string) Result {
	if s == nil {
		return Result{}
	}
	fes := make(map[string]string, len(s.FieldErrors))
	for _, v := range s.FieldErrors {
		if _, ok := fes[v.Field]; ok {
			continue
		}
		if v.Key != "" {
			fes[v.Field] = translator(v.Key, v.Params...)
		} else {
			fes[v.Field] = v.Message
		}
	}
	errStr := s.Message
	if s.Key != "" {
		errStr = translator(s.Key, s.Args...)
	}
	r := Result{
		State:       1,
		Error:       errStr,
		ErrorCode:   s.ErrorCode,
		FieldErrors: fes,
	}

	return r
}
