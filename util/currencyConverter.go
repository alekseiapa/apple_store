package util

func ConvertCur(fromCur string, toCur string, amount float32) float32 {
	var amountConverted float32
	if fromCur == "USD" && fromCur == toCur {
		return amount
	}
	if fromCur == "USD" {
		switch toCur {
		case "RUB":
			amountConverted = amount * 62.5
		case "EUR":
			amountConverted = amount * 0.95
		}
	} else {
		switch fromCur {
		case "RUB":
			amountConverted = amount * 0.016
		case "EUR":
			amountConverted = amount * 1.05
		}
	}
	return amountConverted
}
