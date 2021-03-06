// Copyright 2017 The etcd Authors
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

//go:build v2v3
// +build v2v3

package v2store_test

import (
	"testing"

	"go.etcd.io/etcd/server/v3/etcdserver/api/v2store"
	"go.etcd.io/etcd/server/v3/etcdserver/api/v2v3"
	integration2 "go.etcd.io/etcd/tests/v3/framework/integration"
)

type v2v3TestStore struct {
	v2store.Store
	clus *integration2.ClusterV3
	t    *testing.T
}

func (s *v2v3TestStore) Close() { s.clus.Terminate(s.t) }

func newTestStore(t *testing.T, ns ...string) StoreCloser {
	integration2.BeforeTest(t)
	clus := integration2.NewClusterV3(t, &integration2.ClusterConfig{Size: 1})
	return &v2v3TestStore{
		v2v3.NewStore(clus.Client(0), "/v2/"),
		clus,
		t,
	}
}
