package envs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvVars struct {
	Pubkey  string
	Privkey string
}

func Getenv(envFile string) *EnvVars {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot read env files...")
	}

	envVars := &EnvVars{
		Pubkey:  os.Getenv("PUB_KEY"),
		Privkey: os.Getenv("PRIV_KEY"),
	}

	return envVars

}
