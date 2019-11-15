package github

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequest(t *testing.T) {
	request := CreateRepoRequest{
		Name:        "golang introduction",
		Description: "A golang introduction repository",
		Homepage:    "https://github.com",
		Private:     true,
		HasIssues:   true,
		HasProjects: true,
		HasWiki:     true,
	}

	if request.Private {

	}

	// Mashal takes an input interface and attempts to create a valid json string
	bytes, err := json.Marshal(request)
	assert.Nil(t, err)
	assert.NotNil(t, bytes)
	fmt.Println(string(bytes))

	var target CreateRepoRequest
	err = json.Unmarshal(bytes, target)
	assert.Nil(t, err)

	assert.EqualValues(t, target.Name, request.Name)
}
