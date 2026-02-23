package api

import (
	"kautsarhasby/ewallet-ums/constants"
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RefreshTokenHandler struct {
	RefreshTokenService interfaces.IRefreshTokenService
}

func (api *RefreshTokenHandler) RefreshToken(c *gin.Context) {
	var (
		log          = helpers.Logger
		refreshToken = c.Request.Header.Get("Authorization")
	)

	claim, ok := c.Get("token")
	if !ok {
		log.Error("failed to get claim in context")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	tokenClaim, ok := claim.(*helpers.TokenClaims)
	if !ok {
		log.Error("failed to parse claim in tokenClaim")
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}
	response, err := api.RefreshTokenService.RefreshToken(c.Request.Context(), refreshToken, *tokenClaim)
	if err != nil {
		log.Error("failed on refresh token service  ", err)
		helpers.SendResponseHTTP(c, http.StatusInternalServerError, constants.ErrServerError, nil)
		return
	}

	helpers.SendResponseHTTP(c, http.StatusCreated, constants.SuccessMessage, response)

}
