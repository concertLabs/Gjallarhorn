package servicemanager

import "log"

// Manager handles all Run functions of all services
type Manager struct {
	Services map[string]Service
}

// NewManager returns a new Manager with n entries
func NewManager(n int) *Manager {
	return &Manager{
		Services: make(map[string]Service, n),
	}
}

// Start tries to start all services
func (m *Manager) Start() error {
	for name, srvc := range m.Services {
		go srvc.Run()
		log.Printf("[Service Manager] Service %s started.\n", name)
	}
	return nil
}

// Stop kills all services
func (m *Manager) Stop() {
	// NOTE: we can do some cleanup work in this function
	return
}

// Add adds the service s and tries to start it
func (m *Manager) Add(s Service) {
	m.Services[s.Name()] = s
}
