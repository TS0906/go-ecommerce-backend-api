package controller

import (
	"github.com/TS0906/go-ecommerce-backend-api/internal/service"
	"github.com/TS0906/go-ecommerce-backend-api/pkg/response"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// uc user controller
// us user service

// controller -> service -> repo -> models -> db
func (uc *UserController) GetUserByID(c *gin.Context) {

	// c.JSON(http.StatusOK, gin.H{ // map string interface
	// 	"message": uc.userService.GetInfoUser(),
	// 	"users":   []string{"thinh", "thinh2"},
	// })

	// if err != nil {
	// 	return response.ErrorResponse(c, 20003, "Email is Invalid")
	// }
	response.SuccessResponse(c, 20001, []string{"thinh", "thinh2"})
}
