package response

const (
	ErrorSuccess          = 20001 // ok
	ErrorCodeParamInvalid = 20002 // Invalid
	ErrorTokenInvalid     = 20003 // Invalid

	//register code
	ErrorCodeUserHasExit = 5001
	ErrorCodeBodyInvalid = 5002
	ErrorUserNotFound = 4004

	//JwtData_code
	ErrInvalidKey= 6001
	ErrInvalidKeyType= 6002
	ErrHashUnavailable= 6003
	ErrTokenMalformed= 6004
	ErrTokenUnverifiable= 6005
	ErrTokenSignatureInvalid= 6006
	ErrTokenRequiredClaimMissing= 6007
	ErrTokenInvalidAudience= 6008
	ErrTokenExpired= 6009
	ErrTokenUsedBeforeIssued= 6010
	ErrTokenInvalidIssuer= 6011
	ErrTokenInvalidSubject= 6012
	ErrTokenInvalidId= 6013
	ErrTokenInvalidClaims= 6014
	ErrInvalidType= 6015
	ErrProductNotBody = 7001

	// Redis Error
	ErrKeyNil= 9000
	ErrValueNil= 9001
	// request max
	ErrRequestFullSize = 8888
)

var msg = map[int]string{
	ErrorSuccess:          "success",
	ErrorCodeParamInvalid: "email is  invalid",
	ErrorTokenInvalid:     "token is invalid",
	ErrorCodeUserHasExit:  "email has in db ",
	ErrorCodeBodyInvalid:  "body invalid",
	ErrorUserNotFound: "user Not Found",
	ErrInvalidKey                : "key is invalid",
	ErrInvalidKeyType            : "key is of invalid type",
	ErrHashUnavailable           : "the requested hash function is unavailable",
	ErrTokenMalformed            : "token is malformed",
	ErrTokenUnverifiable         : "token is unverifiable",
	ErrTokenSignatureInvalid     : "token signature is invalid",
	ErrTokenRequiredClaimMissing : "token is missing required claim",
	ErrTokenInvalidAudience      : "token has invalid audience",
	ErrTokenExpired              : "token is expired",
	ErrTokenUsedBeforeIssued     : "token used before issued",
	ErrTokenInvalidIssuer        : "token has invalid issuer",
	ErrTokenInvalidSubject       : "token has invalid subject",
	ErrTokenInvalidId            : "token has invalid id",
	ErrTokenInvalidClaims        : "token has invalid claims",
	ErrInvalidType               : "invalid type for claim",
	ErrProductNotBody : "product request not data",
	ErrRequestFullSize : "request full many",
	ErrKeyNil : "product request not data",
	ErrValueNil : "request full many",
}
