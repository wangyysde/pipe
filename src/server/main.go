package main

import (
	"fmt"
	"os"
	"path/filepath"
	"config"
	"github.com/pkg/errors"                  //https://github.com/pkg/errors       https://godoc.org/github.com/pkg/errors
	kingpin "gopkg.in/alecthomas/kingpin.v2" //https://github.com/alecthomas/kingpin    https://gopkg.in/alecthomas/kingpin.v2
)

var (
	a          = kingpin.New(filepath.Base(os.Args[0]), "A command-line "+config.Progname+" application.")
//	debug      = a.Flag("debug", "Enable debug mode.").Default("false").Bool()
//	serverIP   = a.Flag("server", "Server address.").Default("0.0.0.0").IP()
//	serverPort = a.Flag("port", "Server Port").Default("8080").Int()
	configFile     = a.Flag("config", "Configuration file path").Default(config.DefaultConFile).String()
//	accesslog  = a.Flag("access-log", "Access Log file").Default("/var/log/" + progname + ".log").String()
//	errorlog   = a.Flag("error-log", "Error Log file").Default("/var/log/" + progname + "-error.log").String()
	version    = a.Flag("version", "Show the version information for "+config.Progname).Bool()
)

func main() {

	a.HelpFlag.Short('h')
	_, err := a.Parse(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing commandline arguments"))
		os.Exit(2)
	}

	if *version {
		fmt.Printf(config.Progname + " " + config.Proversion + "\n")
		os.Exit(1)
	}

	if config.ParseConfig(*configFile) > 0 {
		//fmt.Fprintln(os.Stderr, errors.Wrapf(err, "Error parsing config file: %s", config.Cfg.Global.Config))
	       fmt.Printf("Error parsing config file: %s\n", config.Cfg.Global.Config)	
               os.Exit(1)
	}

	fmt.Printf("config values :%+v",config.Cfg)
/*
	if *serverPort <= 1024 || *serverPort >= 65535 {
		fmt.Printf("The listen port should greater than 1024 and less than 65535\n")
		os.Exit(1)
	}

	if *debug {
		log_cfg.log_level = 3
	}

	if accesslog != nil {
		log_cfg.access_log = *accesslog
	}

	if errorlog != nil {
		log_cfg.error_log = *errorlog
	}
*/
	
	//log_init()

	//*debug = true
//	fmt.Printf("Debug is:%v\n", *debug)
//	fmt.Println("config is:", *config)
//	fmt.Printf("this is a test message")
	//kingpin.Usage()
}
