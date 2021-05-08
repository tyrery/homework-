package main

import (
		"net/http"
		"os"
		"os/signal"
		"time"
		"log"
		"golang.org/x/sync/errgroup"
	)


// 主动关闭服务器
var server *http.Server

func main() {
	var g errgroup.Group

   // 一个通知退出的chan
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	// 子协程
	group.Go(func() error {
		server = &http.Server{
			Addr:         ":1210",
			WriteTimeout: time.Second * 4,
			Handler:      mux
		}	
		return nil
	})

	// 捕获err
	if err := group.Wait(); err != nil {
		log.Fatal("Get errors: ", err)
	}else {
		log.Println("Get all num successfully!")
	}

}
go func() {
	// 接收退出信号
	<-quit
	if err := server.Close(); err != nil {
		log.Fatal("Close server:", err)
	}
}()

log.Println("Starting httpServer")
err := server.ListenAndServe()
if err != nil {
	// 正常退出
	if err == http.ErrServerClosed {
		log.Fatal("Server closed under request")
	} else {
		log.Fatal("Server closed unexpected", err)
	}
}
  log.Fatal("Server exited")

}

type myHandler struct{}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte("this is httpServer"))
}

 // 关闭http
func sayBye(w http.ResponseWriter, r *http.Request) {
	  w.Write([]byte("bye bye ,shutdown the server"))     // 没有输出
	  err := server.Shutdown(nil)
	  if err != nil {
		log.Println([]byte("shutdown the server err"))
	  }
 }


