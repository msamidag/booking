package main

import (
	"net/http"

	"github.com/msamidag/booking/pkg/config"
	"github.com/msamidag/booking/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

//web oturumlarının devreye girdiği yer burasıdır

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	/*
	   middleware: (Ara katman)
	   Herhangi bir Mux için ara katman yazılımı,
	   bir eşleşme aranmadan önce yürütülür. Erken yanıt verme fırsatı
	   sağlayan belirli bir işleyiciye yönlendirme,
	   istek yürütmenin gidişatını değiştirin veya için istek kapsamlı
	   değerleri ayarlayın
	*/
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	//	mux.Use(WriteToConsole)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/teverpan", handlers.Repo.Teverpan)
	/*
		mux := pat.New()

		mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
		mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	*/
	return mux

}

//github.com/alexedwards/scs
/*
SCS: Go için HTTP Oturum Yönetimi
Ara katman yazılımı aracılığıyla oturum verilerinin otomatik olarak yüklenmesi ve kaydedilmesi.
PostgreSQL, MySQL, MSSQL, SQLite, Redis ve diğerleri dahil olmak üzere 19 farklı sunucu tarafı oturum mağazası seçeneği.
Özel oturum depoları da desteklenir.
İstek başına birden fazla oturumu, 'flash' mesajlarını, oturum belirteci yenilemeyi,
boşta ve mutlak oturum zaman aşımlarını ve 'beni hatırla' işlevini destekler.
Genişletilmesi ve özelleştirilmesi kolaydır. Oturum belirteçlerini HTTP üstbilgilerindeki veya
istek/yanıt gövdelerindeki istemcilere/istemcilerden iletin.
Verimli tasarım. Daha küçük, daha hızlı ve goril / seanslardan daha az bellek kullanır.
*/
