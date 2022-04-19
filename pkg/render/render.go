package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/msamidag/booking/pkg/config"
	"github.com/msamidag/booking/pkg/models"
)

//FuncMap, ADLARDAN İŞLEVLERE EŞLEMEYİ TANIMLAYAN BİR HARİTA TÜRÜDÜR.
//FuncMap (şablon tutucu da diyebiliriz burada)
var functions = template.FuncMap{}

var app *config.AppConfig

//NewTemplates, şablon paketinin yapılandırmasını ayarlar
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {

	return td
}

//RenderTemplate renders template using html/template
func RenderTemplate(w http.ResponseWriter, tmpl string, td *models.TemplateData) {
	//öncelikli olarak, şablonu önbelleğe alıyoruz. ana şablonumuz base.layout.tmpl
	var tc map[string]*template.Template

	if app.UseCache {
		// get the template cache from the app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()

	}
	t, ok := tc[tmpl]
	if !ok {
		log.Fatal("şablon önbelleğinden şablon alınamadı")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//base şablonunun ön belleğe alınması FuncMap()
	myCache := map[string]*template.Template{}

	//  Glob, desenle eşleşen tüm dosyaların ADLARINI string tipli slice olarak döndürür veya
	//  eşleşen dosya yoksa sıfırdır.
	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return myCache, err
	}
	//sayfa adlarına göre parse işlemi
	for _, page := range pages {
		name := filepath.Base(page) //sayfa adı sıyrılıyor
		//yeni bir HTML şablonu oluşturulur (New),
		//şablon tutucuya haritaya aktarılır (Funcs),
		//adlandırılmış dosyalar ayrıştırılıp ts ile ilişkilendirilir (ParseFile)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}
		/*
			ParseGlob, kalıp tarafından tanımlanan dosyalardaki
			şablon tanımlarını ayrıştırır ve elde edilen şablonları
			 matches ile ilişkilendirir.
		*/
		matches, err := filepath.Glob("./templates/*.layout.tmpl")
		if err != nil {
			return myCache, err
		}
		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}
	return myCache, nil
}

//<form action="/"><button type="submit">Home Page</button></form>
//<form action="/about"><button type="submit">Home Page</button></form>
