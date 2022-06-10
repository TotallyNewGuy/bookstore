package util

import "github.com/spf13/viper"

type Config struct {
	BookGrpcAddress     string `mapstructure:"BOOK_GRPC_ADDRESS"`
	UserGrpcAddress     string `mapstructure:"USER_GRPC_ADDRESS"`
	OrderGrpcAddress    string `mapstructure:"ORDER_GRPC_ADDRESS"`
	PurchaseHttpAddress string `mapstructure:"PURCHASE_HTTP_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
