/*
   Copyright 2019 txn2
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
	"flag"
	"os"

	"github.com/txn2/provision"

	"github.com/gin-gonic/gin"
	"github.com/txn2/ack"
	"github.com/txn2/micro"
)

var (
	provisionPathEnv   = getEnv("PROVISION_PATH", "/provision")
	provisionHostEnv   = getEnv("PROVISION_HOST", "api-provision:8080")
	provisionSchemeEnv = getEnv("PROVISION_SCHEME", "http")
)

func main() {

	provisionPath := flag.String("provisionPath", provisionPathEnv, "Provision path")
	provisionHost := flag.String("provisionHost", provisionHostEnv, "Provision host")
	provisionScheme := flag.String("provisionScheme", provisionSchemeEnv, "Provision scheme")

	serverCfg, _ := micro.NewServerCfg("Op")
	server := micro.NewServer(serverCfg)

	// User token middleware
	server.Router.Use(provision.UserTokenHandler())

	checkSysopHandler := func(c *gin.Context) {
		ak := ack.Gin(c)

		userI, ok := c.Get("User")
		if !ok {
			ak.SetPayloadType("ErrorMessage")
			ak.SetPayload("missing user token")
			ak.GinErrorAbort(401, "E401", "UnauthorizedAccess")
			return
		}

		user := userI.(*provision.User)

		if user.Active && user.Sysop {
			return
		}

		ak.SetPayloadType("ErrorMessage")
		ak.SetPayload("insufficient privileges")
		ak.GinErrorAbort(401, "E401", "UnauthorizedAccess")
		return
	}

	prov := server.Router.Group(*provisionPath, checkSysopHandler)

	prov.Any("/*any",
		server.ReverseProxy(micro.PxyCfg{
			Scheme: provisionScheme,
			Host:   provisionHost,
			Strip:  provisionPath,
		}),
	)

	// run op server
	server.Run()
}

// getEnv gets an environment variable or sets a default if
// one does not exist.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}

	return value
}
