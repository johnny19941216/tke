/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2020 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by lister-gen. DO NOT EDIT.

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	notify "tkestack.io/tke/api/notify"
)

// TemplateLister helps list Templates.
// All objects returned here must be treated as read-only.
type TemplateLister interface {
	// List lists all Templates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*notify.Template, err error)
	// Templates returns an object that can list and get Templates.
	Templates(namespace string) TemplateNamespaceLister
	TemplateListerExpansion
}

// templateLister implements the TemplateLister interface.
type templateLister struct {
	indexer cache.Indexer
}

// NewTemplateLister returns a new TemplateLister.
func NewTemplateLister(indexer cache.Indexer) TemplateLister {
	return &templateLister{indexer: indexer}
}

// List lists all Templates in the indexer.
func (s *templateLister) List(selector labels.Selector) (ret []*notify.Template, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*notify.Template))
	})
	return ret, err
}

// Templates returns an object that can list and get Templates.
func (s *templateLister) Templates(namespace string) TemplateNamespaceLister {
	return templateNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// TemplateNamespaceLister helps list and get Templates.
// All objects returned here must be treated as read-only.
type TemplateNamespaceLister interface {
	// List lists all Templates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*notify.Template, err error)
	// Get retrieves the Template from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*notify.Template, error)
	TemplateNamespaceListerExpansion
}

// templateNamespaceLister implements the TemplateNamespaceLister
// interface.
type templateNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Templates in the indexer for a given namespace.
func (s templateNamespaceLister) List(selector labels.Selector) (ret []*notify.Template, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*notify.Template))
	})
	return ret, err
}

// Get retrieves the Template from the indexer for a given namespace and name.
func (s templateNamespaceLister) Get(name string) (*notify.Template, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(notify.Resource("template"), name)
	}
	return obj.(*notify.Template), nil
}
