package util

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int32) int32 {
	return min + rand.Int31n(max-min+1)
}

func RandomDouble(min, max int64) float64 {
	myRange := max - min + 1

	tempRes := fmt.Sprintf("%2.2f", float32(myRange)*rand.Float32())
	res, err := strconv.ParseFloat(tempRes, 64)
	if err != nil {
		fmt.Println("convert double is fail: ", err)
		return 10.99
	}
	return res
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)
	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomMoney(min, max int64) float64 {
	return RandomDouble(min, max)
}

func RandomStock(min, max int32) int32 {
	return RandomInt(min, max)
}

func RandomAddress() string {
	return RandomString(10)
}

func RandomName() string {
	return RandomString(4)
}
