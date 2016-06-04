package daemon

import "os/exec"

// ServiceMock is the mock implementation of the Service interface.
type ServiceMock struct {
	Command             string
	processes           map[int]*exec.Cmd
	Args                []string
	ServiceStopResult   bool
	ServiceStopError    error
	ServiceList         map[int]*exec.Cmd
	ServiceStartCmd     *exec.Cmd
	ServiceStartCount   int
	ServicePort         int
	ServiceStopCount    int
	ServicesSetupCalled bool

	// ExecFunc sets the function to run when starting commands
	ExecFunc func() *exec.Cmd
}

// Setup the Management services.
func (s *ServiceMock) Setup() {
	s.ServicesSetupCalled = true
}

// Stop a Service and returns the exit status.
func (s *ServiceMock) Stop(pid int) (bool, error) {
	if _, ok := s.processes[pid]; ok {
		s.processes[pid].Process.Kill()
	}
	s.ServiceStopCount++
	return s.ServiceStopResult, s.ServiceStopError
}

// List all Service PIDs.
func (s *ServiceMock) List() map[int]*exec.Cmd {
	return s.ServiceList
}

// Start a Service and log its output.
func (s *ServiceMock) Start() *exec.Cmd {

	s.ServiceStartCount++
	cmd := s.ExecFunc()
	cmd.Start()
	if s.processes == nil {
		s.processes = make(map[int]*exec.Cmd)
	}
	s.processes[cmd.Process.Pid] = cmd

	return cmd
}

// NewService creates a new MockService with default settings.
func (s *ServiceMock) NewService(args []string) (int, Service) {
	return s.ServicePort, s
}
