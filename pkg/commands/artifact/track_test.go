package artifact

import (
	"testing"

	"github.com/gimlet-io/gimlet-cli/pkg/dx"
	"gotest.tools/assert"
)

func Test_extractEndStateWithNoGitopsComimt(t *testing.T) {
	testReleaseStatus := dx.ReleaseStatus{
		Status:       "does not matter, the method checks the gitops commits only",
		StatusDesc:   "",
		GitopsHashes: []dx.GitopsStatus{},
		Results:      []dx.Result{},
	}

	everythingSucceeded, gitopsCommitsHaveFailed := testReleaseStatus.ExtractGitopsEndState()

	assert.Equal(t, everythingSucceeded, true)
	assert.Equal(t, gitopsCommitsHaveFailed, false)
}

func Test_extractEndStateIfGitopsCommitFailed(t *testing.T) {
	testReleaseStatus := dx.ReleaseStatus{
		Status:     "processed",
		StatusDesc: "",
		GitopsHashes: []dx.GitopsStatus{
			{
				Hash:       "abc123",
				Status:     "ReconciliationFailed",
				StatusDesc: "test description one",
			},
			{
				Hash:       "abc456",
				Status:     "ReconciliationSucceeded",
				StatusDesc: "test description two",
			},
		},
		Results: []dx.Result{
			{
				Hash:               "abc123",
				GitopsCommitStatus: "ReconciliationFailed",
				StatusDesc:         "test description one",
			},
			{
				Hash:               "abc456",
				GitopsCommitStatus: "ReconciliationSucceeded",
				StatusDesc:         "test description two",
			},
		},
	}

	everythingSucceeded, gitopsCommitsHaveFailed := testReleaseStatus.ExtractGitopsEndState()

	assert.Equal(t, everythingSucceeded, false)
	assert.Equal(t, gitopsCommitsHaveFailed, true)
}

func Test_extractEndStateIfGitopsCommitsSucceeded(t *testing.T) {
	testReleaseStatus := dx.ReleaseStatus{
		Status:     "processed",
		StatusDesc: "",
		GitopsHashes: []dx.GitopsStatus{
			{
				Hash:       "abc123",
				Status:     "ReconciliationSucceeded",
				StatusDesc: "test description one",
			},
			{
				Hash:       "abc456",
				Status:     "ReconciliationSucceeded",
				StatusDesc: "test description two",
			},
		},
		Results: []dx.Result{
			{
				Hash:               "abc123",
				GitopsCommitStatus: "ReconciliationSucceeded",
				StatusDesc:         "test description one",
			},
			{
				Hash:               "abc456",
				GitopsCommitStatus: "ReconciliationSucceeded",
				StatusDesc:         "test description two",
			},
		},
	}

	everythingSucceeded, gitopsCommitsHaveFailed := testReleaseStatus.ExtractGitopsEndState()

	assert.Equal(t, everythingSucceeded, true)
	assert.Equal(t, gitopsCommitsHaveFailed, false)
}
