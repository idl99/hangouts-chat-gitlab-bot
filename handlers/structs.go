package handlers

// GitlabEvent - struct to represent generic structure of Gitlab Webhook event
type GitlabEvent struct {
	ObjectKind string     `json:"object_kind"`
	User       User       `json:"user"`
	Project    Project    `json:"project"`
	Repository Repository `json:"repository"`
}

// GitlabMergeRequestEvent - struct to represent Gitlab merge request event
type GitlabMergeRequestEvent struct {
	GitlabEvent
	MergeRequest GitlabMergeRequestAttributes `json:"object_attributes"`
}

// GitlabMergeRequestAttributes - struct to represent attributes of the merge request event
type GitlabMergeRequestAttributes struct {
	ID             int    `json:"id"`
	TargetBranch   string `json:"target_branch"`
	SourceBranch   string `json:"source_branch"`
	AuthorID       int    `json:"author_id"`
	AssigneeID     int    `json:"assignee_id"`
	Title          string `json:"title"`
	CreatedAt      string `json:"created_at"`
	Description    string `json:"description"`
	WorkInProgress bool   `json:"work_in_progress"`
	URL            string `json:"url"`
	Assignee       User   `json:"assignee"`
}

// User - details of the Gitlab user related to the Webhook event
type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

// Project - details of the Gitlab project for which are using a Webhook
type Project struct {
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	WebURL            string      `json:"web_url"`
	AvatarURL         interface{} `json:"avatar_url"`
	GitSSHURL         string      `json:"git_ssh_url"`
	GitHTTPURL        string      `json:"git_http_url"`
	Namespace         string      `json:"namespace"`
	VisibilityLevel   int         `json:"visibility_level"`
	PathWithNamespace string      `json:"path_with_namespace"`
	DefaultBranch     string      `json:"default_branch"`
	Homepage          string      `json:"homepage"`
	URL               string      `json:"url"`
	SSHURL            string      `json:"ssh_url"`
	HTTPURL           string      `json:"http_url"`
	ID                *int        `json:"id,omitempty"`
}

// Repository - details of the Gitlab repository for which we are using a Webhook
type Repository struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
}
