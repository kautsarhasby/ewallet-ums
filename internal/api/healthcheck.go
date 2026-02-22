package api

import (
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthCheckHandler struct {
	HealthCheckService interfaces.IHealthCheckServices
}

func (api *HealthCheckHandler) HealthCheckHandlerHTTP(c *gin.Context) {
	msg, err := api.HealthCheckService.HealthCheckServices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusOK, msg, nil)
}
