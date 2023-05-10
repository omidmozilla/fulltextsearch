package config

type serverConfig struct {
	Host  string
	Port  string
}

type rpcClientConfig struct {
	Host  string
	Port  string
}

type dbConfig struct {
	Name  string
}

func GetDBConfig() dbConfig {
	return dbConfig {
		Name: "fulltextsearch",
	}
}

func GetServerConfig() serverConfig {
	return serverConfig {
		Host: "localhost",
		Port: "8080",
	}
}

func GetRPCConfig() rpcClientConfig {
	return rpcClientConfig {
		Host: "http://localhost",
		Port: "8088",
	}
}

func GetServerURL() string {
	serverConfig := GetServerConfig()
	return serverConfig.Host + ":" + serverConfig.Port
}

func GetRPCURL() string {
	rpcConfig := GetRPCConfig()
	return rpcConfig.Host + ":" + rpcConfig.Port
}