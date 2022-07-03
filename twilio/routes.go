package twilio

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (ts TwilioService) defaultMain(w http.ResponseWriter, r *http.Request) {
	ts.Logger.Infow("default main page")

	tmp, err := template.ParseFiles("index.html")

	if err != nil {
		ts.Logger.Error(err)
	}
	w.Header().Add("Content Type", "text/html")

	tmp.Execute(w, nil)

}

func (ts TwilioService) handleMessage(w http.ResponseWriter, r *http.Request) {
	ts.Logger.Infow("handle sending message")

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		ts.Logger.Infow("header is not a form urlencoded")
		return
	}

	err := r.ParseForm()
	if err != nil {
		ts.Logger.Error(err)
	}

	message := r.Form.Get("message")
	to := r.Form.Get("to")

	if to == "" || message == "" {
		ts.Logger.Infow("missing message or phone number")
		return
	}

	ts.TwilioCon.SendMessage(to, message)
	fmt.Fprintf(w, "hello, %s!\n", "message is sending")
}

func (ts *TwilioService) routes() http.Handler {
	r := httprouter.New()

	r.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		ts.defaultMain(writer, request)
	})

	r.POST("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		ts.handleMessage(writer, request)
	})

	return r
}
