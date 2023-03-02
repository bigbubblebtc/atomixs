// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0
// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1beta2 "github.com/atomix/atomix/stores/shared-memory/pkg/client/clientset/versioned/typed/sharedmemory/v1beta2"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeSharedmemoryV1beta2 struct {
	*testing.Fake
}

func (c *FakeSharedmemoryV1beta2) SharedMemoryStores(namespace string) v1beta2.SharedMemoryStoreInterface {
	return &FakeSharedMemoryStores{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeSharedmemoryV1beta2) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}