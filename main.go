package main

import (
	"diploma/internal/data/billing"
	"diploma/internal/data/email"
	"diploma/internal/data/incidents"
	"diploma/internal/data/mms"
	"diploma/internal/data/sms"
	"diploma/internal/data/support"
	"diploma/internal/data/voice"
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
