package docker

import "testing"

func TestStart(t *testing.T) {
	// stop a container with an exist containerID, expected err
	if err := Start("sdfsfsadfasfa"); err == nil {
		t.Error("expected return a error when start a containter with a not exist containerID")
	}

	// Start a container with an exist containerID, expected nil
	if err := Stop("28cb50796bcefec59e1c5c36c26249fbe3761c1c751dd0de399b897ec0a1efa1"); err != nil {
		t.Error(err)
	}

	if err := Start("28cb50796bcefec59e1c5c36c26249fbe3761c1c751dd0de399b897ec0a1efa1"); err != nil {
		t.Error(err)
	}
}
