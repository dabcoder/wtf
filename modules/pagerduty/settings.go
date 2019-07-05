package pagerduty

import (
	"os"

	"github.com/olebedev/config"
	"github.com/wtfutil/wtf/cfg"
)

const defaultTitle = "PagerDuty"

type Settings struct {
	common *cfg.Common

	apiKey           string        `help:"Your PagerDuty API key."`
	escalationFilter []interface{} `help:"An array of schedule names you want to filter on."`
	scheduleIDs      []interface{} `help:"An array of schedule IDs you want to restrict the query to."`
	showIncidents    bool          `help:"Whether or not to list incidents." optional:"true"`
	showSchedules    bool          `help:"Whether or not to show schedules." optional:"true"`
}

func NewSettingsFromYAML(name string, ymlConfig *config.Config, globalConfig *config.Config) *Settings {

	settings := Settings{
		common: cfg.NewCommonSettingsFromModule(name, defaultTitle, ymlConfig, globalConfig),

		apiKey:           ymlConfig.UString("apiKey", os.Getenv("WTF_PAGERDUTY_API_KEY")),
		escalationFilter: ymlConfig.UList("escalationFilter"),
		scheduleIDs:      ymlConfig.UList("scheduleIDs", []interface{}{}),
		showIncidents:    ymlConfig.UBool("showIncidents", true),
		showSchedules:    ymlConfig.UBool("showSchedules", true),
	}

	return &settings
}
