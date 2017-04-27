package config

type Config struct {
	Logging struct {
		Destination string
	}
	Prometheus struct {
		Address string
	}
}
