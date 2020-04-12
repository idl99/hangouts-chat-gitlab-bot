package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/idl99/hangouts-chat-gitlab-bot/config"
)

// GitlabHandler - handler for Gitlab Webhook events
func GitlabHandler(w http.ResponseWriter, r *http.Request) {

	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()

	// NOTE: Attemping to read the request body more than once will result in EOF
	rawData, err := ioutil.ReadAll(r.Body)

	log.Println("Received GitlabEvent from Webhook")

	gitlabEvent := &GitlabEvent{}
	err = json.NewDecoder(bytes.NewBuffer(rawData)).Decode(gitlabEvent)
	log.Println("Gitlab event payload: ", gitlabEvent)

	if err != nil {
		panic("Failed to decode json payload received from webhook.")
	}

	if gitlabEvent.ObjectKind == "merge_request" {
		handleMergeRequest(bytes.NewBuffer(rawData))
	}
}

func handleMergeRequest(b *bytes.Buffer) {

	gitlabMergeRequestEvent := &GitlabMergeRequestEvent{}
	err := json.NewDecoder(b).Decode(gitlabMergeRequestEvent)

	if err != nil {
		log.Println(err.Error())
		panic("Failed to decode json payload as merge request")
	}
	log.Println("Gitlab Merge Request event payload: ", gitlabMergeRequestEvent)

	mergeRequest := &gitlabMergeRequestEvent.MergeRequest
	message := getMergeRequestMessage(mergeRequest)
	payload := map[string]string{"text": message}

	configuration := readConfiguration()

	sendChatMessage(configuration.Chat.WebhookURL, payload)
}

func sendChatMessage(webhookURL string, payload map[string]string) {

	jsonPayload, _ := json.Marshal(payload)
	log.Println("Payload of message to be sent to Hangouts Chat webhook: ", string(jsonPayload))

	_, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonPayload))

	if err != nil {
		panic("Failed to post chat message.")
	}
	log.Println("Sent message to Hangouts chat!")
}

func getMergeRequestMessage(mr *GitlabMergeRequestAttributes) string {

	messageTemplateBody := `Merge Request !{{.ID}} created to merge code from {{.SourceBranch}} to {{.TargetBranch}}. {{ if .Assignee.Name }}Assigned to: {{ .Assignee.Name }}.{{ end }}`
	messageTemplate, err := template.New("Merge Request message template").Parse(messageTemplateBody)

	if err != nil {
		panic("Failed to parse message template")
	}

	var buffer bytes.Buffer
	messageTemplate.Execute(&buffer, mr)

	return buffer.String()
}

func readConfiguration() *config.Configuration {
	var configuration config.Configuration
	config.ReadFile(&configuration)
	return &configuration
}
