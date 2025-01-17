package store

import (
	"github.com/gimlet-io/gimlet-cli/pkg/gimletd/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitopsCommitCRUD(t *testing.T) {
	s := NewTest()
	defer func() {
		s.Close()
	}()

	gitopsCommit := &model.GitopsCommit{
		Sha: "sha",
	}

	err := s.SaveOrUpdateGitopsCommit(gitopsCommit)
	assert.Nil(t, err)

	gitopsCommit.Status = "aStatus"
	err = s.SaveOrUpdateGitopsCommit(gitopsCommit)
	assert.Nil(t, err)

	savedGitopsCommit, err := s.GitopsCommit("sha")
	assert.Nil(t, err)
	assert.Equal(t, "aStatus", savedGitopsCommit.Status)
}
