package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"math"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

const (
	// WIB :
	WIB string = "Asia/Jakarta"
	// UTC :
	UTC string = "UTC"

	ALPHABET = "abcdefghijklmnopqrstuvwxyz"
	NUMBER   = "0123456789"
)

// GetTimeNow :
func GetTimeNow() time.Time {
	return time.Now().In(GetLocation())
}

// GetLocation - get location wib
func GetLocation() *time.Location {
	return time.FixedZone(WIB, 7*3600)
}

// Stringify :
func Stringify(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}

func CheckEmail(e string) bool {
	var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if len(e) < 3 && len(e) > 254 {
		return false
	}

	return emailRegex.MatchString(e)
}

func GenerateNumber(n int) string {
	rand.Seed(time.Now().UnixNano())
	var letterRunes = []rune("0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func GenerateString(n int) string {
	var sb strings.Builder
	k := len(ALPHABET)
	for i := 0; i < n; i++ {
		c := ALPHABET[rand.Intn(k)]
		sb.WriteByte(c)
	}
	return sb.String()
}

func RandomName() string {
	return fmt.Sprintf("%s %s", GenerateString(5), GenerateString(6))
}

func RandomUserName() string {
	return fmt.Sprint(GenerateString(5))
}

func RandomPassword() string {
	return fmt.Sprint(GenerateString(5))
}

func RandomPhone() string {
	return fmt.Sprintf("+%s%s", GenerateNumber(2), GenerateNumber(8))
}

func RandomEmail() string {
	return fmt.Sprintf("%s@mail.com", GenerateString(6))
}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}

func Float64bytes(float float64) []byte {
	bits := math.Float64bits(float)
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, bits)
	return bytes
}

func GetDayOfBirth(year, month, day int, format string) (time.Time, string) {
	dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
	str_dob := fmt.Sprintf("%v", dob.Format(format))
	return dob, str_dob
}
