package tests

import (
	"encoding/json"
	"io/ioutil"
	"testing"

	"github.com/idl99/hangouts-chat-gitlab-bot/handlers"
	"github.com/stretchr/testify/assert"
)

func TestMergeRequestStruct(t *testing.T) {

	var mergeRequestEvent handlers.GitlabMergeRequestEvent

	data, err := ioutil.ReadFile("data/merge_request_payload.json")
	if err != nil {
		t.Errorf("Error occurred when trying to read test data from test file")
	}

	err = json.Unmarshal(data, &mergeRequestEvent)
	if err != nil {
		t.Errorf("Error occurred when trying to unmarshal")
	}

	assert.Equal(t, "merge_request", mergeRequestEvent.ObjectKind)

	mergeRequest := mergeRequestEvent.MergeRequest
	assert.Equal(t, "master", mergeRequest.TargetBranch)
	assert.Equal(t, "ms-viewport", mergeRequest.SourceBranch)
	assert.Equal(t, 51, mergeRequest.AuthorID)
	assert.Equal(t, 6, mergeRequest.AssigneeID)
	assert.Equal(t, "MS-Viewport", mergeRequest.Title)
	assert.Equal(t, false, mergeRequest.WorkInProgress)

	assignee := mergeRequest.Assignee
	assert.Equal(t, "User1", assignee.Name)
	assert.Equal(t, "user1", assignee.Username)

}
