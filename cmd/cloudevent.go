/*
Copyright 2018 Bitnami

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
	"github.com/pavangudiwada/kubewatch/config"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cloudEventConfigCmd = &cobra.Command{
	Use:   "cloudevent",
	Short: "specific cloudevent configuration",
	Long:  `specific cloudevent configuration`,
	Run: func(cmd *cobra.Command, args []string) {
		conf, err := config.New()
		if err != nil {
			logrus.Fatal(err)
		}

		url, err := cmd.Flags().GetString("url")
		if err == nil {
			if len(url) > 0 {
				conf.Handler.CloudEvent.Url = url
			}
		} else {
			logrus.Fatal(err)
		}

		if err = conf.Write(); err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	cloudEventConfigCmd.Flags().StringP("url", "u", "", "Specify CloudEvent url")
}
