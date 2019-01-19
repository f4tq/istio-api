// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha2

import (
	v1alpha2 "istio.io/api/pkg/kube/apis/config/v1alpha2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// AttributeManifestLister helps list AttributeManifests.
type AttributeManifestLister interface {
	// List lists all AttributeManifests in the indexer.
	List(selector labels.Selector) (ret []*v1alpha2.AttributeManifest, err error)
	// AttributeManifests returns an object that can list and get AttributeManifests.
	AttributeManifests(namespace string) AttributeManifestNamespaceLister
	AttributeManifestListerExpansion
}

// attributeManifestLister implements the AttributeManifestLister interface.
type attributeManifestLister struct {
	indexer cache.Indexer
}

// NewAttributeManifestLister returns a new AttributeManifestLister.
func NewAttributeManifestLister(indexer cache.Indexer) AttributeManifestLister {
	return &attributeManifestLister{indexer: indexer}
}

// List lists all AttributeManifests in the indexer.
func (s *attributeManifestLister) List(selector labels.Selector) (ret []*v1alpha2.AttributeManifest, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AttributeManifest))
	})
	return ret, err
}

// AttributeManifests returns an object that can list and get AttributeManifests.
func (s *attributeManifestLister) AttributeManifests(namespace string) AttributeManifestNamespaceLister {
	return attributeManifestNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// AttributeManifestNamespaceLister helps list and get AttributeManifests.
type AttributeManifestNamespaceLister interface {
	// List lists all AttributeManifests in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha2.AttributeManifest, err error)
	// Get retrieves the AttributeManifest from the indexer for a given namespace and name.
	Get(name string) (*v1alpha2.AttributeManifest, error)
	AttributeManifestNamespaceListerExpansion
}

// attributeManifestNamespaceLister implements the AttributeManifestNamespaceLister
// interface.
type attributeManifestNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all AttributeManifests in the indexer for a given namespace.
func (s attributeManifestNamespaceLister) List(selector labels.Selector) (ret []*v1alpha2.AttributeManifest, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha2.AttributeManifest))
	})
	return ret, err
}

// Get retrieves the AttributeManifest from the indexer for a given namespace and name.
func (s attributeManifestNamespaceLister) Get(name string) (*v1alpha2.AttributeManifest, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha2.Resource("attributemanifest"), name)
	}
	return obj.(*v1alpha2.AttributeManifest), nil
}
