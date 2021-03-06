/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "gitlab.com/project-emco/core/emco-base/src/monitor/pkg/apis/k8splugin/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ResourceBundleStateLister helps list ResourceBundleStates.
type ResourceBundleStateLister interface {
	// List lists all ResourceBundleStates in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceBundleState, err error)
	// ResourceBundleStates returns an object that can list and get ResourceBundleStates.
	ResourceBundleStates(namespace string) ResourceBundleStateNamespaceLister
	ResourceBundleStateListerExpansion
}

// resourceBundleStateLister implements the ResourceBundleStateLister interface.
type resourceBundleStateLister struct {
	indexer cache.Indexer
}

// NewResourceBundleStateLister returns a new ResourceBundleStateLister.
func NewResourceBundleStateLister(indexer cache.Indexer) ResourceBundleStateLister {
	return &resourceBundleStateLister{indexer: indexer}
}

// List lists all ResourceBundleStates in the indexer.
func (s *resourceBundleStateLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceBundleState, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceBundleState))
	})
	return ret, err
}

// ResourceBundleStates returns an object that can list and get ResourceBundleStates.
func (s *resourceBundleStateLister) ResourceBundleStates(namespace string) ResourceBundleStateNamespaceLister {
	return resourceBundleStateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ResourceBundleStateNamespaceLister helps list and get ResourceBundleStates.
type ResourceBundleStateNamespaceLister interface {
	// List lists all ResourceBundleStates in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.ResourceBundleState, err error)
	// Get retrieves the ResourceBundleState from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.ResourceBundleState, error)
	ResourceBundleStateNamespaceListerExpansion
}

// resourceBundleStateNamespaceLister implements the ResourceBundleStateNamespaceLister
// interface.
type resourceBundleStateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ResourceBundleStates in the indexer for a given namespace.
func (s resourceBundleStateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ResourceBundleState, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ResourceBundleState))
	})
	return ret, err
}

// Get retrieves the ResourceBundleState from the indexer for a given namespace and name.
func (s resourceBundleStateNamespaceLister) Get(name string) (*v1alpha1.ResourceBundleState, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("resourcebundlestate"), name)
	}
	return obj.(*v1alpha1.ResourceBundleState), nil
}
