package cache

// Config holds cache configuration.
type Config struct {
	Driver      string
	RedisConfig RedisConfig
}
