package service

import (
	"fmt"
	"github.com/neimarkbraga/win-node-svc/app"
	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/debug"
	"golang.org/x/sys/windows/svc/eventlog"
	"os"
	"time"
)

var eventLog debug.Log

type winService struct{}

func (m *winService) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (ssec bool, errno uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.StartPending}
	process, ended := make(chan *os.Process), make(chan bool)
	go app.Run(process, ended)
	proc := <-process
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
loop:
	for {
		select {
		case <-ended:
			eventLog.Error(1, fmt.Sprintf("node process stopped"))
			proc.Kill()
			break loop
		case c := <-r:
			switch c.Cmd {
			case svc.Interrogate:
				changes <- c.CurrentStatus
				time.Sleep(100 * time.Millisecond)
				changes <- c.CurrentStatus
			case svc.Stop, svc.Shutdown:
				proc.Kill()
				break loop
			default:
				eventLog.Error(1, fmt.Sprintf("unexpected control request #%d", c))
			}
		}
	}
	changes <- svc.Status{State: svc.StopPending}
	return
}

func Run(isDebug bool) {
	var err error
	if isDebug {
		eventLog = debug.New(app.Config.Name)
	} else {
		eventLog, err = eventlog.Open(app.Config.Name)
		if err != nil {
			return
		}
	}
	defer eventLog.Close()

	eventLog.Info(0, fmt.Sprintf("starting %s service", app.Config.Name))

	run := svc.Run
	if isDebug {
		run = debug.Run
	}
	err = run(app.Config.Name, &winService{})
	if err != nil {
		eventLog.Error(1, fmt.Sprintf("%s service failed: %v", app.Config.Name, err))
		return
	}

	eventLog.Info(2, fmt.Sprintf("%s service stopped", app.Config.Name))
}