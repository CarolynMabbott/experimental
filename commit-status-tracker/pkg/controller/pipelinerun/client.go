// Copyright 2020 The Tekton Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package pipelinerun

import (
	"context"

	"github.com/jenkins-x/go-scm/scm"
	"github.com/jenkins-x/go-scm/scm/driver/github"
	"github.com/jenkins-x/go-scm/scm/driver/gitlab"
	"golang.org/x/oauth2"
)

type scmClientFactory func(string, string) *scm.Client

func createClient(token, repoType string) *scm.Client {
	var client *scm.Client
	if repoType == "github" {
		client = github.NewDefault()
	} else if repoType == "gitlab" {
		client = gitlab.NewDefault()
	} else {
		return nil
	}
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	client.Client = oauth2.NewClient(context.Background(), ts)
	return client
}
