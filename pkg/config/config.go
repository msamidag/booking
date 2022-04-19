package config

import (
	"html/template"
	"log"

	"github.com/alexedwards/scs/v2"
)

//AppConfig uygulama yapılandırmasını tutar
type AppConfig struct {
	UseCache      bool
	TemplateCache map[string]*template.Template
	InfoLog       *log.Logger //daha sonra kullanılacak. günlük gibi yaoılan şeyleri herşeye kaydedebiliecek. veritabanı gibi
	InProduction  bool
	Session       *scs.SessionManager
}
