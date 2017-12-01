/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"fmt"
	"strings"
)

// Tide is config for the tide pool.
type Tide struct {
	// Queries must not overlap. It must be impossible for any two queries to
	// ever return the same PR.
	// TODO: This will only be possible when we allow specifying orgs. At that
	//       point, verify the above condition.
	Queries []TideQuery `json:"queries,omitempty"`
}

// TideQuery is turned into a GitHub search query. See the docs for details:
// https://help.github.com/articles/searching-issues-and-pull-requests/
type TideQuery struct {
	Repos []string `json:"repos,omitempty"`

	Labels        []string `json:"labels,omitempty"`
	MissingLabels []string `json:"missingLabels,omitempty"`

	ReviewApprovedRequired bool `json:"reviewApprovedRequired,omitempty"`
}

func (tq *TideQuery) Query() string {
	toks := []string{"is:pr", "state:open"}
	for _, r := range tq.Repos {
		toks = append(toks, fmt.Sprintf("repo:\"%s\"", r))
	}
	for _, l := range tq.Labels {
		toks = append(toks, fmt.Sprintf("label:\"%s\"", l))
	}
	for _, l := range tq.MissingLabels {
		toks = append(toks, fmt.Sprintf("-label:\"%s\"", l))
	}
	if tq.ReviewApprovedRequired {
		toks = append(toks, "review:approved")
	}
	return strings.Join(toks, " ")
}
