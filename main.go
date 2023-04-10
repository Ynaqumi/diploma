package main

import (
	"diploma/status/billing"
	"diploma/status/email"
	"diploma/status/incidents"
	"diploma/status/mms"
	"diploma/status/sms"
	"diploma/status/support"
	"diploma/status/voice"
	"fmt"
)

func main() {
	fmt.Println(sms.Sms(), "\n",
		mms.Mms(), "\n",
		voice.VoiceCall(), "\n",
		email.Email(), "\n",
		billing.Billing(), "\n",
		support.Support(), "\n",
		incidents.Incidents())
}
