package controller

import (
	"awesomeProject/pkg/response"
	"awesomeProject/src/dto/request"
	dto "awesomeProject/src/dto/response"
	"awesomeProject/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService service.IUserService
}

func NewUserController(userService service.IUserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// ShowAccount userRegister
// @Summary      Show an account
// @Description  register
// @Tags         accounts
// @Accept       json
// @Produce      json
//@Param        payload body response.UserDto true "payload"
// @Success      200  {object}  response.Response
// @Router       /api/v1/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	var user dto.UserDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := uc.userService.Register(user)
	if err >= 1 {
		response.SuccessResponse(c, 20001, "")
		return
	}
	response.ErrorResponse(c, 5001)
	return

}

// ShowAccount getUser
// @Summary      Show an account
// @Description  register
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        email   path    string  true  "Email of the user"
// @Success      200  {object}  response.Response
// @Router       /api/v1/user/getUser [post]
func (uc *UserController) GetUser(c *gin.Context) {

	email := c.Param("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is required"})
		return
	}

	user := uc.userService.FindByEmail(email)
	if user.ID != 0 {
		response.SuccessResponse(c, 20001, user)
		return
	}
	response.ErrorResponse(c, 4004)
	return

}
// ShowAccount getUser
// @Summary      Show an account
// @Description  register
// @Tags         accounts
// @Accept       json
// @Produce      json
//@Param        payload   body    request.Authen  true  "body of the user"
// @Success      200  {object}  response.Response
// @Router       /api/v1/user/login [post]

func (uc *UserController) Login(c *gin.Context) {
	var auth request.Authen

	if err := c.ShouldBindJSON(&auth); err != nil {
		response.SuccessResponse(c, 6002, nil)
		return
	}

	data := uc.userService.Login(auth)
	if data.Token == "" || data.RefreshToken == "" {
		response.SuccessResponse(c, 6003, nil)
		return
	}

	response.SuccessResponse(c, 20001, data)
	return

}
// ShowAccount getUser
// @Summary      Show an account
// @Description  register
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param Authorization header string true "Bearer token"
// @Success      200  {object}  response.Response
// @Router       /api/v1/user/login [post]
func (uc *UserController) DecodeToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		response.SuccessResponse(c, 6004, nil)
		return
	}
	token := ""
	if len(authHeader) > 7 {
		token = authHeader[7:]
	} else {
		response.SuccessResponse(c, 6005, nil)
		return
	}
	decodedToken := uc.userService.DecodeToken(token)
	if decodedToken.ID != "" {
		response.SuccessResponse(c, 20001, decodedToken)
		return
	}
	response.SuccessResponse(c, 6005, nil)
	return

}
