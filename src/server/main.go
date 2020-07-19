package main

import (
	"fmt"
	"os"
	"strings"
	"path/filepath"
	"config"
	"net"
	"github.com/pkg/errors"                  //https://github.com/pkg/errors       https://godoc.org/github.com/pkg/errors
	kingpin "gopkg.in/alecthomas/kingpin.v2" //https://github.com/alecthomas/kingpin    https://gopkg.in/alecthomas/kingpin.v2
)

var (
	a          = kingpin.New(filepath.Base(os.Args[0]), "A command-line "+config.Progname+" application.")
	configFile     = a.Flag("config", "Configuration file path").Default(config.DefaultConFile).String()
	version    = a.Flag("version", "Show the version information for "+config.Progname).Bool()
)

func main() {

	a.HelpFlag.Short('h')
	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error  commandline arguments\n"))
		os.Exit(10001)          //Error no: AABBB. AA: Package seq,main is 10; BBB: error no 
	}

	if *version {
		myerr := errors.New(config.Progname + " " + config.Proversion + "\n")
		fmt.Fprintln(os.Stderr,errors.WithMessagef(myerr,"Go Version: %s Branch: %s CommitID: %S Build time: %s \n ", config.Goversion, config.BuildBranch,config.Commit, config.Buildstamp))
		os.Exit(10002)
	}

	if config.ParseConfig(*configFile) > 0 {
		myerr := errors.New(fmt.Sprintf("Error parsing config file: %s\n", config.Cfg.Global.Config))
		fmt.Fprintln(os.Stderr, myerr)
    os.Exit(10003)
	}

	if strings.ToUpper(config.Cfg.Log.Loglevel) == "DEBUG" {
		myerr := errors.New(fmt.Sprintf("Parsed config values are:\n %+v\n", config.Cfg))
		fmt.Fprintln(os.Stderr, myerr)
	}
	
	ret := log_init() 
	if ret > 0 {
		os.Exit(ret)
	}

	//Print start message to error log file
	log_error("Starting " + config.Cfg.App.Name +"Server", "Info")
	log_error("Parsing command line parameters successful ", "Info")
	log_error("Parsing config file： "+config.Cfg.Global.Config +" successful ", "Info")

	//Checking the parameters from config file are available
	addr, err := net.ResolveIPAddr("ip", config.Cfg.Global.Listen)
	if err != nil {
		log_startmsg("The serer listen ip " + config.Cfg.Global.Listen + "is invalid !" , "Error")
		os.Exit(10006)
	}else{
		log_startmsg("The serer IP is: " +addr.String() , "Info")
	}

	if config.Cfg.Global.Port <= 1024 || config.Cfg.Global.Port >= 65535 {
		log_startmsg("The listen port should greater than 1024 and less than 65535 !" , "Error")
		os.Exit(10007)
	}else{
		log_startmsg("The listen port is: " + fmt.Sprintf("%d",config.Cfg.Global.Port), "Info")
	}


	os.Exit(0)

}
