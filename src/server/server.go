package main

import "github.com/gin-gonic/gin"
import "context"
import "net/http"
import "os"
import "os/signal"
import "syscall"
import "time"
import "log"
import "config"
import "strconv"
import "fmt"

func init_serer()(ret int) {
		r := gin.Default()
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
			Addr:    ":" + strconv.Itoa(config.Cfg.Global.Port) ,
      Handler: r,
    }

    go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
    }()
    
    quit := make(chan os.Signal)

    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
    log.Println("Shutting down server...")
    
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	
    log.Println("Server exiting")
    
	return 0
}
