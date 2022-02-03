package app

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/neimarkbraga/win-node-svc/libs/utils"
	"os"
	"os/exec"
	"time"
)

func WriteLog(text string, filename string) {
	var err error
	if _, err = os.Stat(Config.LogDirectory); errors.Is(err, os.ErrNotExist) {
		err = os.MkdirAll(Config.LogDirectory, os.ModePerm)
	}
	if err != nil {
		return
	}

	filePath := Config.LogDirectory + filename
	if _, err = os.Stat(filePath); errors.Is(err, os.ErrNotExist) {
		err = os.WriteFile(filePath, []byte(""), os.ModePerm)
	}
	if err != nil {
		return
	}

	file, err := os.OpenFile(filePath, os.O_APPEND, os.ModePerm)
	if err != nil {
		return
	}
	defer file.Close()

	file.WriteString(text + "\n")
}

func Run(process chan *os.Process, ended chan bool) {
	defer func() {
		ended <- true
	}()

	nodejs, err := utils.GetNodeJsExePath()
	utils.PanicOnError(err)

	cmd := exec.Command(nodejs, Config.EntryFile)
	cmd.Dir = Config.WorkingDirectory

	stdout, err := cmd.StdoutPipe()
	utils.PanicOnError(err)

	err = cmd.Start()
	utils.PanicOnError(err)

	process <- cmd.Process

	logFilename := "\\service." + fmt.Sprint(time.Now().Unix()) + ".log"
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		WriteLog(text, logFilename)
		utils.PanicOnError(err)
	}
}