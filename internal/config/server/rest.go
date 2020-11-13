package server

type ServerList struct {
	Rest Rest
}

// Rest object rest
type Rest struct {
	Tracking Server `yaml:"tracking"`
}

type Server struct {
	URL          string `yaml:"url"`
	ClientID     string `yaml:"clientID"`
	ClientSecret string `yaml:"clientSecret"`
	ChannelID    string `yaml:"channelID"`
}
