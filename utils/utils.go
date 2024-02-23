package utils

import (
	"crypto/rand"
	"encoding/base64"
)

// GenerateRandomString menghasilkan string acak sepanjang n byte.
func GenerateRandomString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	// Menggunakan base64 untuk mengkodekan sehingga outputnya aman untuk URL dan panjangnya lebih pendek
	return base64.URLEncoding.EncodeToString(b)
}
