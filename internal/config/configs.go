package config

// type env string
//
// const (
// 	EnvDevelop    env = "develop"
// 	EnvTest       env = "test"
// 	EnvStaging    env = "staging"
// 	EnvProduction env = "production"
// )



type Config struct {
  Environment string
  AppName string
	Port int
  AppVersion string
}

