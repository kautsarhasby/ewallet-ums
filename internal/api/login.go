package api

import (
	"kautsarhasby/ewallet-ums/constants"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginHandler struct {
	LoginService interfaces.ILoginService
}

func (api *LoginHandler) Login(c *gin.Context) {
	var (
		log      = helpers.Logger
		request  models.LoginRequest
		response models.LoginResponse
	)

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Error("failed to Parse request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedParseRequest, nil)
		return
	}

	if err := request.Validate(); err != nil {
		log.Error("failed to Validate request", err)
		helpers.SendResponseHTTP(c, http.StatusBadRequest, constants.ErrFailedValidateRequest, nil)
		return
	}

	response, err := api.LoginService.Login(c.Request.Context(), request)
	if err = request.Validate(); err != nil {
		log.Error("Failed to login", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, constants.SuccessMessage, response)

}
