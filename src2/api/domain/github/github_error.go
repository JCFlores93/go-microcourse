package github

type GithubErrorResponse struct {
	StatusCode    int           `json:"status_code"`
	Message       string        `json:"message"`
	Documentation string        `json:"documentation_url"`
	Errors        []GithubError `json:"errors"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}
