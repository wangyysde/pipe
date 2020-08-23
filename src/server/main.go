package main

import (
	"fmt"
	"os"
	"path/filepath"
	"config"
	"net"
//	"github.com/pkg/errors"                  //https://github.com/pkg/errors       https://godoc.org/github.com/pkg/errors
	kingpin "gopkg.in/alecthomas/kingpin.v2" //https://github.com/alecthomas/kingpin    https://gopkg.in/alecthomas/kingpin.v2
)

var (
	a          = kingpin.New(filepath.Base(os.Args[0]), "A command-line "+config.Progname+" application.")
	configFile     = a.Flag("config", "Configuration file path").Default(config.DefaultConFile).String()
	version    = a.Flag("version", "Show the version information for "+config.Progname).Bool()
)

func main() {
	
	//Initating Stdout log descriptor
	CreateStdLog()

	WriteStartLog(fmt.Sprintf("Starting %s server\n", config.Cfg.App.Name ),"info")
	a.HelpFlag.Short('h')
	_, err := a.Parse(os.Args[1:])
	if err != nil {
		WriteStartLog(fmt.Sprintf("Error commandline arguments:%s\n",err),"fatal")
		os.Exit(10001)          //Error no: AABBB. AA: file seq,main is 1; BBB: error no 
	}
	WriteStartLog("Parsed commandline parameters successful","info")

	if *version {
		WriteStartLog(fmt.Sprintf("%s: %s \n",config.Progname,config.Proversion),"info")
		WriteStartLog(fmt.Sprintf("Go Version: %s Branch: %s CommitID: %S Build time: %s \n",config.Goversion, config.BuildBranch,config.Commit, config.Buildstamp ),"info")
		os.Exit(0)
	}

	if config.ParseConfig(*configFile) > 0 {
		WriteStartLog(fmt.Sprintf("Error parsing config file: %s\n", config.Cfg.Global.Config),"fatal")
    os.Exit(10002)
	}
	
	WriteStartLog("Parsed configure file successful","info")

	WriteStartLog(fmt.Sprintf("Opening error log file: %s \n", config.Cfg.Log.ErrorLog),"info")
	if CreateErrLog(config.Cfg.Log.ErrorLog) > 0 {
		os.Exit(10003)
	}
	defer CloseErrLogFd() 
	WriteStartLog("Openned error log file \n","info")

	WriteStartLog(fmt.Sprintf("Opening access log file: %s \n", config.Cfg.Log.AccessLog),"info")
	if CreateAccLog(config.Cfg.Log.AccessLog) > 0 {
		os.Exit(10004)
	}
	defer CloseErrLogFd()
	WriteStartLog("Openned access log file \n","info")	

	SetLoglevel(config.Cfg.Log.Loglevel)

	//Checking the parameters from config file are available
	addr, err := net.ResolveIPAddr("ip", config.Cfg.Global.Listen)
	if err != nil {
		WriteStartLog(fmt.Sprintf("The serer listen ip:%s is invalid !\n", config.Cfg.Global.Listen), "fatal") 
		os.Exit(10005)
	}else{
		WriteStartLog(fmt.Sprintf("The serer IP is: %s\n",addr.String()),"info")
	}

	if config.Cfg.Global.Port <= 1024 || config.Cfg.Global.Port >= 65535 {
		WriteStartLog("The listen port should greater than 1024 and less than 65535 !","fatal")
		os.Exit(10006)
	}else{
		WriteStartLog(fmt.Sprintf("The listen port is : %d\n",config.Cfg.Global.Port),"info")
	}

	WriteStartLog("Now starting the server....","info")
	WriteStartLog("test error message....","error")
	WriteStartLog("test fatal message....","fatal")
	if ret := init_serer(); ret >0 {
		WriteStartLog("Starting the server ERROR","fatal")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
<-quit
	WriteLog2Errlog("Shutting down server...","info")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
if err := srv.Shutdown(ctx); err != nil {
	log.Fatal("Server forced to shutdown:", err)
}

	log.Println("Server exiting")
	

	os.Exit(0)

}

