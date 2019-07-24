package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"go_demo/config"
	"go_demo/model"
	"go_demo/router"
	"go_demo/router/middleware"
	v "go_demo/package/version"
)

var (
	cfg = pflag.StringP(
		"config",
		"c",
		"conf/dev.yaml",
		"go_demo config file path",
	)
	version = pflag.BoolP(
		"version",
		"v",
		false,
		"show version info.",
	)
)

func main() {
	pflag.Parse()

	if *version {
		v := v.Get()
		marshalIed, err := json.MarshalIndent(&v, "", " ")
		if err != nil {
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}

		fmt.Println(string(marshalIed))
		return
	}

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	model.DB.Init()
	defer model.DB.Close()

	//set gin mode
	gin.SetMode(viper.GetString("runmode"))

	ctx := gin.New()

	router.Load(
		ctx,
		middleware.RequestId(),
		middleware.Logging(),
	)

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	go func() {
		reConnectDB()
	}()

	log.Infof("Start to listening the incoming requests on http address: %s", viper.GetString("addr"))
	log.Info(http.ListenAndServe(viper.GetString("addr"), ctx).Error())
}

func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/check/health")

		if err == nil && resp.StatusCode == 200 {
            return nil
		}

		fmt.Println("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second * 1)
	}

	return errors.New("Cannot connect to the router.")
}

func reConnectDB() {
	for {
		time.Sleep(time.Second * time.Duration(viper.GetInt("db_reconnect")))
		model.DB.Init()
		log.Info("The DB reconnect successfully.")
	}
}
