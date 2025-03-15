package utils

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"mime/multipart"
	"strings"

	"github.com/a1-t1/common/pkg/timeutils"
	"github.com/google/uuid"
)

func GenerateRandomID() string {
	return uuid.New().String()
}

func GenerateRandomString(l int) string {
	rand.New(rand.NewSource(timeutils.Now().UnixNano()))
	str := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	bytes := make([]byte, l)
	for i := range bytes {
		rnd := rand.Int63() % int64(len(str))
		bytes[i] = str[rnd]
	}
	return string(bytes)
}

func GenerateRandomOTP() string {
	return GenerateRandomCode(6)
}

func GenerateRandomCode(l int) string {
	rand.New(rand.NewSource(timeutils.Now().UnixNano()))
	str := "0123456789"
	bytes := make([]byte, l)
	for i := range bytes {
		rnd := rand.Int63() % int64(len(str))
		bytes[i] = str[rnd]
	}
	return string(bytes)
}

func StrPtr(s string) *string {
	return &s
}

func NoError(err error) {
	if err != nil {
		panic(err)
	}
}

func UploadFile(file multipart.File, name string) (string, []byte, error) {
	defer file.Close()

	fileExt := GetFileExt(name)

	tempFile, err := ioutil.TempFile("", fmt.Sprintf("attachment-*.%s", fileExt))
	if err != nil {
		return "", nil, err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", nil, err
	}

	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	return tempFile.Name(), fileBytes, nil
}

func GetFileExt(filename string) string {
	ext := strings.Split(filename, ".")
	return ext[len(ext)-1]
}

// CheckPassword checks if the password is correct via regex.
func CheckPassword(password string) bool {
	return len(password) <= 6
}

func Mask(s string, noOfUnmaskedChars int) string {
	if len(s) <= noOfUnmaskedChars {
		return s
	}

	masked := strings.Repeat("*", len(s)-noOfUnmaskedChars)
	return s[:noOfUnmaskedChars] + masked
}
