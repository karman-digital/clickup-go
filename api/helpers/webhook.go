package helpers

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"hash"

	cuerrors "github.com/karman-digital/clickup/app/errors"
)

func ValidateWebhookSignature(method, signature string, secret, body []byte) error {
	checkSum, err := toCompactJSONString(body)
	if err != nil {
		return err
	}
	hash := encrypt(secret, checkSum)
	if signature != hex.EncodeToString(hash.Sum(nil)) {
		return cuerrors.ErrInvalidSignature
	}
	return nil
}

func encrypt(secret []byte, input string) hash.Hash {
	hash := hmac.New(sha256.New, secret)
	hash.Write([]byte(input))
	return hash
}

func toCompactJSONString(input []byte) (string, error) {
	buffer := new(bytes.Buffer)
	if err := json.Compact(buffer, input); err != nil {
		return "", err
	}
	return buffer.String(), nil
}
