package services

import "kautsarhasby/ewallet-ums/internal/interfaces"

type HealthCheck struct {
	HealthCheckRepository interfaces.IHealthCheckRepository
}

func (s *HealthCheck) HealthCheckServices() (string, error) {
	return "service healthy", nil
}
