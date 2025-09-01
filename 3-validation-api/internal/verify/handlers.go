package verify

import (
	"fmt"
	"net/http"
	"net/smtp"
	"purple2/3-validation-api/configs"

	"github.com/jordan-wright/email"
)

type VarifyHandler struct {
	Config *configs.Config
}

func NewVarifyHandler(router *http.ServeMux, config *configs.Config) {
	handler := &VarifyHandler{
		Config: config,
	}
	router.HandleFunc("POST /send", handler.Send())
	router.HandleFunc("GET /verify/{hash}", handler.Verify())
}

func (handler *VarifyHandler) Send() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		
		toEmail := req.FormValue("toEmail") // пример получения получателя

		e := email.NewEmail()
		e.From = fmt.Sprintf("Отправитель <%s>", handler.Config.Email)
		e.To = []string{toEmail}
		e.Subject = "Проверка почтового адреса"
		e.Text = []byte("Тело письма с текстом....")
		e.HTML = []byte("Тело письма HTML с ссылкой для восстановления")
		e.Send(handler.Config.Address, smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, handler.Config.Address))
		//res.Json(w, data, 200)
	}
}

func (handler *VarifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("Verify")
	}
}

