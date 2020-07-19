package config
import (
	"io/ioutil"
	"log"
	yaml "gopkg.in/yaml.v2"
)

//Defined the default value for config 
var Proversion = "v0.0.1"
var Progname = "pipe"
var Commit = "NA" 
var BuildBranch string
var Buildstamp string
var Goversion string 
var DefaultPrefix = "/usr/local/" + Progname
var DefaultConFile = "/usr/local/" + Progname + "/conf/config.yaml"

// Struct for app block. program name 
type app struct {
	Name string 
	Version string 
	Commit string
	BuildBranch string 
	Buildstamp string
	Goversion string 
}

//Struct for global block 
type  global struct {
	Config string `yaml: "config"`
	Listen string `yaml: "listen"`
	Port int `yaml: "port"`
}

//Struct for log block 
type struLog struct {
	Loglevel string `yaml: "loglevel"`
	AccessLog string `yaml: "accesslog"`
	ErrorLog string `yaml: "errorlog"`
}

//Main config struct
type config struct{
	App app
	Global global
	Log struLog
}

var Cfg = new(config)

func ParseConfig(confFile string)(result int){
	// Set default value to cfg struct
	Cfg.App = app { 
		Progname,
		Proversion,
		Commit,
		BuildBranch,
		Buildstamp,
		Goversion,
	}

	if len(confFile) <= 0{
		confFile = DefaultConFile
                Cfg.Global.Config =DefaultConFile
	}else{
            confFile = confFile 
            Cfg.Global.Config = confFile
        }

	yamlFile, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("yamlFile.Get err %v", err)
		return 1
	}

	err = yaml.Unmarshal(yamlFile, Cfg)

	if err != nil {
			log.Fatalf("Unmarshal: %v when to struct", err)
			return 1
	}
	
	return 0
}
