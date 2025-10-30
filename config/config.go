package config

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	Server      ServerConfig
	Postgres    PostgresConfig
	Logger      LoggerConfig
	Casbin      CasbinConfig
	GDrive      GDriveConfig
	Redis       RedisConfig
	RateLimiter RateLimiterConfig
	RedisKey    RedisKey
	Minio       MinioConfig
	SMTP        SMTPConfig
	Nats        NatsConfig
}

type ServerConfig struct {
	Host         string
	Port         int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	Mode         string
	SSL          bool
	JWTSecretKey string
	EncryptKey   string
}

type PostgresConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Dbname   string
	SSLMode  string
}

type LoggerConfig struct {
	Level       string
	Caller      bool
	Encoding    string
	Development bool
}

type CasbinConfig struct {
	Model  string
	Policy string
}

type GDriveConfig struct {
	FOLDERIDDRIVE string
	CREDPATHDRIVE string
	LOCALHOST     string
}

type RedisConfig struct {
	REDISADDR string
	REDISDB   string
	REDISPW   string

	// Pooling
	PoolSize        int
	MinIdleConns    int
	PoolTimeout     time.Duration
	ConnMaxIdleTime time.Duration
	ConnMaxLifeTime time.Duration
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
}

// RateLimiterConfig holds all rate limiter configurations
type RateLimiterConfig struct {
	Enabled bool // Global enable/disable flag
	API     RateLimitConfig
	Auth    RateLimitConfig
	Device  RateLimitConfig
	Admin   RateLimitConfig
	Client  RateLimitConfig
}

// RateLimitConfig holds individual rate limit configuration
type RateLimitConfig struct {
	RequestsPerMinute int
	Burst             int
}

// RedisKey
type RedisKey struct {
	ProvinceKey string
}

// Minio
type MinioConfig struct {
	MinioEndpoint   string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
	BucketName      string
	DefaultAvatar   string
}

type SMTPConfig struct {
	Host     string
	Port     string
	Email    string
	Password string
}

type NatsConfig struct {
	URL      string
	Username string
	Password string
}

func NewAppConfig(configPath string) (*Config, error) {
	v := viper.New()

	v.AutomaticEnv()

	if _, err := os.Stat(".env"); err == nil {
		v.SetConfigFile(".env")
		v.SetConfigType("env")
		if err := v.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading .env file: %v", err)
		}
	}

	cfg := new(Config)

	// Server
	cfg.Server.Host = v.GetString("SERVER_HOST")
	cfg.Server.Port = v.GetInt("SERVER_PORT")
	cfg.Server.ReadTimeout = time.Duration(v.GetInt("SERVER_READ_TIMEOUT")) * time.Second
	cfg.Server.WriteTimeout = time.Duration(v.GetInt("SERVER_WRITE_TIMEOUT")) * time.Second
	cfg.Server.Mode = v.GetString("SERVER_MODE")
	cfg.Server.SSL = v.GetBool("SERVER_SSL")
	cfg.Server.JWTSecretKey = v.GetString("SERVER_JWT_SECRET_KEY")
	cfg.Server.EncryptKey = v.GetString("SERVER_ENCRYPT_KEY")

	// Postgres
	cfg.Postgres.User = v.GetString("POSTGRES_USER")
	cfg.Postgres.Password = v.GetString("POSTGRES_PASSWORD")
	cfg.Postgres.Host = v.GetString("POSTGRES_HOST")
	cfg.Postgres.Port = v.GetInt("POSTGRES_PORT")
	cfg.Postgres.Dbname = v.GetString("POSTGRES_DBNAME")
	cfg.Postgres.SSLMode = v.GetString("POSTGRES_SSL_MODE")

	// Logger
	cfg.Logger.Level = v.GetString("LOGGER_LEVEL")
	cfg.Logger.Caller = v.GetBool("LOGGER_CALLER")
	cfg.Logger.Encoding = v.GetString("LOGGER_ENCODING")
	cfg.Logger.Development = v.GetBool("LOGGER_DEVELOPMENT")

	// Casbin
	cfg.Casbin.Model = v.GetString("CASBIN_MODEL")
	cfg.Casbin.Policy = v.GetString("CASBIN_POLICY")

	// Google drive
	cfg.GDrive.FOLDERIDDRIVE = v.GetString("FOLDERIDDRIVE")
	cfg.GDrive.CREDPATHDRIVE = v.GetString("CREDPATHDRIVE")
	cfg.GDrive.LOCALHOST = v.GetString("LOCALHOST")

	// Redis
	cfg.Redis.REDISADDR = v.GetString("REDIS_ADDR")
	cfg.Redis.REDISDB = v.GetString("REDIS_DB")
	cfg.Redis.REDISPW = v.GetString("REDIS_PW")
	cfg.Redis.PoolSize = v.GetInt("REDIS_POOL_SIZE")
	cfg.Redis.MinIdleConns = v.GetInt("REDIS_MIN_IDLE_CONNS")
	cfg.Redis.PoolTimeout = v.GetDuration("REDIS_POOL_TIMEOUT")
	cfg.Redis.ConnMaxIdleTime = v.GetDuration("REDIS_CONN_MAX_IDLE_TIME")
	cfg.Redis.ConnMaxLifeTime = v.GetDuration("REDIS_CONN_MAX_LIFE_TIME")
	cfg.Redis.ReadTimeout = v.GetDuration("REDIS_READ_TIMEOUT")
	cfg.Redis.WriteTimeout = v.GetDuration("REDIS_WRITE_TIMEOUT")

	cfg.RedisKey.ProvinceKey = v.GetString("PROVINCE_KEY_REDIS")

	// Rate Limiter
	cfg.RateLimiter.Enabled = getBoolWithDefault(v, "RATE_LIMITER_ENABLED", true)

	cfg.RateLimiter.API.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_API_REQUESTS_PER_MINUTE", 100)
	cfg.RateLimiter.API.Burst = getIntWithDefault(v, "RATE_LIMIT_API_BURST", 20)

	cfg.RateLimiter.Auth.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_AUTH_REQUESTS_PER_MINUTE", 10)
	cfg.RateLimiter.Auth.Burst = getIntWithDefault(v, "RATE_LIMIT_AUTH_BURST", 3)

	cfg.RateLimiter.Device.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_DEVICE_REQUESTS_PER_MINUTE", 50)
	cfg.RateLimiter.Device.Burst = getIntWithDefault(v, "RATE_LIMIT_DEVICE_BURST", 10)

	cfg.RateLimiter.Admin.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_ADMIN_REQUESTS_PER_MINUTE", 200)
	cfg.RateLimiter.Admin.Burst = getIntWithDefault(v, "RATE_LIMIT_ADMIN_BURST", 30)

	cfg.RateLimiter.Client.RequestsPerMinute = getIntWithDefault(v, "RATE_LIMIT_CLIENT_REQUESTS_PER_MINUTE", 100)
	cfg.RateLimiter.Client.Burst = getIntWithDefault(v, "RATE_LIMIT_CLIENT_BURST", 10)

	cfg.Minio.MinioEndpoint = v.GetString("MINIO_ENDPOINT")
	cfg.Minio.AccessKeyID = v.GetString("MINIO_ACCESS_KEY_ID")
	cfg.Minio.SecretAccessKey = v.GetString("MINIO_SECRET_ACCESS_KEY")
	cfg.Minio.UseSSL = v.GetBool("MINIO_USE_SSL")
	cfg.Minio.BucketName = v.GetString("MINIO_BUCKET_NAME")
	cfg.Minio.DefaultAvatar = v.GetString("MINIO_DEFAULT_AVATAR")

	cfg.SMTP.Host = v.GetString("SMTP_HOST")
	cfg.SMTP.Port = v.GetString("SMTP_PORT")
	cfg.SMTP.Email = v.GetString("SMTP_EMAIL")
	cfg.SMTP.Password = v.GetString("SMTP_PASSWORD")

	// NATS
	cfg.Nats.URL = v.GetString("NATS_URL")
	cfg.Nats.Username = v.GetString("NATS_USERNAME")
	cfg.Nats.Password = v.GetString("NATS_PASSWORD")

	return cfg, nil
}

func getIntWithDefault(v *viper.Viper, key string, defaultValue int) int {
	if v.IsSet(key) {
		return v.GetInt(key)
	}
	return defaultValue
}

func getBoolWithDefault(v *viper.Viper, key string, defaultValue bool) bool {
	if v.IsSet(key) {
		return v.GetBool(key)
	}
	return defaultValue
}

func getStringWithDefault(v *viper.Viper, key string, defaultValue string) string {
	if v.IsSet(key) {
		return v.GetString(key)
	}
	return defaultValue
}
