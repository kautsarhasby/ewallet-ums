package interfaces

type IHealthCheckServices interface {
	HealthCheckServices() (string, error)
}

type IHealthCheckRepo interface {
	HealthCheckRepository() (string, error)
}
