package config

type Config struct {
	Logging struct {
		Logfile string
		Level   string
	}
	Prometheus struct {
		Address string
	}
}
