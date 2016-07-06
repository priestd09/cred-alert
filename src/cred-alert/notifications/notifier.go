package notifications

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pivotal-golang/lager"
)

//go:generate counterfeiter . Notifier

type Notifier interface {
	SendNotification(message string) error
}

type slackNotifier struct {
	logger lager.Logger

	webhookURL string
	client     *http.Client
}

type slackNotification struct {
	Text string `json:"text"`
}

func NewSlackNotifier(logger lager.Logger, webhookURL string) Notifier {
	if webhookURL == "" {
		return &nullSlackNotifier{
			logger: logger,
		}
	}

	return &slackNotifier{
		webhookURL: webhookURL,
		client:     &http.Client{},
		logger:     logger,
	}
}

func (n *slackNotifier) SendNotification(message string) error {
	logger := n.logger.Session("send-notification", lager.Data{
		"message": message,
	})

	notification := &slackNotification{Text: message}
	body, err := json.Marshal(notification)
	if err != nil {
		logger.Error("unmarshal-faiiled", err)
		return err
	}

	req, err := http.NewRequest("POST", n.webhookURL, bytes.NewBuffer(body))
	if err != nil {
		logger.Error("request-failed", err)
		return err
	}

	resp, err := n.client.Do(req)
	if err != nil {
		logger.Error("response-error", err)
		return err
	}

	if resp.StatusCode != http.StatusOK {
		logger.Info(fmt.Sprintf("bad responze (!200): %d", resp.StatusCode))
		return fmt.Errorf("bad responze (!200): %d", resp.StatusCode)
	}

	return nil
}

type nullSlackNotifier struct {
	logger lager.Logger
}

func (n *nullSlackNotifier) SendNotification(message string) error {
	n.logger.Session("send-notification", lager.Data{
		"message": message,
	}).Debug("sent")
	return nil
}
