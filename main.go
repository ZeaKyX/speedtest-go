package main

import (
	_ "time/tzdata"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/librespeed/speedtest/config"
	"github.com/librespeed/speedtest/database"
	"github.com/librespeed/speedtest/results"
	"github.com/librespeed/speedtest/web"

	_ "github.com/breml/rootcerts"
	log "github.com/sirupsen/logrus"
)

var (
	optConfig = pflag.StringP("configfilepath", "c", "", "config file to be used, defaults to settings.toml in the same directory")
)

func init() {
	pflag.StringP("listen_port", "p", "", "backend listen port, default is 8989")
	pflag.StringP("proxyprotocol_port", "", "", "proxy protocol port, use 0 to disable")
	pflag.StringP("download_chunks", "d", "", "")
	pflag.StringP("distance_unit", "", "", "")
	pflag.StringP("enable_cors", "e", "", "")
	pflag.StringP("statistics_password", "s", "", "password for logging into statistics page, change this to enable stats page")
	pflag.StringP("redact_ip_addresses", "r", "", "redact IP addresses")
	pflag.StringP("database_type", "", "", "database type for statistics data, currently supports: none, memory, bolt, mysql, postgresql")
	pflag.StringP("database_hostname", "", "", "")
	pflag.StringP("database_name", "", "", "")
	pflag.StringP("database_username", "", "", "")

	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	pflag.Parse()
	conf := config.Load(*optConfig)
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Warnf("No config file found in search paths, using default values")
		} else {
			log.Fatalf("Error reading config: %s", err)
		}
	}

	if err := viper.Unmarshal(&conf); err != nil {
		log.Fatalf("Error parsing config: %s", err)
	}
	web.SetServerLocation(&conf)
	results.Initialize(&conf)
	database.SetDBInfo(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}
