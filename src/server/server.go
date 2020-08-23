package main

import (
	"bzhyserver"
	"config"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"string"
)

type Server struct {
	context    context.Context
	shutdownFn context.CancelFunc
	//	childRoutines      *errgroup.Group
	cfg                *config.Config
	shutdownReason     string
	shutdownInProgress bool

	configFile string
	rootPath   string
	indexs		 []string
	index			 string 
	pidFile    string

	r *bzhyserver.Engine
	//RouteRegister routing.RouteRegister `inject:""`
	//HTTPServer    *api.HTTPServer       `inject:""`
}

var Svr = new(Server)

func NewServer(Cfg *config.Config) *Server {
	Svr.cfg = Cfg
	Svr.configFile = Cfg.Config
	Svr.indexs = string.Split(Cfg.Index)
	if len(Svr.indexs) == 1 && Svr.indexs[0] == "" {
		WriteStartLog("The default page of the server is empty. index.html will be used", "warn")
		Svr.indexs[0] = "index.html"
	}
	Svr.r =  bzhyserver.Default()
	Svr.r.SetAccLogHandler(WriteLog2Acclog)
	Svr.r.SetErrLogHandler(WriteLog2Errlog)
	Svr.r.Get("/",HandlerGetRequest)

	return Svr
}

func(c *bzhyserver.Context){


}

/*
func init_serer() (ret int) {

	r := bzhyserver.Default()
	r.SetAccLogHandler(WriteLog2Acclog)
	r.SetErrLogHandler(WriteLog2Errlog)

	r.GET("/", GetHandler)
	r.POST("/somePost", posting)
		  r.PUT("/somePut", putting)
		  r.DELETE("/someDelete", deleting)
		  r.PATCH("/somePatch", patching)
		  r.HEAD("/someHead", head)
		  r.OPTIONS("/someOptions", options)
			gin.DebugPrintRouteFunc = func(httpMethod, absolutePath, handlerName string, nuHandlers int) {
				//log.Printf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
				logmsg := fmt.Sprintf("endpoint %v %v %v %v\n", httpMethod, absolutePath, handlerName, nuHandlers)
				WriteStartLog(logmsg,"info")
				//log_startmsg(logmsg,"Info")
			}

	    r.GET("/", func(c *gin.Context) {
	        time.Sleep(5*time.Second)
	        c.String(http.StatusOK, "Welcome Gin Server")
	        c.JSON(200, gin.H{
	            "Blog":"www.flysnow.org",
	            "wechat":"flysnow_org",
	        })
	    })
	
	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.Cfg.Global.Port),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			WriteStartLog(fmt.Sprintf("Listen %s:%s %s", config.Cfg.Host, config.Cfg.Port, err), "fatal")
		}
	}()

	return 0
}

*/