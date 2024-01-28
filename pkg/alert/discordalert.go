package alert

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	"lavaprovider-monitor/pkg/config"
	"lavaprovider-monitor/pkg/log"
	"net/http"
	"time"
)

type webhookPayload struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds"`
}

// SendMsgViaWebhook  send msg via discord websocket
func SendMsgViaWebhook(webhookUrl string, webhookPayload webhookPayload) error {
	b, err := json.Marshal(webhookPayload)
	if err != nil {
		return fmt.Errorf("failed to alert due to fail to marshal webhook payload: %v", err)
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", webhookUrl, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("failed to alert due to fail to create http request: %v", err)
	}

	req.Header.Add("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to alert due to fail to send http request: %v", err)
	}
	if res.StatusCode != 204 {
		return fmt.Errorf("failed to alert,: %v", res)
	}
	return nil
}

// getEmbedAlert get Embed alert
func getEmbedAlert(alerts []Alert) []Embed {
	var embeds []Embed
	for _, alert := range alerts {
		timestamp := fmt.Sprint(alert.Timestamp.Unix())
		embed := Embed{
			Title:       alert.Title,
			Description: alert.Message,
			Color:       16711680,
			Fields: []*EmbedField{
				{
					Name:   "Type",
					Value:  alert.Type,
					Inline: false,
				},
				{
					Name:   "Origin",
					Value:  alert.Origin,
					Inline: false,
				},
				{
					Name:   "Details",
					Value:  alert.Details,
					Inline: true,
				},
			},
			Thumbnail: &EmbedThumbnail{
				URL: "https://i.imgur.com/5NmtLLy.jpg",
			},
		}
		if alert.AdditionalData != nil {
			for key, value := range alert.AdditionalData {
				field := &EmbedField{
					Name:   key,
					Value:  value,
					Inline: false,
				}
				embed.Fields = append(embed.Fields, field)
			}
		}
		time := &EmbedField{
			Name:   "Time",
			Value:  "<t:" + timestamp + ">",
			Inline: false,
		}
		embed.Fields = append(embed.Fields, time)
		embeds = append(embeds, embed)
	}
	return embeds
}

// SendAlertViaDiscord send alert
func SendAlertViaDiscord(config config.DiscordAlertConfig, alerts []Alert) error {
	embeds := getEmbedAlert(alerts)
	var mention string
	if len(config.Alertuserid) != 0 && config.Alertuserid[0] != "" {
		for _, userid := range config.Alertuserid {
			mention += "<@" + userid + "> "
		}
	}
	if len(config.Alertroleid) != 0 && config.Alertroleid[0] != "" {
		for _, roleid := range config.Alertroleid {
			mention += "<@&" + roleid + "> "
		}
	}

	payload := webhookPayload{
		Content: "" + mention,
		Embeds:  embeds,
	}
	var err error
	retry.Do(
		func() error {
			err = SendMsgViaWebhook(config.Webhook, payload)
			if err != nil {
				return err
			}
			return nil
		},
		retry.Delay(time.Second*3),
		retry.Attempts(3),
	)
	if err != nil {
		return fmt.Errorf("failed to alert via discord: %v", err)
	}
	log.Log.Info("alert via discord successfully")
	return nil
}
