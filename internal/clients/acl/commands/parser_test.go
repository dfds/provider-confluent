package commands

import (
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	saID  = "sa-12351"
	rName = "testingaclResourceName"
)

func TestParsePatternType(t *testing.T) {

	assert := assert.New(t)
	err := parsePatternType(&exec.Cmd{}, "")
	if err == nil {
		t.Errorf("Expected error from parsePatternType with empty patternType arg")
	} else if err.Error() != errPatternInvalid {
		t.Error(err)
	}

	err = parsePatternType(&exec.Cmd{}, "LITERAL")
	if err != nil {
		t.Error(err)
	}

	cmd := exec.Cmd{}
	err = parsePatternType(&cmd, "PREFIXED")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 1 {
			assert.Equal(cmd.Args[0], "--prefix")
		} else {
			t.Errorf("cmd.Args contains more or less than 1 element.")
		}
	}
}

func TestParsePermission(t *testing.T) {
	assert := assert.New(t)

	err := parsePermission(&exec.Cmd{}, "")
	if err == nil {
		t.Errorf("Expected error from parsePermission with empty permission arg")
	} else if err.Error() != errPermissionInvalid {
		t.Error(err)
	}

	cmd := exec.Cmd{}
	err = parsePermission(&cmd, "ALLOW")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 1 {
			assert.Equal(cmd.Args[0], "--allow")
		} else {
			t.Errorf("cmd.Args contains more or less than 1 element.")
		}
	}

	cmd = exec.Cmd{}
	err = parsePermission(&cmd, "DENY")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 1 {
			assert.Equal(cmd.Args[0], "--deny")
		} else {
			t.Errorf("cmd.Args contains more or less than 1 element.")
		}
	}
}

func TestParseServiceAccount(t *testing.T) {
	assert := assert.New(t)

	cmd := exec.Cmd{}
	err := parseServiceAccount(&cmd, fmt.Sprintf("User:%s", saID))
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 2 {
			assert.Equal(cmd.Args[0], "--service-account")
			assert.Equal(cmd.Args[1], saID)
		} else {
			t.Errorf("cmd.Args contains more or less than 2 element.")
		}
	}

}

func TestParsePrincipal(t *testing.T) {
	assert := assert.New(t)
	_, err := ParsePrincipal("")
	if err != nil {
		assert.Equal(err.Error(), errPrincipalInvalid)
	}

	_, err = ParsePrincipal("Cluster:sa-12351")
	if err != nil {
		assert.Equal(err.Error(), errPrincipalInvalid)
	}

	_, err = ParsePrincipal("User:sb-12351")
	if err != nil {
		assert.Equal(err.Error(), errPrincipalInvalid)
	}

	sa, err := ParsePrincipal(fmt.Sprintf("User:%s", saID))
	if err != nil {
		t.Error(err)
	}

	assert.Equal(sa, saID)
}

func TestParseResource(t *testing.T) {
	assert := assert.New(t)

	err := parseResource(&exec.Cmd{}, "", "")
	if err == nil {
		t.Errorf("Expected error from parseResource with empty rName arg && empty rType arg")
	} else if err.Error() != errResourceTypeInvalid {
		t.Error(err)
	}

	cmd := exec.Cmd{}
	err = parseResource(&cmd, rName, "TOPIC")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 2 {
			assert.Equal(cmd.Args[0], "--topic")
			assert.Equal(cmd.Args[1], rName)
		} else {
			t.Errorf("cmd.Args contains more or less than 2 element.")
		}
	}

	cmd = exec.Cmd{}
	err = parseResource(&cmd, rName, "GROUP")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 2 {
			assert.Equal(cmd.Args[0], "--consumer-group")
			assert.Equal(cmd.Args[1], rName)
		} else {
			t.Errorf("cmd.Args contains more or less than 2 element.")
		}
	}

	cmd = exec.Cmd{}
	err = parseResource(&cmd, "kafka-cluster", "CLUSTER")
	if err != nil {
		t.Error(err)
	} else {
		if len(cmd.Args) == 1 {
			assert.Equal(cmd.Args[0], "--cluster-scope")
		} else {
			t.Errorf("cmd.Args contains more or less than 1 element.")
		}
	}

	cmd = exec.Cmd{}
	err = parseResource(&cmd, rName, "CLUSTER")
	if err == nil {
		t.Errorf("Expected error from parseResource with rType arg set to CLUSTER & rName arg not being empty")
	} else if err.Error() != errResourceNameSpecifiedWithResourceTypeCluster {
		t.Error(err)
	}
}
