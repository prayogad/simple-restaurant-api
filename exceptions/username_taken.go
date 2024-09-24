package exceptions

type UsernameTakenError struct {
	Error string
}

func NewUsernameTakenError(error string) UsernameTakenError {
	return UsernameTakenError{Error: error}
}
