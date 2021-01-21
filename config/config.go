package configuration

type Config struct {
	Port     string `json:"port"`
	DBName   string `json:"dbName"`
	CollName string `json:"collName"`
}
