package commands

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

func parsePatternType(cmd *exec.Cmd, patternType string) error {
	switch patternType {
	case "PREFIXED":
		cmd.Args = append(cmd.Args, "--prefix")
		return nil
	case "LITERAL":
		return nil
	default:
		return errors.New(errPatternInvalid)
	}
}

func parsePermission(cmd *exec.Cmd, permission string) error {
	switch permission {
	case "ALLOW":
		cmd.Args = append(cmd.Args, "--allow")
		return nil
	case "DENY":
		cmd.Args = append(cmd.Args, "--deny")
		return nil
	default:
		return errors.New(errPermissionInvalid)
	}
}

func parseServiceAccount(cmd *exec.Cmd, principal string) error {
	serviceAccount, err := ParsePrincipal(principal)
	if err != nil {
		return err
	}
	cmd.Args = append(cmd.Args, "--service-account", serviceAccount)
	return nil
}

// ParsePrincipal helper method
func ParsePrincipal(principal string) (string, error) {
	fmt.Println("Principal: ", principal)
	split := strings.Split(principal, ":")
	if len(split) != 2 {
		return "", errors.New(errPrincipalInvalid)
	}
	user := split[0]
	if user != "User" {
		return "", errors.New(errPrincipalInvalid)
	}

	serviceAccount := split[1]
	if !strings.Contains(serviceAccount, "sa-") {
		return "", errors.New(errPrincipalInvalid)
	}

	return serviceAccount, nil
}

func parseResource(cmd *exec.Cmd, rName string, rType string) error {
	switch rType {
	case "TOPIC":
		cmd.Args = append(cmd.Args, "--topic", rName)
		return nil
	case "GROUP":
		cmd.Args = append(cmd.Args, "--consumer-group", rName)
		return nil
	case "CLUSTER":
		if rName != "" {
			// return errors.New(errResourceNameSpecifiedWithResourceTypeCluster)
			rName = ""
		}
		cmd.Args = append(cmd.Args, "--cluster-scope")
		return nil
	default:
		return errors.New(errResourceTypeInvalid)
	}
}
