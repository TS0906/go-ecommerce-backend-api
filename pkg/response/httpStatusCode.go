package response

// ResponseData struct

const (
	ErrCodeSuccess      = 20001 // Success
	ErrCodeParamInvalid = 20003 // Email is Invalid
	ErrInvalidToken     = 20004 // Invalid Token
)

//message
var msg = map[int]string{
	ErrCodeSuccess:      "Success",
	ErrCodeParamInvalid: "Email is Invalid",
	ErrInvalidToken:     "Invalid Token",
}
