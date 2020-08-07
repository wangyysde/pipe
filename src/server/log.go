package main

import (
	"os"
	"strings"
	"bzhylog"
	"fmt"
)

//Define log file descriptor
var (
	ErrFd *os.File = nil
	AccFd *os.File = nil
)

//Define loggger for 
var (
	StdLog *bzhylog.Logger = nil 
	AccLog *bzhylog.Logger = nil 
	ErrLog *bzhylog.Logger = nil 
)

// Create a new instance of the logger for StdOut.
func CreateStdLog() {
	StdLog =  bzhylog.New()
	StdLog.Out = os.Stdout
	StdLog.SetLevel(bzhylog.TraceLevel)
	StdLog.SetFormatter(&bzhylog.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
}

// Create a new instance of the logger for Access Log.
func CreateAccLog(AccLogFile string)(ret int)  {
	AccLog = bzhylog.New()
        AccFd, err := os.OpenFile(AccLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if err == nil {
		AccLog.Out = AccFd
	} else {
		WriteLog2Stdout(fmt.Sprintf("Failed to open the ACCESS log file %s Error message: %s",AccLogFile,err), "fatal")
		return 200001
	}


	return 0
}

// Create a new instance of the logger for Error Log.
func CreateErrLog(ErrLogFile string )(ret int)  {
        ErrLog = bzhylog.New()
        ErrFd, err := os.OpenFile(ErrLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
        if err == nil {
                ErrLog.Out = ErrFd
        } else {
                WriteLog2Stdout(fmt.Sprintf("Failed to open the ERROR log file %s Error message: %s",ErrLogFile,err), "fatal")
                return 200002
        }


        return 0
}

//Write log msg to StdOut
func WriteLog2Stdout(msg string, level string) (ret int){
	switch strings.ToLower(level) {
        	case "panic":
			StdLog.Panic(msg)		
        	case "fatal":
                	StdLog.Fatal(msg)
        	case "error":
                	StdLog.Error(msg)
        	case "warn", "warning":
                	StdLog.Warn(msg)
        	case "info":
                	StdLog.Info(msg)
        	case "debug":
                	StdLog.Debug(msg)
        	case "trace":
                	StdLog.Trace(msg)
		default:
			WriteLog2Stdout("We got a log message without UNKNOW log level", "warn")
        }
	
	return 0
			
}

//Write log msg to Access log file
func WriteLog2Acclog(msg string, level string) (ret int){
        switch strings.ToLower(level) {
                case "panic":
                        AccLog.Panic(msg)
                case "fatal":
                        AccLog.Fatal(msg)
                case "error":
                        AccLog.Error(msg)
                case "warn", "warning":
                        AccLog.Warn(msg)
                case "info":
                        AccLog.Info(msg)
                case "debug":
                        AccLog.Debug(msg)
                case "trace":
                        AccLog.Trace(msg)
                default:
                        WriteLog2Stdout("We got a log message without UNKNOW log level", "warn")
        }

        return 0

}


//Write log msg to Error log file
func WriteLog2Errlog(msg string, level string) (ret int){
  switch strings.ToLower(level) {
    case "panic":
      ErrLog.Panic(msg)
    case "fatal":
      ErrLog.Fatal(msg)
    case "error":
    	ErrLog.Error(msg)
    case "warn", "warning":
      ErrLog.Warn(msg)
    case "info":
      ErrLog.Info(msg)
    case "debug":
      ErrLog.Debug(msg)
    case "trace":
      ErrLog.Trace(msg)
    default:
    	WriteLog2Stdout("We got a log message without UNKNOW log level", "warn")
  }

  return 0

}


//Set the level of access logger, error logger and stdout
func SetLoglevel(level string ) (ret int ) {
	switch strings.ToLower(level) {
		case "panic":
			StdLog.SetLevel(bzhylog.PanicLevel)
			AccLog.SetLevel(bzhylog.PanicLevel)
			ErrLog.SetLevel(bzhylog.PanicLevel)
		case "fatal":
      StdLog.SetLevel(bzhylog.FatalLevel) 
    	AccLog.SetLevel(bzhylog.FatalLevel)
      ErrLog.SetLevel(bzhylog.FatalLevel)
		case "error":
      StdLog.SetLevel(bzhylog.ErrorLevel) 
      AccLog.SetLevel(bzhylog.ErrorLevel)
      ErrLog.SetLevel(bzhylog.ErrorLevel)
		case "warn":
      StdLog.SetLevel(bzhylog.WarnLevel) 
    	AccLog.SetLevel(bzhylog.WarnLevel)
      ErrLog.SetLevel(bzhylog.WarnLevel)
		case "info":
      StdLog.SetLevel(bzhylog.InfoLevel) 
      AccLog.SetLevel(bzhylog.InfoLevel)
      ErrLog.SetLevel(bzhylog.InfoLevel)
		case "debug":
      StdLog.SetLevel(bzhylog.DebugLevel) 
      AccLog.SetLevel(bzhylog.DebugLevel)
      ErrLog.SetLevel(bzhylog.DebugLevel)
		case "trace":
      StdLog.SetLevel(bzhylog.TraceLevel) 
      AccLog.SetLevel(bzhylog.TraceLevel)
      ErrLog.SetLevel(bzhylog.TraceLevel)
		default:
			WriteLog2Stdout("We got a invalid level of log. We will set log level to FATAL", "fatal")
			StdLog.SetLevel(bzhylog.FatalLevel) 
      AccLog.SetLevel(bzhylog.FatalLevel)
      ErrLog.SetLevel(bzhylog.FatalLevel)
	}
	
	return 0
}

//Closing the file descriptors of accesss log.
func CloseAccLogFd()(ret int){
	if AccFd != nil {
		err := AccFd.Close()
		if err != nil {
			WriteLog2Errlog(fmt.Sprintf("Closing Access log err %s", err),"error")
		}
		AccFd = nil
	}

	return 0
} 

//Closing the file descriptors of error log.
func CloseErrLogFd()(ret int){
	if ErrFd != nil {
		 ErrFd.Close()
	}
	return 0
}

//Writing Starting log message to Stdout and Error log file 
func WriteStartLog(msg string, level string){
	WriteLog2Stdout(msg, level)
	if ErrFd != nil {
		WriteLog2Errlog(msg, level)
	}
	
}

