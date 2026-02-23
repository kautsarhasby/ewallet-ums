package cmd

import (
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/api"
	"kautsarhasby/ewallet-ums/internal/interfaces"
	"kautsarhasby/ewallet-ums/internal/repository"
	"kautsarhasby/ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()
	r := gin.Default()

	r.GET("/health", dependency.HealthCheckAPI.HealthCheckHandlerHTTP)

	usersV1 := r.Group("/users/v1")
	usersV1.POST("/register", dependency.RegisterAPI.Register)
	usersV1.POST("/login", dependency.LoginAPI.Login)

	usersV1WithAuth := usersV1.Use()
	usersV1WithAuth.DELETE("/logout", dependency.MiddlewareValidateAuth, dependency.LogoutAPI.Logout)
	usersV1WithAuth.PUT("/refresh-token", dependency.MiddlewareValidateRefresh, dependency.RefreshAPI.RefreshToken)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	UserRepository interfaces.IUserRepository
	HealthCheckAPI *api.HealthCheckHandler
	RegisterAPI    *api.RegisterHandler
	LoginAPI       *api.LoginHandler
	LogoutAPI      *api.LogoutHandler
	RefreshAPI     *api.RefreshTokenHandler
}

func dependencyInject() Dependency {

	userRepo := &repository.UserRepository{
		DB: helpers.DB,
	}

	// check health
	healthCheckSvc := &services.HealthCheck{}
	healthCheckAPI := &api.HealthCheckHandler{
		HealthCheckService: healthCheckSvc,
	}

	// register
	registerSvc := &services.RegisterService{
		UserRepository: userRepo,
	}
	registerAPI := &api.RegisterHandler{
		RegisterService: registerSvc,
	}

	// login
	loginSvc := &services.LoginService{
		UserRepository: userRepo,
	}
	loginAPI := &api.LoginHandler{
		LoginService: loginSvc,
	}

	// logout
	logoutSvc := &services.LogoutService{
		UserRepository: userRepo,
	}
	logoutAPI := &api.LogoutHandler{
		LogoutService: logoutSvc,
	}

	// refresh token
	refreshSvc := &services.RefreshTokenService{
		UserRepository: userRepo,
	}
	refreshAPI := &api.RefreshTokenHandler{
		RefreshTokenService: refreshSvc,
	}

	return Dependency{
		UserRepository: userRepo,
		HealthCheckAPI: healthCheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
		LogoutAPI:      logoutAPI,
		RefreshAPI:     refreshAPI,
	}
}
