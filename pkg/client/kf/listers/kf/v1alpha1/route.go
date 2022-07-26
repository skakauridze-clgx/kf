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

package v1alpha1

import (
	v1alpha1 "github.com/google/kf/v2/pkg/apis/kf/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RouteLister helps list Routes.
// All objects returned here must be treated as read-only.
type RouteLister interface {
	// List lists all Routes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Route, err error)
	// Routes returns an object that can list and get Routes.
	Routes(namespace string) RouteNamespaceLister
	RouteListerExpansion
}

// routeLister implements the RouteLister interface.
type routeLister struct {
	indexer cache.Indexer
}

// NewRouteLister returns a new RouteLister.
func NewRouteLister(indexer cache.Indexer) RouteLister {
	return &routeLister{indexer: indexer}
}

// List lists all Routes in the indexer.
func (s *routeLister) List(selector labels.Selector) (ret []*v1alpha1.Route, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route))
	})
	return ret, err
}

// Routes returns an object that can list and get Routes.
func (s *routeLister) Routes(namespace string) RouteNamespaceLister {
	return routeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RouteNamespaceLister helps list and get Routes.
// All objects returned here must be treated as read-only.
type RouteNamespaceLister interface {
	// List lists all Routes in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.Route, err error)
	// Get retrieves the Route from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.Route, error)
	RouteNamespaceListerExpansion
}

// routeNamespaceLister implements the RouteNamespaceLister
// interface.
type routeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Routes in the indexer for a given namespace.
func (s routeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Route, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Route))
	})
	return ret, err
}

// Get retrieves the Route from the indexer for a given namespace and name.
func (s routeNamespaceLister) Get(name string) (*v1alpha1.Route, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("route"), name)
	}
	return obj.(*v1alpha1.Route), nil
}
