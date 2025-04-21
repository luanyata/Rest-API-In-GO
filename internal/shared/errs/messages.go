package errs

type ErrorMessage struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}

var UserError = struct {
	EmailInvalid       ErrorMessage
	EmailAlreadyExists ErrorMessage
	FailHashPassword   ErrorMessage
	NotFound           ErrorMessage
}{
	EmailInvalid:       ErrorMessage{"Invalid email address", "INVALID_EMAIL"},
	EmailAlreadyExists: ErrorMessage{"Email already exists", "EMAIL_ALREADY_EXISTS"},
	FailHashPassword:   ErrorMessage{"Failed to hash password", "HASH_PASSWORD_FAILED"},
	NotFound:           ErrorMessage{"User not found", "USER_NOT_FOUND"},
}
