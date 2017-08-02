package servicemanager

// Service is a minimal footprint to start a service
type Service interface {
	Name() string
	Run()
}
