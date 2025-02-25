package configs

import (
	"log"
	"time"

	"github.com/caarlos0/env/v11"
)

type App struct {
	FileSavePath  string   `env:"FILE_SAVE_PATH" envDefault:"./files/upload/"`
	FileMaxSize   int      `env:"FILE_MAX_SIZE" envDefault:"50"`
	FileAllowExts []string `env:"FILE_ALLOW_EXTS" envDefault:".mp4a,.mp3,.wav"`
	FileTargetExt string   `env:"FILE_TARGET_EXT" envDefault:".wav"`
}

var AppSetting = &App{}

type Server struct {
	HttpPort     string        `env:"HTTP_PORT" envDefault:"8081"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"300s"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"300s"`
}

var ServerSetting = &Server{}

func ConfigInit() {
	if err := env.Parse(AppSetting); err != nil {
		log.Fatalf("[FAIL] Failed to parse AppSetting: %v", err)
	}
	if err := env.Parse(ServerSetting); err != nil {
		log.Fatalf("[FAIL] Failed to parse ServerSetting: %v", err)
	}
}
