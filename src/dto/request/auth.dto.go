package request

type Authen struct {
	Email	string `json:"email" binding:"required"`
	PassWord string `json:"password" binding:"required"`
}