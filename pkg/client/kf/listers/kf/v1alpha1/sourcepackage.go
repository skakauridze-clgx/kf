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

// SourcePackageLister helps list SourcePackages.
// All objects returned here must be treated as read-only.
type SourcePackageLister interface {
	// List lists all SourcePackages in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SourcePackage, err error)
	// SourcePackages returns an object that can list and get SourcePackages.
	SourcePackages(namespace string) SourcePackageNamespaceLister
	SourcePackageListerExpansion
}

// sourcePackageLister implements the SourcePackageLister interface.
type sourcePackageLister struct {
	indexer cache.Indexer
}

// NewSourcePackageLister returns a new SourcePackageLister.
func NewSourcePackageLister(indexer cache.Indexer) SourcePackageLister {
	return &sourcePackageLister{indexer: indexer}
}

// List lists all SourcePackages in the indexer.
func (s *sourcePackageLister) List(selector labels.Selector) (ret []*v1alpha1.SourcePackage, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcePackage))
	})
	return ret, err
}

// SourcePackages returns an object that can list and get SourcePackages.
func (s *sourcePackageLister) SourcePackages(namespace string) SourcePackageNamespaceLister {
	return sourcePackageNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// SourcePackageNamespaceLister helps list and get SourcePackages.
// All objects returned here must be treated as read-only.
type SourcePackageNamespaceLister interface {
	// List lists all SourcePackages in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.SourcePackage, err error)
	// Get retrieves the SourcePackage from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.SourcePackage, error)
	SourcePackageNamespaceListerExpansion
}

// sourcePackageNamespaceLister implements the SourcePackageNamespaceLister
// interface.
type sourcePackageNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all SourcePackages in the indexer for a given namespace.
func (s sourcePackageNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.SourcePackage, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.SourcePackage))
	})
	return ret, err
}

// Get retrieves the SourcePackage from the indexer for a given namespace and name.
func (s sourcePackageNamespaceLister) Get(name string) (*v1alpha1.SourcePackage, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("sourcepackage"), name)
	}
	return obj.(*v1alpha1.SourcePackage), nil
}
