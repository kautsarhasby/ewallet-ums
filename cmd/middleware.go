package cmd

import (
	"kautsarhasby/ewallet-ums/constants"
	"kautsarhasby/ewallet-ums/helpers"
	"net/http"

	"log"

	"github.com/gin-gonic/gin"
)

func (dependency *Dependency) MiddlewareValidateAuth(ctx *gin.Context) {
	token := ctx.Request.Header.Get("Authorization")
	if token == "" {
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnathorized, nil)
		ctx.Abort()
		return
	}

	claimToken, err := helpers.ValidateToken(ctx.Request.Context(), token)
	if err != nil {
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnathorized, nil)
		ctx.Abort()
		return
	}

	err = dependency.LogoutAPI.LogoutService.Logout(ctx.Request.Context(), token)
	if err != nil {
		helpers.SendResponseHTTP(ctx, http.StatusInternalServerError, constants.ErrServerError, nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claimToken)

	ctx.Next()
	return
}

func (dependency *Dependency) MiddlewareValidateRefresh(ctx *gin.Context) {
	refreshToken := ctx.Request.Header.Get("Authorization")
	if refreshToken == "" {
		log.Println("authorization empty")
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnathorized, nil)
		ctx.Abort()
		return
	}

	_, err := dependency.UserRepository.GetUserSessionByRefreshToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnathorized, nil)
		ctx.Abort()
		return
	}

	claim, err := helpers.ValidateToken(ctx.Request.Context(), refreshToken)
	if err != nil {
		helpers.SendResponseHTTP(ctx, http.StatusUnauthorized, constants.ErrUnathorized, nil)
		ctx.Abort()
		return
	}

	ctx.Set("token", claim)
	return
}
