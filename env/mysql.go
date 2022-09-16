package env

import "os"

func MySQLUser() string {
	return os.Getenv("MYSQL_USER")
}
func MySQLPassword() string {
	return os.Getenv("MYSQL_PASSWORD")
}
func MySQLHost() string {
	return os.Getenv("MYSQL_HOST")
}
func MySQLPort() string {
	return os.Getenv("MYSQL_PORT")
}
func MySQLDatabase() string {
	return os.Getenv("MYSQL_DATABASE")
}
