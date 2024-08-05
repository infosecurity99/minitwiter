package handler

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"test/api/models"
	"test/pkg/jwt"
	"time"
)

// UserLogin godoc
// @Router       /auth/user/login [POST]
// @Summary      User login
// @Description  admin login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        login body models.UserLoginRequest true "login"
// @Success      201  {object}  models.UserLoginResponse
// @Failure      400  {object}  models.Response
// @Failure      404  {object}  models.Response
// @Failure      500  {object}  models.Response
func (h Handler) UserLogin(c *gin.Context) {
	userLogin := models.UserLoginRequest{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := c.ShouldBindJSON(&userLogin); err != nil {
		handleResponse(c, h.log, "error while reading body", http.StatusBadRequest, err.Error())
		return
	}

	loginResponse, err := h.services.AuthService().UserLogin(ctx, userLogin)
	if err != nil {
		handleResponse(c, h.log, "error while admin login", http.StatusBadRequest, err.Error())
		return
	}

	handleResponse(c, h.log, "success", http.StatusOK, loginResponse)
}

func getAuthInfo(c *gin.Context) (models.AuthInfo, error) {
	accessToken := c.GetHeader("Authorization")
	if accessToken == "" {
		return models.AuthInfo{}, errors.New("unauthorized")
	}

	m, err := jwt.ExtractClaims(accessToken)
	if err != nil {
		return models.AuthInfo{}, err
	}

	return models.AuthInfo{
		UserID:   m["user_id"].(string),
		UserRole: m["user_role"].(string),
	}, nil
}
