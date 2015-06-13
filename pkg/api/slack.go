package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/grafana/grafana/pkg/api/dtos"
	"github.com/grafana/grafana/pkg/log"
	"github.com/grafana/grafana/pkg/middleware"
	"github.com/grafana/grafana/pkg/util"
)

func ShareWithSlack(c *middleware.Context, cmd dtos.ShareWithIntegrationCommand) Response {
	webhookUrl := "https://hooks.slack.com/services/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"

	message := util.DynMap{
		"username": "Grafana",
		"icon_url": "http://grafana.org/img/fav32.png",
		"channel":  "#grafana",
		"attachments": []util.DynMap{
			util.DynMap{
				"title":      "Cool Graph from Grafana",
				"title_link": cmd.ShareUrl,
				"color":      "#EF843C",
				"image_url":  cmd.ImageUrl,
			},
		},
	}

	messageJson, _ := json.Marshal(message)

	values := url.Values{
		"payload": []string{string(messageJson)},
	}

	client := http.Client{Timeout: time.Duration(5 * time.Second)}

	res, err := client.PostForm(webhookUrl, values)
	if err != nil {
		return ApiError(500, "Slack returned error", err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return ApiError(500, "Slack response failed to parse", err)
	}

	// var resData map[string]interface{}
	// if err := json.Unmarshal(body, &resData); err != nil {
	// 	return ApiError(500, "Failed to parse slack response to json", err)
	// }
	//
	log.Info("Slack response: %s", string(body))

	return ApiSuccess("Message sent to slack")
}
