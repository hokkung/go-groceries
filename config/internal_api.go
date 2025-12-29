package config

type InternalAPIConfiguration struct {
	AppProperties    AppProperties
	ServerProperties ServerProperties
}

// NewConfiguration returns new config instance
func NewInternalAPIConfiguration(
	app AppProperties,
) *InternalAPIConfiguration {
	return &InternalAPIConfiguration{
		AppProperties: app,
	}
}

// ProvideConfiguration provides configuration for di
func ProvideInternalAPIConfiguration(
	app AppProperties,
) InternalAPIConfiguration {
	return *NewInternalAPIConfiguration(app)
}
