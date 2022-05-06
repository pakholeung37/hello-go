package main

import (
	"fmt"
	"net/http"
	"time"

	"example.com/hello-gin/pkg/setting"
	"example.com/hello-gin/routers"
)

func main() {
	var router = routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    time.Duration(setting.ReadTimeout) * time.Second,
		WriteTimeout:   time.Duration(setting.WriteTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil {
		return
	}
}
