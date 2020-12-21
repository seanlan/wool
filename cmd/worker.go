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
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/packages/task_queue"
	tasks2 "github.com/seanlan/wool/tasks"

	"github.com/spf13/cobra"
)

func registerTasks() {
	tasks := map[string]interface{}{
		"PrintOK":     task_queue.PrintOK,
		"SayHi":       task_queue.SayHi,
		"Add":         task_queue.Add,
		"SaveMessage": tasks2.SaveChatMessage,
	}
	task_queue.RegisterTasks(tasks)
}

var consumerTag string
var concurrency int
var queueName string

// workerCmd represents the worker command
var workerCmd = &cobra.Command{
	Use:   "worker",
	Short: "start task queue worker",
	Long:  `start task queue worker`,
	Run: func(cmd *cobra.Command, args []string) {
		registerTasks()
		logging.Logger.Infof("consumerTag : %v", consumerTag)
		logging.Logger.Infof("concurrency : %v", concurrency)
		task_queue.RunWorker(consumerTag, concurrency, queueName)
	},
}

func init() {
	workerCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 5, "concurrency number")
	workerCmd.Flags().StringVarP(&consumerTag, "consumer", "s", "", "consumer tag")
	workerCmd.Flags().StringVarP(&queueName, "queue", "q", "", "queue name")
	rootCmd.AddCommand(workerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
