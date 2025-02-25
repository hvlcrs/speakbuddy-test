package main

import (
	"log"
	"speakbuddy/pkg/configs"
	"speakbuddy/pkg/db"
	"speakbuddy/routers"

	"github.com/fvbock/endless"
)

func init() {
	// Load application configurations
	configs.ConfigInit()
	configs.DatabaseInit()

	// Initialize data layer connections
	db.Init()
}

func main() {
	r := routers.Init()

	log.Printf("[INFO] Server start listening on port %s\n", configs.ServerSetting.HttpPort)

	// Graceful restart using endless
	endless.DefaultReadTimeOut = configs.ServerSetting.ReadTimeout
	endless.DefaultWriteTimeOut = configs.ServerSetting.WriteTimeout
	srv := endless.NewServer(":"+configs.ServerSetting.HttpPort, r)

	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("[FAIL] Server fail to listen: %s\n", err)
	}
}
