package variables

type ClientConfig struct {
	Host        string
	Prefix      string
	TopicPrefix string
}

type Config struct {
	PrefixPath   string
	Port         string
	KafkaBrokers []string
	JwtIssuer    string
}
