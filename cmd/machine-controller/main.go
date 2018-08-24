/*
Copyright 2018 The Kubernetes Authors.

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

	"github.com/golang/glog"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"k8s.io/apiserver/pkg/util/logs"
	"sigs.k8s.io/cluster-api/pkg/controller/config"

	"sigs.k8s.io/cluster-api-provider-aws/cloud/aws/controllers/machine"
	"sigs.k8s.io/cluster-api-provider-aws/cloud/aws/controllers/machine/options"
)

var (
	logLevel string
)

func init() {
	config.ControllerConfig.AddFlags(pflag.CommandLine)
	pflag.CommandLine.StringVar(&logLevel, "log-level", defaultLogLevel, "Log level (debug,info,warn,error,fatal)")
}

const (
	controllerLogName = "awsMachine"
	defaultLogLevel   = "info"
)

func main() {
	// the following line exists to make glog happy, for more information, see: https://github.com/kubernetes/kubernetes/issues/17162
	flag.CommandLine.Parse([]string{})
	pflag.Parse()

	logs.InitLogs()
	defer logs.FlushLogs()

	log.SetOutput(os.Stdout)
	if lvl, err := log.ParseLevel(logLevel); err != nil {
		log.Panic(err)
	} else {
		log.SetLevel(lvl)
	}

	machineServer := options.NewServer()
	if err := machine.Run(machineServer); err != nil {
		glog.Errorf("Failed to start cluster controller. Err: %v", err)
	}
}
