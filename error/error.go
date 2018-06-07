package error

// XCPCTypeError indicate the message is not a valid xcpc message
type XCPCTypeError struct {
	Msg string
}

func (e XCPCTypeError) Error() string {
	return e.Msg
}
