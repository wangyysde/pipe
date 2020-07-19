package main

import (
	"io"
	"log"
	"os"
	"strings"
  "config"
)

//var log_level = [4]string{"Info", "Warnging", "Error", "debug"}
var log_level = make(map[string]int)

type str_log_conf struct {
	log_level string 
	access_log string
	error_log  string
	errLogFd  *os.File
	accLogFd  *os.File
}

var (
	Access_logger *log.Logger
	Error_logger  *log.Logger
	Std_logger	*log.Logger
)

var log_cfg str_log_conf

func log_init() (ret int) {

	//Initating log_level map
	log_level = map[string]int{
		"INFO": 0, 
		"WARNGING": 1, 
		"ERROR": 2, 
		"DEBUG": 3,
	}
	Std_logger = log.New(io.Writer(os.Stdout), config.Progname+"-", log.Ldate|log.Ltime|log.Lshortfile)

	Std_logger.Printf("Starting open the  error log file: %s", config.Cfg.Log.ErrorLog)
	errLogFd, err := os.OpenFile(config.Cfg.Log.ErrorLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Std_logger.Printf("Can not open the error log file: %s, Error message is: %s", config.Cfg.Log.ErrorLog, err)
		return 10004
	}
	Error_logger = log.New(io.Writer(errLogFd), config.Progname, log.Ldate|log.Ltime|log.Lshortfile)
	Std_logger.Printf("Opened error log file: %s successful", config.Cfg.Log.ErrorLog)
	err = nil

	Std_logger.Printf("Starting open the  access log file: %s", config.Cfg.Log.AccessLog)
	accLogFd, err := os.OpenFile(config.Cfg.Log.AccessLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		Std_logger.Printf("Can not open the error log file: %s, Error message is: %s",config.Cfg.Log.AccessLog, err)
		return 10005
	}
	Access_logger = log.New(io.Writer(accLogFd), config.Progname, log.Ldate|log.Ltime|log.Lshortfile)
	Std_logger.Printf("Opened acess log file: %s successful", config.Cfg.Log.AccessLog)
	err = nil

	log_cfg = str_log_conf{strings.ToUpper(config.Cfg.Log.Loglevel), config.Cfg.Log.AccessLog,config.Cfg.Log.ErrorLog,errLogFd,accLogFd}

	return 0
}

func log_access(logmsg string) (ret int) {
	if len(logmsg) > 0 {
		Access_logger.Printf(logmsg)
	}

	return 0
}

func log_error(logmsg string, loglevel string) (ret int) {
	if len(logmsg) > 0 {
		if log_level[strings.ToUpper(log_cfg.log_level)] >= log_level[strings.ToUpper(loglevel)]{
			logPrefix := config.Progname + " - " + loglevel + "-"
			Error_logger.SetPrefix(logPrefix)
			Error_logger.Printf(logmsg)
		}
	}
	return 0
}

func log_startmsg(logmsg string,loglevel string) (ret int){
	if len(logmsg) > 0 {
		logPrefix := config.Progname + " - " + loglevel + "-"
		Std_logger.SetPrefix(logPrefix)
		Std_logger.Printf(logmsg)
		if log_level[strings.ToUpper(log_cfg.log_level)] >= log_level[strings.ToUpper(loglevel)]{
			Error_logger.SetPrefix(logPrefix)
			Error_logger.Printf(logmsg)
		}
	}

	return 0
}

func log_close() ( ret int){
	err := log_cfg.errLogFd.Close()
	if err != nil {
		Std_logger.Printf("Close error log file: Error %s, Error message is: %s", log_cfg.error_log, err)
	}
	err = nil
	 
	err = log_cfg.accLogFd.Close()
	if err != nil {
		Std_logger.Printf("Close access log file: Error %s, Error message is: %s", log_cfg.access_log, err)
	}
	err = nil

	return 0
}
