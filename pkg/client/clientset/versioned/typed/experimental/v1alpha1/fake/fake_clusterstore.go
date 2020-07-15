/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/pivotal/kpack/pkg/apis/experimental/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterStores implements ClusterStoreInterface
type FakeClusterStores struct {
	Fake *FakeExperimentalV1alpha1
}

var clusterstoresResource = schema.GroupVersionResource{Group: "experimental.kpack.pivotal.io", Version: "v1alpha1", Resource: "clusterstores"}

var clusterstoresKind = schema.GroupVersionKind{Group: "experimental.kpack.pivotal.io", Version: "v1alpha1", Kind: "ClusterStore"}

// Get takes name of the clusterStore, and returns the corresponding clusterStore object, and an error if there is any.
func (c *FakeClusterStores) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusterstoresResource, name), &v1alpha1.ClusterStore{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStore), err
}

// List takes label and field selectors, and returns the list of ClusterStores that match those selectors.
func (c *FakeClusterStores) List(opts v1.ListOptions) (result *v1alpha1.ClusterStoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusterstoresResource, clusterstoresKind, opts), &v1alpha1.ClusterStoreList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ClusterStoreList{ListMeta: obj.(*v1alpha1.ClusterStoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.ClusterStoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterStores.
func (c *FakeClusterStores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusterstoresResource, opts))
}

// Create takes the representation of a clusterStore and creates it.  Returns the server's representation of the clusterStore, and an error, if there is any.
func (c *FakeClusterStores) Create(clusterStore *v1alpha1.ClusterStore) (result *v1alpha1.ClusterStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusterstoresResource, clusterStore), &v1alpha1.ClusterStore{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStore), err
}

// Update takes the representation of a clusterStore and updates it. Returns the server's representation of the clusterStore, and an error, if there is any.
func (c *FakeClusterStores) Update(clusterStore *v1alpha1.ClusterStore) (result *v1alpha1.ClusterStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusterstoresResource, clusterStore), &v1alpha1.ClusterStore{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStore), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterStores) UpdateStatus(clusterStore *v1alpha1.ClusterStore) (*v1alpha1.ClusterStore, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusterstoresResource, "status", clusterStore), &v1alpha1.ClusterStore{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStore), err
}

// Delete takes name of the clusterStore and deletes it. Returns an error if one occurs.
func (c *FakeClusterStores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(clusterstoresResource, name), &v1alpha1.ClusterStore{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterStores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusterstoresResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ClusterStoreList{})
	return err
}

// Patch applies the patch and returns the patched clusterStore.
func (c *FakeClusterStores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterStore, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusterstoresResource, name, pt, data, subresources...), &v1alpha1.ClusterStore{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ClusterStore), err
}
