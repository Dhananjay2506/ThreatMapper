package directory

import (
	"os"
	"strconv"

	"github.com/deepfence/ThreatMapper/deepfence_utils/log"
)

const (
	GlobalDirKey  = NamespaceID("global")
	NonSaaSDirKey = NamespaceID("default")
	NamespaceKey  = "namespace"
)

type NamespaceID string

type RedisConfig struct {
	Endpoint string
	Password string
	Database int
}

type Neo4jConfig struct {
	Endpoint string
	Username string
	Password string
}

type PostgresqlConfig struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	SslMode  string
}

type DBConfigs struct {
	Redis    RedisConfig
	Neo4j    *Neo4jConfig
	Postgres *PostgresqlConfig
}

var directory map[NamespaceID]DBConfigs

func init() {

	redisCfg := initRedis()

	postgresqlCfg := init_posgres()

	saasMode := false
	saasModeOn, has := os.LookupEnv("DEEPFENCE_SAAS_MODE")
	if !has {
		log.Warn().Msg("DEEPFENCE_SAAS_MODE defaults to: off")
	} else if saasModeOn == "on" {
		saasMode = true
	}

	directory = map[NamespaceID]DBConfigs{}

	if !saasMode {
		neo4jCfg := init_neo4j()
		directory[NonSaaSDirKey] = DBConfigs{
			Redis:    redisCfg,
			Neo4j:    &neo4jCfg,
			Postgres: nil,
		}
	}

	directory[GlobalDirKey] = DBConfigs{
		Redis:    redisCfg,
		Neo4j:    nil,
		Postgres: &postgresqlCfg,
	}
}

func initRedis() RedisConfig {
	redisHost, has := os.LookupEnv("DEEPFENCE_REDIS_HOST")
	if !has {
		redisHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_REDIS_HOST defaults to: %v", redisHost)
	}
	redisPort, has := os.LookupEnv("DEEPFENCE_REDIS_PORT")
	if !has {
		redisPort = "6379"
		log.Warn().Msgf("DEEPFENCE_REDIS_PORT defaults to: %v", redisPort)
	}
	redisEndpoint := redisHost + ":" + redisPort
	redisPassword := os.Getenv("DEEPFENCE_REDIS_PASSWORD")
	redisDbNumber := 0
	var err error
	redisDbNumberStr := os.Getenv("DEEPFENCE_REDIS_DB_NUMBER")
	if redisDbNumberStr != "" {
		redisDbNumber, err = strconv.Atoi(redisDbNumberStr)
		if err != nil {
			redisDbNumber = 0
		}
	}
	return RedisConfig{
		Endpoint: redisEndpoint,
		Password: redisPassword,
		Database: redisDbNumber,
	}
}

func init_posgres() PostgresqlConfig {
	var err error
	postgresHost, has := os.LookupEnv("DEEPFENCE_POSTGRES_USER_DB_HOST")
	if !has {
		postgresHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_POSTGRES_USER_DB_HOST defaults to: %v", postgresHost)
	}
	postgresPort := 5432
	postgresPortStr := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_PORT")
	if postgresPortStr == "" {
		log.Warn().Msgf("DEEPFENCE_POSTGRES_USER_DB_PORT defaults to: %d", postgresPort)
	} else {
		postgresPort, err = strconv.Atoi(postgresPortStr)
		if err != nil {
			postgresPort = 5432
		}
	}
	postgresUsername := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_USER")
	postgresPassword := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_PASSWORD")
	postgresDatabase := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_NAME")
	postgresSslMode := os.Getenv("DEEPFENCE_POSTGRES_USER_DB_SSLMODE")

	return PostgresqlConfig{
		Host:     postgresHost,
		Port:     postgresPort,
		Username: postgresUsername,
		Password: postgresPassword,
		Database: postgresDatabase,
		SslMode:  postgresSslMode,
	}
}

func init_neo4j() Neo4jConfig {
	neo4jHost, has := os.LookupEnv("DEEPFENCE_NEO4J_HOST")
	if !has {
		neo4jHost = "localhost"
		log.Warn().Msgf("DEEPFENCE_NEO4J_HOST defaults to: %v", neo4jHost)
	}
	neo4jBoltPort, has := os.LookupEnv("DEEPFENCE_NEO4J_BOLT_PORT")
	if !has {
		neo4jBoltPort = "7687"
		log.Warn().Msgf("DEEPFENCE_NEO4J_BOLT_PORT defaults to: %v", neo4jBoltPort)
	}
	neo4jEndpoint := "bolt://" + neo4jHost + ":" + neo4jBoltPort
	neo4jUsername := os.Getenv("DEEPFENCE_NEO4J_USER")
	neo4jPassword := os.Getenv("DEEPFENCE_NEO4J_PASSWORD")
	return Neo4jConfig{
		Endpoint: neo4jEndpoint,
		Username: neo4jUsername,
		Password: neo4jPassword,
	}
}
