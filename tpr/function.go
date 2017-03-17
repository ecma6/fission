/*
Copyright 2016 The Fission Authors.

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

package tpr

import (
	"k8s.io/client-go/1.5/pkg/api"
	"k8s.io/client-go/1.5/pkg/watch"
	"k8s.io/client-go/1.5/rest"
)

type (
	FunctionInterface interface {
		Create(*Function) (*Function, error)
		Get(name string) (*Function, error)
		Update(*Function) (*Function, error)
		Delete(name string, options *api.DeleteOptions) error
		List(opts api.ListOptions) (*FunctionList, error)
		Watch(opts api.ListOptions) (watch.Interface, error)
	}

	functionClient struct {
		client    *rest.RESTClient
		namespace string
	}
)

func MakeFunctionInterface(tprClient *rest.RESTClient, namespace string) FunctionInterface {
	return &functionClient{
		client:    tprClient,
		namespace: namespace,
	}
}

func (fc *functionClient) Create(f *Function) (*Function, error) {
	var result Function
	err := fc.client.Post().
		Resource("functions").
		Namespace(fc.namespace).
		Body(f).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (fc *functionClient) Get(name string) (*Function, error) {
	var result Function
	err := fc.client.Get().
		Resource("functions").
		Namespace(fc.namespace).
		Name(name).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (fc *functionClient) Update(f *Function) (*Function, error) {
	var result Function
	err := fc.client.Put().
		Resource("functions").
		Namespace(fc.namespace).
		Name(f.Metadata.Name).
		Body(f).
		Do().Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (fc *functionClient) Delete(name string, opts *api.DeleteOptions) error {
	return fc.client.Delete().
		Namespace(fc.namespace).
		Resource("functions").
		Name(name).
		Body(opts).
		Do().
		Error()
}

func (fc *functionClient) List(opts api.ListOptions) (*FunctionList, error) {
	var result FunctionList
	err := fc.client.Get().
		Namespace(fc.namespace).
		Resource("functions").
		VersionedParams(&opts, api.ParameterCodec).
		Do().
		Into(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (fc *functionClient) Watch(opts api.ListOptions) (watch.Interface, error) {
	return fc.client.Get().
		Prefix("watch").
		Namespace(fc.namespace).
		Resource("functions").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}
