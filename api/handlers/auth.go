package handlers

import (
	"api_gateway/api/http"
	"api_gateway/genproto/auth_service"
	"errors"

	"github.com/gin-gonic/gin"
)

// Login godoc
// @ID login
// @Router /login [POST]
// @Summary Login
// @Description Login
// @Tags Authentication	
// @Accept json
// @Produce json
// @Param login body auth_service.LoginRequest true "LoginRequestBody"
// @Success 201 {object} http.Response{data=auth_service.TokenResponse} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Login(c *gin.Context) {
	var login auth_service.LoginRequest

	err := c.ShouldBindJSON(&login)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	resp, err := h.services.AuthService().Login(
		c.Request.Context(),
		&login,
	)

	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}

// Register godoc
// @ID register
// @Router /register [POST]
// @Summary Register
// @Description Register
// @Tags Authentication
// @Accept json
// @Produce json
// @Param register body auth_service.CreateUser true "CreateUserRequest"
// @Success 201 {object} http.Response{data=auth_service.User} "User data"
// @Response 400 {object} http.Response{data=string} "Bad Request"
// @Failure 500 {object} http.Response{data=string} "Server Error"
func (h *Handler) Register(c *gin.Context) {

	var register auth_service.CreateUser

	err := c.ShouldBindJSON(&register)
	if err != nil {
		h.handleResponse(c, http.BadRequest, err.Error())
		return
	}

	res, err := h.services.UserService().CheckUser(
		c.Request.Context(),
		&auth_service.CheckUserRequest{
			Name:   register.Name,
			Secret: register.Secret,
		},
	)
	if err != nil {
		h.handleResponse(c, http.GRPCError, err.Error())
		return
	}

	if res.Exists && res.Registered {
		h.handleResponse(c, http.BadRequest, errors.New("user registered").Error())
		return
	}

	resp, err := h.services.UserService().Create(
		c.Request.Context(),
		&register,
	)

	if err != nil {
		h.handleResponse(c, http.InternalServerError, err.Error())
		return
	}

	h.handleResponse(c, http.Created, resp)
}
