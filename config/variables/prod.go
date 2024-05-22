package variables

var Prod = Config{
	PrefixPath:   "/pinares",
	Port:         "8082",
	KafkaBrokers: []string{"localhost:9092"},
}
