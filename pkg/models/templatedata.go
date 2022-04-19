package models

//TemplateData, işleyicilerden şablona gönderilen verileri tutar
type TemplateData struct {
	StringMap map[string]string      //email vs gibi veriler için
	IntMap    map[string]int         //keyi string valuesi int olan veriler için
	FloatMap  map[string]float64     //keyi string valuesi float64 olan veriler için
	Data      map[string]interface{} //ne olduğundan emin olamadığımız veri türleri için
	CSRFToken string                 //her sayfanın güvenliği için kullanılacak. CSR, CSR belirteci(token) siteler arası sahteciliği belirtici anlamına gelir. ara katmanda belirlenir
	Flash     string                 //başarılı işlemler için mesaj olarak kullanılacak
	Warning   string                 //uyarı mesajları için kullanılacak
	Error     string                 //hatalar için kullanılacak

}
