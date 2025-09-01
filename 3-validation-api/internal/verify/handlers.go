package verify

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"purple2/3-validation-api/configs"
	"purple2/3-validation-api/pkg/res"

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
		err := e.Send(handler.Config.Address, smtp.PlainAuth("", handler.Config.Email, handler.Config.Password, handler.Config.Address))
		if err != nil {
			log.Printf("Ошибка отправки письма: %v", err)
		}
		//res.Json(w, data, 200)
	}
}

func (handler *VarifyHandler) Verify() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		verCode := req.FormValue("verCode")
		applyCode := "123456" // пример получения кода из базы данных
		if verCode != applyCode {
			res.Json(w, nil, 300)
		}
		res.Json(w, nil, 200)
		
		fmt.Println("Verify")
	}
}

