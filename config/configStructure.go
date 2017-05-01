package config

//Config represent the yaml config file
type Config struct {
	Logging struct {
		Destination string
	}
	Prometheus struct {
		Address string
	}
}
