/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/router"
	"github.com/seanlan/wool/web/routers"

	"github.com/spf13/cobra"
)

// webCmd represents the web command
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "start web api server",
	Long:  `start web api server`,
	Run: func(cmd *cobra.Command, args []string) {
		router.Setup(config.GetString("WEB", "GIN_MODE"))
		routerGroup := router.Router.Group("")
		routers.InitWebSocketRouter(routerGroup)
		imApiGroup := router.Router.Group("api/v1")
		routers.InitIMApiRouter(imApiGroup)
		router.Run(config.GetString("WEB", "HOST_NAME"))
	},
}

func init() {
	rootCmd.AddCommand(webCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// webCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// webCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
