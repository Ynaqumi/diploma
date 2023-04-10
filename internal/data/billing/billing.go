package billing

import (
	"diploma/internal/structs"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

var numder uint8

func Billing() (billing []structs.BillingData) {
	data, err := os.ReadFile("simulator/billing.data")
	if err != nil {
		log.Printf("Не удалось прочитать файл billing.data. Ошибка: %v", err)
	}

	line := strings.Split(string(data), "")
	elem := make([]bool, len(line))

	for i := len(line) - 1; i >= 0; i-- {
		if err != nil {
			fmt.Println(err)
		}
		if line[i] == "0" {
			elem[i] = false
		} else if line[i] == "1" {
			elem[i] = true
			numder += uint8(math.Pow(2, float64(i)))
		}
	}
	billing = append(billing, structs.BillingData{CreateCustomer: elem[0], Purchase: elem[1], Payout: elem[2], Recurring: elem[3], FraudControl: elem[4], CheckoutPage: elem[5]})
	return
}
