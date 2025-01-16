package response

type ResponseJwt struct {
	Token	string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}