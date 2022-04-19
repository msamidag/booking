package handlers

import (
	"net/http"

	"github.com/msamidag/booking/pkg/config"
	"github.com/msamidag/booking/pkg/models"
	"github.com/msamidag/booking/pkg/render"
)

//TemplateData structının burada bulunması cylic paket hatasına yol açıyor

var Repo *Repository

//Repository, depo tipinde
type Repository struct {
	App *config.AppConfig
}

//NewRepo, yeni bir Repo oluşturur
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

//işleyiciler için Repo yu ayarlar
func NewHandlers(r *Repository) {
	Repo = r
}

//Home is the home page handler
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	//kullanıcı IP sini alınır
	remoteIP := r.RemoteAddr

	/*
		RemoteAddr, HTTP sunucularının ve diğer yazılımların,
		genellikle günlüğe kaydetme için isteği gönderen ağ adresini
		kaydetmesine izin verir. Bu alan ReadRequest tarafından
		doldurulmamıştır ve tanımlanmış bir formatı yoktur.
		Bu paketteki HTTP sunucusu, bir işleyiciyi çağırmadan önce
		RemoteAddr'ı bir "IP:port" adresine ayarlar.
		Bu alan HTTP istemcisi tarafından yok sayılır.
	*/
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP) // -->about
	/*
		Put, oturum verilerine bir anahtar ve karşılık gelen değeri ekler.
		Anahtar için mevcut herhangi bir değer değiştirilecektir.
		Oturum verileri durumu Değiştirildi olarak ayarlanacaktır.
	*/
	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})

}

//About is the about page handler
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	//ekledğimiz &TemplateData{} yı test edelim
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	//oturumdan her bilgiyi alabiliriz. açılır listede alabileceğimiz bilgiler sıralanmış.
	stringMap["remote_ip"] = remoteIP

	/*
		GetString, oturum verilerinden belirli bir anahtarın dize değerini
		döndürür. Bir dizgenin ("") sıfır değeri,
		anahtar yoksa veya değer bir dizgeye atanamadıysa döndürülür.
	*/
	//	m.app.session oturum açma kodları buraya
	//send data to template
	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
func (m *Repository) Teverpan(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Merhaba, teverpan sayfasındayız"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, "teverpan.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})

}
