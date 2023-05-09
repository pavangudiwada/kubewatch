/*
Copyright 2016 Skippbox, Ltd.

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

package client

import (
	"github.com/pavangudiwada/kubewatchconfig"
	"github.com/pavangudiwada/kubewatchpkg/controller"
	"github.com/pavangudiwada/kubewatchpkg/handlers"
	"github.com/pavangudiwada/kubewatchpkg/handlers/cloudevent"
	"github.com/pavangudiwada/kubewatchpkg/handlers/flock"
	"github.com/pavangudiwada/kubewatchpkg/handlers/hipchat"
	"github.com/pavangudiwada/kubewatchpkg/handlers/lark"
	"github.com/pavangudiwada/kubewatchpkg/handlers/mattermost"
	"github.com/pavangudiwada/kubewatchpkg/handlers/msteam"
	"github.com/pavangudiwada/kubewatchpkg/handlers/slack"
	"github.com/pavangudiwada/kubewatchpkg/handlers/slackwebhook"
	"github.com/pavangudiwada/kubewatchpkg/handlers/smtp"
	"github.com/pavangudiwada/kubewatchpkg/handlers/webhook"
	"github.com/sirupsen/logrus"
)

// Run runs the event loop processing with given handler
func Run(conf *config.Config) {

	var eventHandler = ParseEventHandler(conf)
	controller.Start(conf, eventHandler)
}

// ParseEventHandler returns the respective handler object specified in the config file.
func ParseEventHandler(conf *config.Config) handlers.Handler {

	var eventHandler handlers.Handler
	switch {
	case len(conf.Handler.Slack.Channel) > 0 || len(conf.Handler.Slack.Token) > 0:
		eventHandler = new(slack.Slack)
	case len(conf.Handler.SlackWebhook.Channel) > 0 || len(conf.Handler.SlackWebhook.Username) > 0 || len(conf.Handler.SlackWebhook.Slackwebhookurl) > 0:
		eventHandler = new(slackwebhook.SlackWebhook)
	case len(conf.Handler.Hipchat.Room) > 0 || len(conf.Handler.Hipchat.Token) > 0:
		eventHandler = new(hipchat.Hipchat)
	case len(conf.Handler.Mattermost.Channel) > 0 || len(conf.Handler.Mattermost.Url) > 0:
		eventHandler = new(mattermost.Mattermost)
	case len(conf.Handler.Flock.Url) > 0:
		eventHandler = new(flock.Flock)
	case len(conf.Handler.Webhook.Url) > 0:
		eventHandler = new(webhook.Webhook)
	case len(conf.Handler.CloudEvent.Url) > 0:
		eventHandler = new(cloudevent.CloudEvent)
	case len(conf.Handler.MSTeams.WebhookURL) > 0:
		eventHandler = new(msteam.MSTeams)
	case len(conf.Handler.SMTP.Smarthost) > 0 || len(conf.Handler.SMTP.To) > 0:
		eventHandler = new(smtp.SMTP)
	case len(conf.Handler.Lark.WebhookURL) > 0:
		eventHandler = new(lark.Webhook)
	default:
		eventHandler = new(handlers.Default)
	}
	if err := eventHandler.Init(conf); err != nil {
		logrus.Fatal(err)
	}
	return eventHandler
}
