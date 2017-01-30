/*
Copyright Pascal Limeux. 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"github.com/pascallimeux/auth-go/src/common"
	"github.com/pascallimeux/auth-go/src/controllers"
	"github.com/pascallimeux/auth-go/src/model"
	"github.com/pascallimeux/auth-go/src/utils"
	"github.com/pascallimeux/auth-go/src/utils/log"
	"net/http"
	"os"
	"time"
)

func main() {

	// Init configuration
	config_file := "config.json"
	var configuration common.Configuration
	err := utils.Read_Conf(config_file, &configuration)
	if err != nil {
		panic(err.Error())
	}

	// Init logger
	f := log.Init_log(configuration.LogFileName, configuration.Logger)
	defer f.Close()

	// Get local IP address if possible
	ipAddress, err := utils.GetOutboundIP()
	if err != nil {
		log.Error(log.Here(), " Impossible to retrieve the IP address of this machine")
	} else {
		configuration.HttpHostUrl = ipAddress + ":8088"
	}

	// Write configuration in log
	log.Info(log.Here(), utils.Get_fields(configuration))

	// Init application context
	appContext := controllers.AppContext{Configuration: configuration}

	// Init sqliteDB
	sqlContext, err := model.GetSqlContext(configuration.DataSourceName)
	if err != nil {
		log.Fatal(log.Here(), err.Error())
	}
	defer sqlContext.Db.Close()
	appContext.SqlContext = sqlContext

	args := os.Args[1:]
	if len(args) == 1 {
		if args[0] == "init" {
			sqlContext.InitBDD()
		}
	}

	// Start http server
	router := appContext.CreateRoutes()
	log.Info(log.Here(), "Listening on: ", configuration.HttpHostUrl)
	s := &http.Server{
		Addr:         configuration.HttpHostUrl,
		Handler:      router,
		ReadTimeout:  configuration.ReadTimeout * time.Nanosecond,
		WriteTimeout: configuration.WriteTimeout * time.Nanosecond,
	}
	log.Fatal(log.Here(), s.ListenAndServe().Error())
}
