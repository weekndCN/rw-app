package docker

import "testing"

func TestStop(t *testing.T) {
	// stop a container with an exist containerID, expected err
	if err := Stop("fsafs"); err == nil {
		t.Error("expected return a error with passing a not exist containerID")
	}

	// stop a container with an not exist containerID, expected nil
	if err := Start("28cb50796bcefec59e1c5c36c26249fbe3761c1c751dd0de399b897ec0a1efa1"); err != nil {
		t.Error(err)
	}

	if err := Stop("28cb50796bcefec59e1c5c36c26249fbe3761c1c751dd0de399b897ec0a1efa1"); err != nil {
		t.Error(err)
	}
	// don't block other test file
	if err := Start("28cb50796bcefec59e1c5c36c26249fbe3761c1c751dd0de399b897ec0a1efa1"); err != nil {
		t.Error(err)
	}

}
