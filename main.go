package main

import (
	"fmt"
	"github.com/neimarkbraga/win-node-svc/app"
	"github.com/neimarkbraga/win-node-svc/libs/service"
	"golang.org/x/sys/windows/svc"
	"log"
	"os"
	"strings"
)

func usage(errmsg string) {
	fmt.Fprintf(os.Stderr,
		"%s\n\n"+
			"usage: %s <command>\n"+
			"       where <command> is one of\n"+
			"       install, remove, debug, start, stop.\n",
		errmsg, os.Args[0])
	os.Exit(2)
}

func main() {
	inService, err := svc.IsWindowsService()
	if err != nil {
		log.Fatalf("failed to determine if we are running in service: %v", err)
	}

	if inService {
		service.Run(false)
		return
	}

	if len(os.Args) < 2 {
		usage("no command specified")
	}

	cmd := strings.ToLower(os.Args[1])
	switch cmd {
	case "debug":
		service.Run(true)
		return
	case "install":
		err = service.Install()
	case "remove":
		err = service.Remove()
	case "start":
		err = service.Start()
	case "stop":
		err = service.Control(svc.Stop, svc.Stopped)
	default:
		usage(fmt.Sprintf("invalid command %s", cmd))
	}
	if err != nil {
		log.Fatalf("failed to %s %s: %v", cmd, app.Config.Name, err)
	}
	return
}