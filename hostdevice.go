package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

type HostDevice struct {
	HostPath      string `json:"hostPath"`
	ContainerPath string `json:"containerPath"`
	Permission    string `json:"permission"`
}

func (d *HostDevice) validate() error {
	numGlobHostPath := strings.Count(d.HostPath, "*")
	if numGlobHostPath > 1 {
		return fmt.Errorf("HostPath can container only one '*' character: %s", d.HostPath)
	}

	if numGlobHostPath == 1 {
		if !strings.HasSuffix(d.ContainerPath, "*") {
			return fmt.Errorf("ContainerPath should ends with '*' character when HostPath container '*': %s", d.ContainerPath)
		}
		return nil
	}

	if strings.Contains(d.ContainerPath, "*") {
		return fmt.Errorf("ContainerPath must not contain '*' when HostPath does not contain '*': %s", d.ContainerPath)
	}

	return nil
}

type ExpandedHostDevice struct {
	HostPath      string
	ContainerPath string
	Permission    string
}

func (d HostDevice) Expand() ([]*ExpandedHostDevice, error) {
	if err := d.validate(); err != nil {
		return nil, err
	}

	matchedHostPath, err := filepath.Glob(d.HostPath)
	if err != nil {
		return nil, err
	}

	expanded := []*ExpandedHostDevice{}
	baseHostPath := strings.Split(d.HostPath, "*")[0]
	baseContainerPath := strings.Split(d.ContainerPath, "*")[0]
	for _, hp := range matchedHostPath {
		expanded = append(expanded, &ExpandedHostDevice{
			HostPath:      hp,
			ContainerPath: strings.Replace(hp, baseHostPath, baseContainerPath, 1),
			Permission:    d.Permission,
		})
	}
	return expanded, nil
}
