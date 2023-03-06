package dep

import "html/template"

// LOGGER
type Logger struct {
	Error func(error)
	Ok    func(...string)
}

// CONFIG
type Config struct {
	HostAddr     string
	ReadTimeout  int
	WriteTimeout int
	TmpDir       string
	GamesDir     string
}

// HANDLERS
type Handler struct {
	Name string
	Dir  string
}

type Handlers []Handler

// DEPENDENCIES
type Dependencies struct {
	Log      *Logger
	Cfg      *Config
	Tmp      *template.Template
	Handlers Handlers
}
