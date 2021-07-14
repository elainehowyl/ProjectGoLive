package envfile

import (
	"github.com/joho/godotenv"
)

func RetrieveEnv(key string) string {
	var envs map[string]string
	envs, _ = godotenv.Read("./.env")
	myKey := envs[key]
	return myKey
}
