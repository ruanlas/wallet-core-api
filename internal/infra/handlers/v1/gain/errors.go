package gain

type InvalidArgs struct {
	message string
}

func (invalidArgs *InvalidArgs) Error() string {
	return invalidArgs.message
}
