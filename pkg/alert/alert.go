package alert

import (
	"fmt"
	"lavaprovider-monitor/pkg/config"
	"lavaprovider-monitor/pkg/log"
	"time"
)

type Alert struct {
	Type           string
	Title          string
	Message        string
	Details        string
	Timestamp      time.Time
	Origin         string
	AdditionalData map[string]string
}

func CheckAlertConfig(config config.AlertConfig) {
	if config.Alert.Enable {
		if config.Discord.Enable {
			if config.Discord.Webhook == "" {
				panic("Discord alert is enabled , but Discord webhook is empty")
			}
		}

	}
}

func SendAlert(aconfig config.AlertConfig, alerts []Alert) error {
	if alerts == nil {
		log.Log.Info("No alerts to send")
		return nil
	}
	var errDiscord error
	if aconfig.Alert.Enable {
		if aconfig.Discord.Enable {
			log.Log.Info("start to send alert via discord")
			errDiscord = SendAlertViaDiscord(aconfig.Discord, alerts)
		}
	} else {
		log.Log.Info("Alert is disabled")
	}

	if errDiscord != nil {
		return fmt.Errorf("discord error: %v", errDiscord)
	}
	return nil
}
