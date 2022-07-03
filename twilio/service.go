package twilio

import (
	"html/template"
	"net/http"

	"go.uber.org/zap"
)

type TwilioService struct {
	Logger    *zap.SugaredLogger
	router    http.Handler
	TwilioCon TwilioConfig
	Template  *template.Template
}

func (ts *TwilioService) Init() {
	ts.router = ts.routes()
}

func (ts TwilioService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ts.router.ServeHTTP(w, r)
}
