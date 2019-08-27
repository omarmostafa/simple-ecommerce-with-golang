package Errors

type NotAuthorizedToken struct {
	Message string
}

func (err NotAuthorizedToken)Error() string  {
	return  err.Message
}