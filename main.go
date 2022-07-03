package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"

	"web.dev.learning/twilio"
)

func run() (err error) {
	flags := parseFlags()
	closeLogger, logger := newLogger()
	defer closeLogger()

	twilioCon := twilio.TwilioConfig{
		AccountSid:       flags.TwilioAccountSid,
		AccountToken:     flags.TwilioAuthToken,
		TwilioFromNumber: flags.TwilioFromNumber,
	}

	service := twilio.TwilioService{
		Logger:    logger,
		TwilioCon: twilioCon,
	}
	service.Init()

	logger.Infow("parse flag result", "twilioFromNumber", flags.TwilioFromNumber)

	logger.Infow("service listening", "host", flags.Host, "port", flags.Port)
	return http.ListenAndServe(flags.Host+":"+flags.Port, service)
}

func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func newLogger() (func(), *zap.SugaredLogger) {
	logger, _ := zap.NewProduction()
	closeLogger := func() {
		logger.Sync()
	}
	sugar := logger.Sugar()
	return closeLogger, sugar
}

type Flags struct {
	Host             string
	Port             string
	Env              string
	TwilioAccountSid string
	TwilioAuthToken  string
	TwilioFromNumber string
}

func parseFlags() Flags {
	hostPtr := flag.String("host", "0.0.0.0", "Host of this service")
	portPtr := flag.String("port", "8080", "Port of this service")

	twilioAccountSid := flag.String("twilioAccountSid", "placeholder", "Twilio account sid")
	twilioAuthToken := flag.String("twilioAuthToken", "placeholder", "Twilio account token")
	twilioFromNumber := flag.String("twilioFromNumber", "+123456790", "Twilio send message phone number")

	env := flag.String("env", "dev", "Environment (dev, prod)")

	flag.Parse()

	return Flags{
		Port:             *portPtr,
		Host:             *hostPtr,
		TwilioAccountSid: *twilioAccountSid,
		TwilioAuthToken:  *twilioAuthToken,
		TwilioFromNumber: *twilioFromNumber,
		Env:              *env,
	}
}
