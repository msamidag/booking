/*
gitbashta ana klasörde iken
$ git init
$ git add .
$ git commit -m "initial commit"
$

Klasörün silinmesi git deponuzda sorunlara neden olabilir.
Tüm yürütme geçmişinizi silmek, ancak kodu geçerli durumunda tutmak istiyorsanız,
bunu aşağıdaki gibi yapmak çok güvenlidir: .git Checkout
		git checkout --orphan latest_branch
Add all the files
		git add -A
Commit the changes
		git commit -am "commit message"
Delete the branch
		git branch -D master
Rename the current branch to master
		git branch -m master
Finally, force update your repository
		git push -f origin master
Not: Bu, eski taahhüt geçmişinizi etrafta tutmayacak

*/
//github.com/msamidag/booking
package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	//	"webrender/cmd/web/routes"
	"github.com/msamidag/booking/pkg/config"
	"github.com/msamidag/booking/pkg/handlers"
	"github.com/msamidag/booking/pkg/render"

	//"webrender/cmd/web/routers"
	"github.com/alexedwards/scs/v2"
)

const port = ":9000"

var app config.AppConfig //şablon oluşturmak için
var session *scs.SessionManager

//iki farklı session var. buna değişken gölgeleme denir. Ve düzeltilmesi zor bir sorundur

//main is the main function
func main() {

	//Production 'da iken bunu true olarak değiştirin
	app.InProduction = false

	session = scs.New()                            //yeni oturum aç
	session.Lifetime = 24 * time.Hour              //7-24 çalışsın
	session.Cookie.Persist = true                  //çerz kabul et
	session.Cookie.SameSite = http.SameSiteLaxMode //benzer sayfalarda esnek ol
	session.Cookie.Secure = app.InProduction       //çerez güvenliğini gözardı et gibi gibi
	//bu oturum bilgilerini nerede kullanırız? elbetteki işleyicilerde (handlers)

	app.Session = session //Session ile session her ikisi de scs paketinin oturum açma yönetimini belirtir.
	//ancak, Session AppConfig yapısının bir alanı iken, session oturum açma seçeneklerini belirlemek için burada tanımlanmış değişkendir.
	//session özellikleri burada belirlenip uygulamanın Session alanına atanır. buna gömme işlemi denir

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app) //uygulamamıza, render bileşenlerine erişim izni alınır

	fmt.Printf("Starting application on port: %s\n", port)
	//	http.ListenAndServe(port, nil) //_ hata varsa da umurumda değil anlamındadır

	srv := &http.Server{
		Addr:    port,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

// HTML şablonu için: Ctrl+Shift+p --> Tmpl html: Create HTML Template
