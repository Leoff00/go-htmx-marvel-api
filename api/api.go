package api

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	envs "github.com/leoff00/go-marvel-api/env"
)

func calculateHash(ts int64, privateKey, publicKey string) string {
	h := md5.New()
	h.Write([]byte(fmt.Sprintf("%d%s%s", ts, privateKey, publicKey)))
	return hex.EncodeToString(h.Sum(nil))
}

func DoRequest(input string) []byte {
	timestamp := time.Now().Unix()
	envVars := envs.Getenv("env")

	hash := calculateHash(timestamp, envVars.Privkey, envVars.Pubkey)
	url := fmt.Sprintf("https://gateway.marvel.com/v1/public/characters?name=%s&apikey=%s&hash=%s&ts=%d", input, envVars.Pubkey, hash, timestamp)

	resp, err := http.Get(url)

	if err != nil {
		log.Println("Error During Request...", err)
		return nil
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Error During Parse Response body...", err)
		return nil
	}

	return body
}
