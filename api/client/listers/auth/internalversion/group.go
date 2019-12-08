/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
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
	auth "tkestack.io/tke/api/auth"
)

// GroupLister helps list Groups.
type GroupLister interface {
	// List lists all Groups in the indexer.
	List(selector labels.Selector) (ret []*auth.Group, err error)
	// Get retrieves the Group from the index for a given name.
	Get(name string) (*auth.Group, error)
	GroupListerExpansion
}

// groupLister implements the GroupLister interface.
type groupLister struct {
	indexer cache.Indexer
}

// NewGroupLister returns a new GroupLister.
func NewGroupLister(indexer cache.Indexer) GroupLister {
	return &groupLister{indexer: indexer}
}

// List lists all Groups in the indexer.
func (s *groupLister) List(selector labels.Selector) (ret []*auth.Group, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*auth.Group))
	})
	return ret, err
}

// Get retrieves the Group from the index for a given name.
func (s *groupLister) Get(name string) (*auth.Group, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(auth.Resource("group"), name)
	}
	return obj.(*auth.Group), nil
}
