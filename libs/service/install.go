package service

import (
	"fmt"
	"github.com/neimarkbraga/win-node-svc/app"
	"golang.org/x/sys/windows/svc/eventlog"
	"golang.org/x/sys/windows/svc/mgr"
	"os"
	"path/filepath"
)

func getExePath() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("%s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("%s is directory", p)
		}
	}
	return "", err
}

func Install() error {
	exepath, err := getExePath()
	if err != nil {
		return err
	}

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(app.Config.Name)
	if err == nil {
		s.Close()
		return fmt.Errorf("service %s already exists", app.Config.Name)
	}

	s, err = m.CreateService(app.Config.Name, exepath, mgr.Config{
		ServiceType: app.Config.ServiceType,
		StartType: app.Config.StartType,
		ErrorControl: app.Config.ErrorControl,
		LoadOrderGroup: app.Config.LoadOrderGroup,
		TagId: app.Config.TagId,
		Dependencies: app.Config.Dependencies,
		ServiceStartName: app.Config.ServiceStartName,
		DisplayName: app.Config.DisplayName,
		Password: app.Config.Password,
		Description: app.Config.Description,
		SidType: app.Config.SidType,
		DelayedAutoStart: app.Config.DelayedAutoStart,
	})
	if err != nil {
		return err
	}
	defer s.Close()

	err = eventlog.InstallAsEventCreate(app.Config.Name, eventlog.Error|eventlog.Warning|eventlog.Info)
	if err != nil {
		s.Delete()
		return fmt.Errorf("SetupEventLogSource() failed: %s", err)
	}
	return nil
}

func Remove() error {
	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(app.Config.Name)
	if err != nil {
		return fmt.Errorf("service %s is not installed", app.Config.Name)
	}

	defer s.Close()
	err = s.Delete()
	if err != nil {
		return err
	}

	err = eventlog.Remove(app.Config.Name)
	if err != nil {
		return fmt.Errorf("RemoveEventLogSource() failed: %s", err)
	}
	return nil
}