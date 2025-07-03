package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

func VerifyTelegramWebApp(initData string) bool {

	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		fmt.Println("BOT_TOKEN not set")
		return false
	}

	values, err := url.ParseQuery(initData)
	if err != nil {
		fmt.Println(err)

		return false
	}

	// Extract hash and remove it
	receivedHash := values.Get("hash")
	values.Del("hash")

	// Get sorted keys
	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// Build data_check_string
	var builder strings.Builder
	for i, k := range keys {
		if i > 0 {
			builder.WriteByte('\n')
		}
		// Preserve original URL-encoded values
		builder.WriteString(k)
		builder.WriteByte('=')
		builder.WriteString(values.Get(k))
	}
	dataCheckStr := builder.String()
	// Compute secret key
	mac := hmac.New(sha256.New, []byte("WebAppData"))
	mac.Write([]byte(botToken))
	secretKey := mac.Sum(nil)

	// Compute expected hash
	mac = hmac.New(sha256.New, secretKey)
	mac.Write([]byte(dataCheckStr))
	expectedHash := mac.Sum(nil)
	expectedHashHex := hex.EncodeToString(expectedHash)
	// Constant-time comparison
	return hmac.Equal([]byte(expectedHashHex), []byte(strings.ToLower(receivedHash)))
}
