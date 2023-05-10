package jaeger

type ENV string

const (
	Dev  ENV = "dev"
	Test ENV = "test"
	Stag ENV = "stag"
	Prod ENV = "prod"
)

var (
	defaultJaegerAddress = "127.0.0.1:6831"

	EnvMapHttp = map[ENV]string{
		Dev:  defaultJaegerAddress,
		Test: defaultJaegerAddress,
		Stag: defaultJaegerAddress,
		Prod: defaultJaegerAddress,
	}
)
