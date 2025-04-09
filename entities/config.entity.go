package entities

type TraderBotConfig struct {
	Database   TraderBotDatabase    `mapstructure:"database"`
	Redis      TraderBotRedis       `mapstructure:"redis"`
	Indicators []TraderBotIndicator `mapstructure:"indicators"`
}

type TraderBotDatabase struct {
	URL string `mapstructure:"url"`
}

type TraderBotRedis struct {
	Addr     string `mapstructure:"addr"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type TraderBotIndicator struct {
	ID     int     `mapstructure:"id"`
	Name   string  `mapstructure:"name"`
	TP     float64 `mapstructure:"tp"`
	SL     float64 `mapstructure:"sl"`
	Volume float64 `mapstructure:"volume"`
}
