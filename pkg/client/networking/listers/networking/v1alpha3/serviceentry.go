// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha3

import (
	v1alpha3 "github.com/google/kf/v2/pkg/apis/networking/v1alpha3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ServiceEntryLister helps list ServiceEntries.
// All objects returned here must be treated as read-only.
type ServiceEntryLister interface {
	// List lists all ServiceEntries in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error)
	// ServiceEntries returns an object that can list and get ServiceEntries.
	ServiceEntries(namespace string) ServiceEntryNamespaceLister
	ServiceEntryListerExpansion
}

// serviceEntryLister implements the ServiceEntryLister interface.
type serviceEntryLister struct {
	indexer cache.Indexer
}

// NewServiceEntryLister returns a new ServiceEntryLister.
func NewServiceEntryLister(indexer cache.Indexer) ServiceEntryLister {
	return &serviceEntryLister{indexer: indexer}
}

// List lists all ServiceEntries in the indexer.
func (s *serviceEntryLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceEntry))
	})
	return ret, err
}

// ServiceEntries returns an object that can list and get ServiceEntries.
func (s *serviceEntryLister) ServiceEntries(namespace string) ServiceEntryNamespaceLister {
	return serviceEntryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ServiceEntryNamespaceLister helps list and get ServiceEntries.
// All objects returned here must be treated as read-only.
type ServiceEntryNamespaceLister interface {
	// List lists all ServiceEntries in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error)
	// Get retrieves the ServiceEntry from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha3.ServiceEntry, error)
	ServiceEntryNamespaceListerExpansion
}

// serviceEntryNamespaceLister implements the ServiceEntryNamespaceLister
// interface.
type serviceEntryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ServiceEntries in the indexer for a given namespace.
func (s serviceEntryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha3.ServiceEntry, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha3.ServiceEntry))
	})
	return ret, err
}

// Get retrieves the ServiceEntry from the indexer for a given namespace and name.
func (s serviceEntryNamespaceLister) Get(name string) (*v1alpha3.ServiceEntry, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha3.Resource("serviceentry"), name)
	}
	return obj.(*v1alpha3.ServiceEntry), nil
}
