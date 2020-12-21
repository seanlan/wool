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
package main

import (
	machinery_conf "github.com/RichardKnop/machinery/v1/config"
	"github.com/seanlan/packages/config"
	"github.com/seanlan/packages/db"
	"github.com/seanlan/packages/gredis"
	"github.com/seanlan/packages/logging"
	"github.com/seanlan/packages/task_queue"
	"github.com/seanlan/wool/cmd"
)

func setEnv() {
	//初始化设置
	config.Setup("./conf.d/conf.yaml")
	//日志初始化
	logging.Setup(
		config.GetBoolean("LOG_DEBUG"),
		config.GetString("APP_NAME"))
	//redis配置初始化
	_ = gredis.Setup(config.GetString("REDIS"))
	//任务队列初始化
	var celeryServerConfig machinery_conf.Config
	_ = config.GetValue("TASK_QUEUE").Populate(&celeryServerConfig)
	task_queue.Setup(celeryServerConfig)
	db.Setup(config.GetString("MYSQL"))
}

func main() {
	setEnv()
	cmd.Execute()
}
