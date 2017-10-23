<<<<<<< HEAD
// Copyright 2017 The Kubernetes Dashboard Authors.
=======
// Copyright 2017 The Kubernetes Authors.
>>>>>>> upstream/master
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package sync

import (
	"testing"

	"k8s.io/client-go/kubernetes/fake"
)

func TestNewSynchronizerManager(t *testing.T) {
	manager := NewSynchronizerManager(fake.NewSimpleClientset())
	if manager == nil {
		t.Fatalf("NewSynchronizerManager(): Expected synchronizer manager not to be nil")
	}
}

func TestSynchronizerManager_Secret(t *testing.T) {
	manager := NewSynchronizerManager(fake.NewSimpleClientset())
	if manager.Secret("", "") == nil {
		t.Fatalf("Secret(%s, %s): Expected secret synchronizer not to be nil", "", "")
	}
}
