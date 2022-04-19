package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

//var app config.AppConfig //şablon oluşturmak için
//ara katman fonksiyonuna örnek
func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page (sayfaya tıkla)")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	//yeni bir çerez (cookie) oluşturuyoruz
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,                 //sadece http olsun, o yol için doğru diyelim, çerez yolu
		Path:     "/",                  // "/" ile bir çerez güvenliği için tüm siteye atıfta bulunma şekli
		Secure:   app.InProduction,     //false diyoruz, çünkü şu anda HTTP 'de çalışmıyoruz. Ancak üretimde, aslında, bunu doğru ve aynı siteye değiştirin
		SameSite: http.SameSiteLaxMode, //bunun için aynı site, relax mod olan in standardını kullanacağız
	})
	//temel çerezi ayarlamamız gerekiyor, çünkü oluşturduğu belirtecin
	//kullanılabilir olduğundan emin olmak için çerezleri kullanıyor.
	return csrfHandler
}

// SessionLoad loads and saves session data for current request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// nosurf --------------------------------------------------------
//github.com/justinas/nosurf ... bu paket
/*
nosurf , Go için Siteler Arası İstek Sahteciliği saldırılarını önlemenize yardımcı olan bir HTTP paketidir.
Bir ara yazılım gibi davranır ve bu nedenle temel olarak herhangi bir Go HTTP uygulamasıyla uyumludur.

CSRF önemli bir güvenlik açığı olmasına rağmen, Go'nun web ile ilgili paket altyapısı çoğunlukla CSRF denetimlerini uygulamayan
ve uygulamaması gereken mikro çerçevelerden oluşur.

nosurf, güvenli olmayan her (GET / HEAD / OPTIONS / TRACE) yönteminde CSRF saldırılarını sarmalayan
ve kontrol eden bir yöntem sağlayarak bu sorunu çözer. CSRFHandler http.Handler

nosurf, Go 1.1 veya üstünü gerektirir.
*/
