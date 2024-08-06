package cli

type Config struct {
	Next     *string
	Previous *string
}

func NewConfig() Config {
	return Config{}
}

func (cfg *Config) setNext(next *string) {
	if next == nil {
		return
	}

	cfg.Next = next
}
