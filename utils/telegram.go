package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"os"
	"sort"
	"strings"
)

func VerifyTelegramWebApp(query url.Values) error {
	fmt.Println("entered: vapp")
	botToken := os.Getenv("BOT_TOKEN")
	if botToken == "" {
		fmt.Println("no token")
		return errors.New("BOT_TOKEN not set")
	}

	receivedHash := query.Get("hash")
	if receivedHash == "" {
		fmt.Println("no hash")
		return errors.New("missing hash")
	}

	// Step 1: Remove 'hash' and sort remaining keys
	dataCheck := make([]string, 0)
	for key, values := range query {
		if key == "hash" {
			continue
		}
		// Assume 1 value per key
		dataCheck = append(dataCheck, key+"="+values[0])
	}
	sort.Strings(dataCheck)
	dataCheckString := strings.Join(dataCheck, "\n")
	fmt.Println("sorted:", dataCheckString)
	// Step 2: Create HMAC using the secret key
	sha256BotToken := sha256.Sum256([]byte(botToken))
	secretKey := sha256BotToken[:]
	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(dataCheckString))
	calculatedHash := hex.EncodeToString(h.Sum(nil))
	fmt.Println(calculatedHash)
	// Step 3: Compare hashes
	if calculatedHash != receivedHash {
		return errors.New("invalid hash")
	}

	return nil
}
