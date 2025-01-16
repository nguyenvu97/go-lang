package response

type DecodeToken struct {
	Subject string `json:"subject"`
	ID string `json:"id"`
	ExpiresAt int64 `json:"expires_at"`
	IssuedAt int64 `json:"issued_at"`
}
