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
	"encoding/json"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/packages/task_queue"
	"github.com/seanlan/wool/core"
	"github.com/seanlan/wool/utils"
	"github.com/spf13/cobra"
)

const RoutingKey = ""

func testTaskQueue() {
	_, _ = task_queue.SendTask("PrintOK", []tasks.Arg{}, RoutingKey, 0)
	_, _ = task_queue.SendTask("SayHi", []tasks.Arg{
		{
			Name:  "arg",
			Type:  "string",
			Value: "World",
		},
	}, RoutingKey, 0)
	_, _ = task_queue.SendTask("Add", []tasks.Arg{
		{
			Type:  "int64",
			Value: 40,
		},
		{
			Type:  "int64",
			Value: 20,
		},
	},
		RoutingKey,
		0)
}

func testGorm() {
}

func makeToken(args string) {
	j := utils.NewJWT("u98ef98sudfae98")
	token, _ := j.CreateToken(args, 3600*24*30)
	logging.Logger.Info(token)
	msg, _ := json.Marshal(core.WSMessage{
		From:    "123123",
		To:      "123123",
		Event:   0,
		Content: "asdfasdf",
	})
	logging.Logger.Info(string(msg))
}

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test script",
	Long:  `test script`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		script := args[0]
		switch script {
		case "queue":
			testTaskQueue()
		case "gorm":
			testGorm()
		case "token":
			makeToken(args[1])
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
