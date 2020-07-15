package config
import (
	"fmt"
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v2"
)

//Defined the default value for config 
var proversion = "v0.0.1"
var progname = "pipe"
var commit = "NA" 
var buildBranch = "master"
var buildstamp string
var defaultPrefix = "/usr/local/" + progname
var defaultConFile = "/usr/local/" + progname + "/conf/config.yaml"

// Struct for app block. program name 
type app struct {
	name string  `yaml: "name"`
	version string 
	commit string
	buildBranch string 
	buildstamp string 
}

//Struct for global block 
type  global struct {
	Config string `yaml: "config"`
	Listen string `yaml: "listen"`
	Port int `yaml: "port"`
}

//Struct for log block 
type struLog struct {
	Loglevel int `yaml: "loglevel"`
	AccessLog string `yaml: "accesslog"`
	ErrorLog string `yaml: "errorlog"`
}

//Main config struct
type config struct{
	App app
	Global global
	Log stru_log
}

var cfg = new(config)

func ParseConfig(confFile string)(result int){
	// Set default value to cfg struct
	cfg.App = app { 
		progname,
		proversion,
		commit,
		buildBranch,
		buildstamp
	}

	if len(confFile) <= 0{
		confFile = defaultConFile
	}

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
		return 1
	}

	err = yaml.Unmarshal(yamlFile, cfg)

	if err != nil {
			log.Fatalf("Unmarshal: %v when to struct", err)
			return 1
	}
	
	return 0
}
