package errors

func (e CheckerError) Error() string {
	return e.Message
}
