/*
 * Copyright 2022 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Package main provides the protoc-gen-go-hz plugin entry point.
// This is a protoc plugin that generates Hertz HTTP framework code from Protocol Buffer definitions.
//
// Usage:
//
//	protoc --go_out=. --go-hz_out=. your.proto
//
// The plugin can be configured with parameters:
//
//	protoc --go_out=. --go-hz_out=. --go-hz_opt=verbose=true,out_dir=. your.proto
package main

import (
	"os"

	"github.com/ca-x/protoc-gen-go-hz/pkg/plugin"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/compiler/protogen"
)

// main is the entry point for the protoc-gen-go-hz plugin.
// It uses the protogen framework to handle protobuf compilation pipeline integration.
func main() {
	// 使用 protogen 框架的标准插件入口点
	// protogen.Options{}.Run() 会自动处理：
	// 1. 从 stdin 读取 CodeGeneratorRequest
	// 2. 解析 proto 文件
	// 3. 调用插件的处理函数
	// 4. 将 CodeGeneratorResponse 写入 stdout
	protogen.Options{}.Run(func(gen *protogen.Plugin) error {
		// 创建 HZ 插件实例，封装 protogen.Plugin
		hzPlugin := plugin.NewHZPlugin(gen)

		// 设置日志输出到 stderr（protoc 要求插件使用 stderr 记录日志）
		logrus.SetOutput(os.Stderr)
		logrus.SetLevel(logrus.InfoLevel)

		// 运行插件的主要逻辑
		return hzPlugin.Run()
	})
}
