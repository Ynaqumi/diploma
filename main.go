package main

import (
	"diploma2/status/billing"
	"diploma2/status/email"
	"diploma2/status/incidents"
	"diploma2/status/mms"
	"diploma2/status/sms"
	"diploma2/status/support"
	"diploma2/status/voice"
)

func main() {
	sms.Sms()
	mms.Mms()
	voice.VoiceCall()
	email.Email()
	billing.Billing()
	support.Support()
	incidents.Incidents()
}
