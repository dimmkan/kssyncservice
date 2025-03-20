package config

import (
  "os"
  "strconv"
  "strings"
)

type DbConfig struct {
  DB_HOST string;
  DB_PORT string;
  DB_USER string;
  DB_PASSWORD string;
  DB_DATABASE string;
}

type ExternalServiceConfig struct {
  KS_SERVICE_HOST string;
  KS_SERVICE_PORT string;
  KS_SERVICE_METHOD string;
}

type CronJobsConfig struct {
  KS_SYNC_SERVICES_CRON string;
}

type Config struct {
  DbConfig DbConfig;
  ExternalServiceConfig ExternalServiceConfig;
  CronJobsConfig CronJobsConfig;
  ExternalFileConfig ExternalFileConfig;
}

type ExternalFileConfig struct {
  FILE_SERVICE_HOST string;
  FILE_SERVICE_PORT string;
  FILE_SERVICE_METHOD string;
}

func New() *Config {
  return &Config {
    DbConfig: DbConfig{
      DB_HOST: getEnv("DB_HOST", ""),
      DB_PORT: getEnv("DB_PORT", ""),
      DB_USER: getEnv("DB_USER", ""),
      DB_PASSWORD: getEnv("DB_PASSWORD", ""),
      DB_DATABASE: getEnv("DB_DATABASE", ""),
    },
    ExternalServiceConfig: ExternalServiceConfig {
      KS_SERVICE_HOST: getEnv("KS_SERVICE_HOST", ""),
      KS_SERVICE_PORT: getEnv("KS_SERVICE_PORT", ""),
      KS_SERVICE_METHOD: getEnv("KS_SERVICE_METHOD", ""),
    },
    CronJobsConfig: CronJobsConfig {
      KS_SYNC_SERVICES_CRON: getEnv("KS_SYNC_SERVICES_CRON", ""),
    },
    ExternalFileConfig: ExternalFileConfig{
      FILE_SERVICE_HOST: getEnv("FILE_SERVICE_HOST", ""),
      FILE_SERVICE_PORT: getEnv("FILE_SERVICE_PORT", ""),
      FILE_SERVICE_METHOD: getEnv("FILE_SERVICE_METHOD", ""),
    },
  };
}

func getEnv(key string, defaultVal string) string {
  if value, exists := os.LookupEnv(key); exists {
return value
  }

  return defaultVal
}

func getEnvAsInt(name string, defaultVal int) int {
  valueStr := getEnv(name, "")
  if value, err := strconv.Atoi(valueStr); err == nil {
return value
  }

  return defaultVal
}

func getEnvAsBool(name string, defaultVal bool) bool {
  valStr := getEnv(name, "")
  if val, err := strconv.ParseBool(valStr); err == nil {
return val
  }

  return defaultVal
}

func getEnvAsSlice(name string, defaultVal []string, sep string) []string {
  valStr := getEnv(name, "")

  if valStr == "" {
return defaultVal
  }

  val := strings.Split(valStr, sep)

  return val
}