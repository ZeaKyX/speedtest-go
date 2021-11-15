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
	pflag.StringP("statistics_password", "s", "", "")
	pflag.StringP("redact_ip_addresses", "r", "", "")
	pflag.StringP("database_type", "", "", "")
	pflag.StringP("database_hostname", "", "", "")
	pflag.StringP("database_name", "", "", "")
	pflag.StringP("database_username", "", "", "")

	viper.BindPFlags(pflag.CommandLine)
}

func main() {
	pflag.Parse()
	conf := config.Load(*optConfig)
	viper.Unmarshal(&conf)
	web.SetServerLocation(&conf)
	results.Initialize(&conf)
	database.SetDBInfo(&conf)
	log.Fatal(web.ListenAndServe(&conf))
}
