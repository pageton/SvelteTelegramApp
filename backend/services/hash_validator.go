package services

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

func ParseHash(hash string) (map[string]string, error) {
	values, err := url.ParseQuery(hash)
	if err != nil {
		return nil, err
	}

	data := make(map[string]string)
	for key, value := range values {
		if len(value) > 0 {
			data[key] = value[0]
		}
	}
	return data, nil
}

func IsHashValid(data map[string]string, botToken string) (bool, error) {
	keys := make([]string, 0, len(data))
	for key := range data {
		if key != "hash" {
			keys = append(keys, key)
		}
	}
	sort.Strings(keys)

	var checkString strings.Builder
	for i, key := range keys {
		if i > 0 {
			checkString.WriteString("\n")
		}
		checkString.WriteString(key + "=" + data[key])
	}

	secret := hmac.New(sha256.New, []byte("WebAppData"))
	secret.Write([]byte(botToken))
	secretKey := secret.Sum(nil)

	h := hmac.New(sha256.New, secretKey)
	h.Write([]byte(checkString.String()))
	signature := h.Sum(nil)

	hexSignature := hex.EncodeToString(signature)
	return data["hash"] == hexSignature, nil
}
