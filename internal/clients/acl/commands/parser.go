package commands

import (
	"errors"
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
	split := strings.Split(principal, ":")
	if len(split) != 2 {
		errors.New(errPrincipalInvalid)
	}
	user := split[0]
	if user != "User" {
		errors.New(errPrincipalInvalid)
	}

	serviceaccount := split[1]
	if strings.Contains(serviceaccount, "sa-") {
		errors.New(errPrincipalInvalid)
	}

	cmd.Args = append(cmd.Args, "--service-account", serviceaccount)
	return nil
}

func parseResource(cmd *exec.Cmd, rName string, rType string) error {
	switch rType {
	case "TOPIC":
		cmd.Args = append(cmd.Args, "--topic", rName)
		return nil
	case "CONSUMER_GROUP":
		cmd.Args = append(cmd.Args, "--consumer-group", rName)
		return nil
	case "CLUSTER":
		cmd.Args = append(cmd.Args, "--cluster-scope")
		return nil
	default:
		return errors.New(errResourceTypeInvalid)
	}
}
