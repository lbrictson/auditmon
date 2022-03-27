package auth

import (
	"bytes"
	"image/png"
	"math/rand"
	"time"

	"github.com/pquerna/otp/totp"
	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt takes in a plan text password and returns the hashed and salted value
func HashAndSalt(plainTextPassword string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(plainTextPassword), bcrypt.DefaultCost)
	hashedString := string(hashed)
	return hashedString
}

// ComparePassword compares a plain text password with an encrypted password, returning true if they match
func ComparePassword(plainTextPassword string, encryptedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(plainTextPassword))
	if err != nil {
		// Sleep for a random amount of time to prevent password timing attacks
		rand.Seed(time.Now().UnixNano())
		n := rand.Intn(100)
		time.Sleep(time.Duration(n) * time.Millisecond)
		return false
	}
	return true
}

type NewTOTPOutput struct {
	ImageBytes []byte
	Secret     string
}

func GenerateMFA(issuer string, username string) NewTOTPOutput {
	key, _ := totp.Generate(totp.GenerateOpts{
		Issuer:      issuer,
		AccountName: username,
	})
	img, _ := key.Image(200, 200)
	var buf bytes.Buffer
	png.Encode(&buf, img)
	return NewTOTPOutput{
		ImageBytes: buf.Bytes(),
		Secret:     key.Secret(),
	}
}

func ValidateMFA(userSuppliedCode string, secret string) bool {
	return totp.Validate(userSuppliedCode, secret)
}
