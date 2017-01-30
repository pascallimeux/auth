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

package common

import (
	"time"
)

const (
	ADMINLOGIN = "admin"
	ADMINPWD   = "orangeadmin"
	ADMINEMAIL = "admin@orange.com"
)

type Configuration struct {
	Logger                  string
	DataSourceName          string
	LogFileName             string
	Expire_in_token_in_hour time.Duration
	HttpHostUrl             string
	ReadTimeout             time.Duration
	WriteTimeout            time.Duration
	HandlerTimeout          time.Duration
	HLTimeout               time.Duration
	DeployTimeout           time.Duration
	TransactionTimeout      time.Duration
}
