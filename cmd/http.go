package cmd

import (
	"kautsarhasby/ewallet-ums/helpers"
	"kautsarhasby/ewallet-ums/internal/api"
	"kautsarhasby/ewallet-ums/internal/repository"
	"kautsarhasby/ewallet-ums/internal/services"
	"log"

	"github.com/gin-gonic/gin"
)

func ServeHTTP() {
	dependency := dependencyInject()
	r := gin.Default()

	r.GET("/health", dependency.HealthCheckAPI.HealthCheckHandlerHTTP)

	usersGroup := r.Group("/users")
	usersV1Group := usersGroup.Group("/v1")
	usersV1Group.POST("/register", dependency.RegisterAPI.Register)
	usersV1Group.POST("/login", dependency.LoginAPI.Login)

	err := r.Run(":" + helpers.GetEnv("PORT", "8080"))
	if err != nil {
		log.Fatal(err)
	}
}

type Dependency struct {
	HealthCheckAPI *api.HealthCheckHandler
	RegisterAPI    *api.RegisterHandler
	LoginAPI       *api.LoginHandler
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

	return Dependency{
		HealthCheckAPI: healthCheckAPI,
		RegisterAPI:    registerAPI,
		LoginAPI:       loginAPI,
	}
}
