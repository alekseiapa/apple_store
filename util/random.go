package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnoprstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// random int between min and max
func RandomInt(min, max int) int {
	return min + rand.Intn(max-min+1) // min + 0->max-min
}

// random float between min and max
func RandomFloat(min, max float32) float32 {
	return min + rand.Float32()*(max-min) // min + 0->max-min
}

// random string of n length
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()

}

// UserTable
func RandomUserFirstName() string {
	return RandomString(6)
}

func RandomUserLastName() string {
	return RandomString(6)

}

func RandomUserMiddleName() string {
	return RandomString(6)

}

func RandomUserGender() string {
	return RandomString(1)

}

func RandomUserAge() int {
	return RandomInt(1, 20)
}

func RandomUserBalance() float32 {
	return RandomFloat(200.00, 1000.00)
}

// Product Table
func RandomProductDescription() string {
	return RandomString(6)
}

func RandomProductPrice() float32 {
	return RandomFloat(10.00, 20.00)

}

func RandomProductInStock() int32 {
	return int32(RandomInt(1, 20))

}

// Order Table
func RandomOrderUseruuid() int64 {
	return int64(RandomInt(1, 20))
}
func RandomOrderQuantity() int64 {
	return int64(RandomInt(1, 5))
}

// OrderProduct Table
func RandomOrderUuid() int64 {
	return int64(RandomInt(1, 20))
}

func RandomProductUuid() int64 {
	return int64(RandomInt(1, 20))
}
