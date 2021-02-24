package main

import (
	"fmt"
	"github.com/csby/gwsf/gtype"
	"github.com/kardianos/service"
)

type Service struct {
}

func (s *Service) Start(name string) error {
	cfg := &service.Config{Name: name}
	svc, err := service.New(nil, cfg)
	if err != nil {
		return err
	}

	return svc.Start()
}

func (s *Service) Stop(name string) error {
	cfg := &service.Config{Name: name}
	svc, err := service.New(nil, cfg)
	if err != nil {
		return err
	}

	return svc.Stop()
}

func (s *Service) Restart(name string) error {
	cfg := &service.Config{Name: name}
	svc, err := service.New(nil, cfg)
	if err != nil {
		return err
	}

	return svc.Restart()
}

func (s *Service) Status(name string) (gtype.ServerStatus, error) {
	cfg := &service.Config{Name: name}
	svc, err := service.New(nil, cfg)
	if err != nil {
		return gtype.ServerStatusUnknown, err
	}

	status, err := svc.Status()
	return gtype.ServerStatus(status), err
}

func (s *Service) Install(name, path string) error {
	cfg := &service.Config{
		Name:        name,
		DisplayName: name,
		Description: name,
		Executable:  path,
	}
	svc, err := service.New(nil, cfg)
	if err != nil {
		return err
	}

	return svc.Install()
}

func (s *Service) RemoteInfo() (*gtype.SvcUpdResult, error) {
	return nil, fmt.Errorf("not implement")
}

func (s *Service) RemoteRestart(name string) error {
	return fmt.Errorf("not implement")
}

func (s *Service) RemoteUpdate(name, path, updateFile, updateFolder string) error {
	return fmt.Errorf("not implement")
}
