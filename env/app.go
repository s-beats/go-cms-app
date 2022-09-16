package env

import "os"

func HostName() string {
	return os.Getenv("HOSTNAME")
}

func Port() string {
	return os.Getenv("PORT")
}

func Envcode() string {
	return os.Getenv("ENVCODE")
}
