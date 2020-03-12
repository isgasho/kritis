/*
Copyright 2018 Google LLC

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

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	v1beta1 "github.com/grafeas/kritis/pkg/kritis/apis/kritis/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// VulnzSigningPolicyLister helps list VulnzSigningPolicies.
type VulnzSigningPolicyLister interface {
	// List lists all VulnzSigningPolicies in the indexer.
	List(selector labels.Selector) (ret []*v1beta1.VulnzSigningPolicy, err error)
	// VulnzSigningPolicies returns an object that can list and get VulnzSigningPolicies.
	VulnzSigningPolicies(namespace string) VulnzSigningPolicyNamespaceLister
	VulnzSigningPolicyListerExpansion
}

// vulnzSigningPolicyLister implements the VulnzSigningPolicyLister interface.
type vulnzSigningPolicyLister struct {
	indexer cache.Indexer
}

// NewVulnzSigningPolicyLister returns a new VulnzSigningPolicyLister.
func NewVulnzSigningPolicyLister(indexer cache.Indexer) VulnzSigningPolicyLister {
	return &vulnzSigningPolicyLister{indexer: indexer}
}

// List lists all VulnzSigningPolicies in the indexer.
func (s *vulnzSigningPolicyLister) List(selector labels.Selector) (ret []*v1beta1.VulnzSigningPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VulnzSigningPolicy))
	})
	return ret, err
}

// VulnzSigningPolicies returns an object that can list and get VulnzSigningPolicies.
func (s *vulnzSigningPolicyLister) VulnzSigningPolicies(namespace string) VulnzSigningPolicyNamespaceLister {
	return vulnzSigningPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// VulnzSigningPolicyNamespaceLister helps list and get VulnzSigningPolicies.
type VulnzSigningPolicyNamespaceLister interface {
	// List lists all VulnzSigningPolicies in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1beta1.VulnzSigningPolicy, err error)
	// Get retrieves the VulnzSigningPolicy from the indexer for a given namespace and name.
	Get(name string) (*v1beta1.VulnzSigningPolicy, error)
	VulnzSigningPolicyNamespaceListerExpansion
}

// vulnzSigningPolicyNamespaceLister implements the VulnzSigningPolicyNamespaceLister
// interface.
type vulnzSigningPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all VulnzSigningPolicies in the indexer for a given namespace.
func (s vulnzSigningPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.VulnzSigningPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.VulnzSigningPolicy))
	})
	return ret, err
}

// Get retrieves the VulnzSigningPolicy from the indexer for a given namespace and name.
func (s vulnzSigningPolicyNamespaceLister) Get(name string) (*v1beta1.VulnzSigningPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("vulnzsigningpolicy"), name)
	}
	return obj.(*v1beta1.VulnzSigningPolicy), nil
}