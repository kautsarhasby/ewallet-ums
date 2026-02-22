package api

import (
	"kautsarhasby/ewallet-ums/constants"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RegisterHandler struct {
	RegisterService interfaces.IRegisterService
}

func (api *RegisterHandler) Register(c *gin.Context) {
	var (
		log = helpers.Logger
	)
	request := models.User{}

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

	response, err := api.RegisterService.Register(c.Request.Context(), request)
	if err != nil {
		log.Error("failed to register new user", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, response)
	return
}
