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

package controllers

import (
	"github.com/gorilla/mux"
	"github.com/pascallimeux/auth/src/model"
	"github.com/pascallimeux/auth/src/utils"
	"github.com/pascallimeux/auth/src/utils/log"
	"net/http"
	"time"
)

//HTTP Get - /o/logs
func (a *AppContext) getLogs(w http.ResponseWriter, r *http.Request) {
	log.Trace(log.Here(), "getLogs() : calling method -")

	err0 := a.checkPermissionFromToken(w, r, "getLogs", "")
	if err0 != nil {
		return
	}

	logs, err := a.SqlContext.GetLogs()
	if err != nil {
		sendError(log.Here(), w, err)
		return
	}
	buildHttp200Response(w, logs)
}

//HTTP Get - /o/logs/{from}/{to}
func (a *AppContext) getLogs4dates(w http.ResponseWriter, r *http.Request) {
	log.Trace(log.Here(), "getLogs() : calling method -")

	err0 := a.checkPermissionFromToken(w, r, "getLogs", "")
	if err0 != nil {
		return
	}

	vars := mux.Vars(r)
	var err error
	var date1, date2 time.Time
	var logs []model.Logg

	date1, err = utils.DateParse(vars["from"])
	if err != nil {
		sendError(log.Here(), w, err)
		return
	}
	date2, err = utils.DateParse(vars["to"])
	if err != nil {
		sendError(log.Here(), w, err)
		return
	}
	logs, err = a.SqlContext.GetLog4Period(date1, date2)
	if err != nil {
		sendError(log.Here(), w, err)
		return
	}
	buildHttp200Response(w, logs)
}
