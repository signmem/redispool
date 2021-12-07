package g

import (
	"log"
	"os"
)

func InitLog() {

	logfile := Config().LogFile
	loggerFile, err := os.OpenFile(logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0664)

	if err != nil {
		panic(err) // Check for error
	}

	log.SetOutput(loggerFile)
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Ltime)
	return
}
