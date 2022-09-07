package error




//TO DO >>> can i ceate this as model use as :-- repo -> service -> handler
type Error struct{
	code int
	message string 
}

func NewError(code int , message string )*Error{
	return &Error{code: code , message: message}
}
func (e *Error) Error() string {
	return e.message
}
func (e *Error)Code() int{
	return e.code
}
//New returns nil error 
func New(err string)*Error{
	return &Error{}
}