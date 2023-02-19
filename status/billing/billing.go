package billing

import (
	"bufio"
	"log"
	"main/config"
	"main/structs"
	"math"
	"os"
	"strconv"
)

func GetBilling(billingChan chan structs.BillingData) {
	billingData, err := os.Open(config.BillingDataFile)
	if err != nil {
		log.Println("Не удалось открыть файл Data", err)
		billingChan <- structs.BillingData{}
		return
	}
	defer billingData.Close()

	scanner := bufio.NewScanner(billingData)
	var field structs.BillingData
	for scanner.Scan() {
		line := scanner.Text()

		var newSlice []int
		for i := len(line); i > 0; i-- {
			circleInt, _ := strconv.Atoi(line[i-1:])
			newSlice = append(newSlice, circleInt)
			line = line[:i-1]
		}

		var number uint8
		for i := 0; i < len(newSlice); i++ {
			number = number + uint8(newSlice[i])*uint8(math.Pow(2, float64(i)))
		}

		CreateCustomer := number&1 == 1
		Purchase := number>>1&1 == 1
		Payout := number>>2&1 == 1
		Recurring := number>>3&1 == 1
		FraudControl := number>>4&1 == 1
		CheckoutPage := number>>5&1 == 1

		field = structs.BillingData{
			CreateCustomer: CreateCustomer,
			Purchase:       Purchase,
			Payout:         Payout,
			Recurring:      Recurring,
			FraudControl:   FraudControl,
			CheckoutPage:   CheckoutPage,
		}
	}
	billingChan <- field
}
