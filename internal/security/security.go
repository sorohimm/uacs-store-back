package security

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"golang.org/x/crypto/bcrypt"
)

func GenerateSalt(password string) string {
	return Base64(16)
}

func SaltPassword(password, salt string) string {
	return password + salt
}

// HashPassword hash password
func HashPassword(password string) (string, error) {
	// Convert password string to byte slice
	var passwordBytes = []byte(password)

	// Hash password with Bcrypt's min cost
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	return string(hashedPasswordBytes), err
}

// DoPasswordsMatch Check if two passwords match using Bcrypt's CompareHashAndPassword
// which return nil on success and an error on failure.
func DoPasswordsMatch(hashedPassword, currPassword string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(hashedPassword), []byte(currPassword))
	return err == nil
}

// Bytes generates n random bytes
func Bytes(n int) []byte {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return b
}

// Base64 generates a random base64 string with length of n
func Base64(n int) string {
	return String(n, "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ+/")
}

// list of default letters that can be used to make a random string when calling String
// function with no letters provided
var defLetters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

// String generates a random string using only letters provided in the letter's parameter
// if user ommit letters parameters, this function will use defLetters instead
func String(n int, letters ...string) string {
	var letterRunes []rune
	if len(letters) == 0 {
		letterRunes = defLetters
	} else {
		letterRunes = []rune(letters[0])
	}

	var bb bytes.Buffer
	bb.Grow(n)
	l := uint32(len(letterRunes))
	// on each loop, generate one random rune and append to output
	for i := 0; i < n; i++ {
		bb.WriteRune(letterRunes[binary.BigEndian.Uint32(Bytes(4))%l])
	}
	return bb.String()
}
