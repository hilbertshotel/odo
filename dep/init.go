package dep

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

// PLACE ALL DEPENDENCIES IN THIS FUNCTION
func Init() (*Dependencies, error) {

	// 1. logger
	log := initLogger()

	// 2. config
	cfg, err := loadConfig()
	if err != nil {
		return nil, err
	}

	// 3. template
	tmp, err := template.ParseGlob(cfg.TmpDir)
	if err != nil {
		return nil, err
	}

	// 4. handlers
	handlers, err := loadHandlers(cfg.GamesDir)
	if err != nil {
		return nil, err
	}

	return &Dependencies{
		Log:      log,
		Cfg:      cfg,
		Tmp:      tmp,
		Handlers: handlers,
	}, nil
}

// INITIALIZE LOGGER
func initLogger() *Logger {
	okLog := log.New(os.Stdout, "OK ", log.Ldate|log.Ltime)
	errLog := log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime)

	return &Logger{
		Error: func(err error) {
			_, file, line, _ := runtime.Caller(1)
			msg := fmt.Sprintf("(%v %v) %v", filepath.Base(file), line, err)
			errLog.Println(msg)
		},

		Ok: func(ss ...string) {
			var msg string
			for _, s := range ss {
				msg += s + " "
			}
			okLog.Println(msg[:len(msg)-1])
		},
	}
}

// LOAD CONFIG
func loadConfig() (*Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var cfg Config
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

// LOAD HANDLERS LIST
func loadHandlers(dir string) (Handlers, error) {
	folders, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var handlers Handlers
	for _, folder := range folders {
		name := "/" + folder.Name() + "/"

		handler := Handler{
			Name: name,
			Dir:  dir + name,
		}

		handlers = append(handlers, handler)
	}

	return handlers, nil
}
