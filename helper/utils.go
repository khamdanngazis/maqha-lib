package helper

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"math/big"

	"github.com/nfnt/resize"
	"golang.org/x/crypto/bcrypt"
)

// GenerateRandomString generates a random string of a specified length.
func GenerateRandomString(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		idx, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		b[i] = charset[idx.Int64()]
	}
	return string(b), nil
}

func HashPassword(password string) (string, error) {
	// Generate a hashed representation of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func CompareHashAndPassword(storedHash string, password string) error {
	// Convert the stored hash from string to []byte
	hashBytes := []byte(storedHash)

	// Compare the stored hash with the password
	return bcrypt.CompareHashAndPassword(hashBytes, []byte(password))
}

func GenerateBcryptSalt() (string, error) {
	// Generate 16 random bytes for the salt
	saltBytes := make([]byte, 16)
	_, err := rand.Read(saltBytes)
	if err != nil {
		return "", err
	}

	// Encode the random bytes to base64 to get a string representation
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	// Prefix with the bcrypt identifier and cost factor
	return fmt.Sprintf("$2a$12$%s", salt), nil
}

func CompressImage(img image.Image, format string) ([]byte, error) {
	// Resize image to ensure it's under 1MB
	resizedImg := resize.Resize(800, 0, img, resize.Lanczos3)

	var buffer bytes.Buffer
	switch format {
	case "jpeg", "jpg":
		err := jpeg.Encode(&buffer, resizedImg, nil)
		if err != nil {
			return nil, err
		}
	case "png":
		err := png.Encode(&buffer, resizedImg)
		if err != nil {
			return nil, err
		}
	case "gif":
		// Optionally, you could support GIFs, though resize.Lanczos3 may not be ideal for GIFs
		// For simplicity, let's return an error for unsupported formats
		return nil, fmt.Errorf("GIF format not supported for compression")
	default:
		return nil, fmt.Errorf("unsupported image format: %s", format)
	}

	if buffer.Len() > 1<<20 { // 1MB in bytes
		return nil, fmt.Errorf("compressed image exceeds 1MB")
	}

	return buffer.Bytes(), nil
}
