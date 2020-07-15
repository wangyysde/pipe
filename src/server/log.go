package main

import (
	"io"
	"log"
	"os"
        "config"
)

var log_level = [4]string{"Info", "Warnging", "Error", "debug"}

type str_log_conf struct {
	log_level  int
	access_log string
	error_log  string
}

var log_cfg = str_log_conf{config.Cfg.Log.Loglevel, config.Cfg.Log.AccessLog,config.Cfg.Log.ErrorLog}

var (
	Access_logger *log.Logger
	Error_logger  *log.Logger
)

func log_init() (ret int) {

	Std_logger := log.New(io.Writer(os.Stdout), config.Progname+"-", log.Ldate|log.Ltime|log.Lshortfile)

	Std_logger.Printf("Starting open the  error log file: %s", log_cfg.error_log)
	errLogFd, err := os.OpenFile(log_cfg.error_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Std_logger.Printf("Can not open the error log file: %s, Error message is: %s", log_cfg.error_log, err)
		os.Exit(1)
	}
	Error_logger = log.New(io.Writer(errLogFd), config.Progname, log.Ldate|log.Ltime|log.Lshortfile)
	Std_logger.Printf("Opened error log file: %s successful", log_cfg.error_log)
	err = nil

	Std_logger.Printf("Starting open the  access log file: %s", log_cfg.access_log)
	accLogFd, err := os.OpenFile(log_cfg.access_log, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Std_logger.Printf("Can not open the error log file: %s, Error message is: %s", log_cfg.access_log, err)
		os.Exit(1)
	}
	Access_logger = log.New(io.Writer(accLogFd), config.Progname, log.Ldate|log.Ltime|log.Lshortfile)
	Std_logger.Printf("Opened acess log file: %s successful", log_cfg.access_log)
	err = nil

	return 0
}

func log_access(logmsg string, loglevel int) (ret int) {
	if len(logmsg) > 0 {
		Access_logger.Printf(logmsg)
	}

	return 0
}

func log_error(logmsg string, loglevel int) (ret int) {
	if len(logmsg) > 0 {
		if loglevel >= log_cfg.log_level {
			logPrefix := config.Progname + " - " + log_level[loglevel] + "-"
			Error_logger.SetPrefix(logPrefix)
			Error_logger.Printf(logmsg)
		}
	}

	return 0

}
