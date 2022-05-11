package config

type Config struct {
	Port     uint32   `json:"port"`
	TaskPath string   `json:"taskPath"`
	Bootlist []string `json:"bootlist"`
}
