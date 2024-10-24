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

// Code generated by client-gen. DO NOT EDIT.

package v1beta1

import (
	context "context"

	storagev1beta1 "k8s.io/api/storage/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationsstoragev1beta1 "k8s.io/client-go/applyconfigurations/storage/v1beta1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// StorageClassesGetter has a method to return a StorageClassInterface.
// A group's client should implement this interface.
type StorageClassesGetter interface {
	StorageClasses() StorageClassInterface
}

// StorageClassInterface has methods to work with StorageClass resources.
type StorageClassInterface interface {
	Create(ctx context.Context, storageClass *storagev1beta1.StorageClass, opts v1.CreateOptions) (*storagev1beta1.StorageClass, error)
	Update(ctx context.Context, storageClass *storagev1beta1.StorageClass, opts v1.UpdateOptions) (*storagev1beta1.StorageClass, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*storagev1beta1.StorageClass, error)
	List(ctx context.Context, opts v1.ListOptions) (*storagev1beta1.StorageClassList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *storagev1beta1.StorageClass, err error)
	Apply(ctx context.Context, storageClass *applyconfigurationsstoragev1beta1.StorageClassApplyConfiguration, opts v1.ApplyOptions) (result *storagev1beta1.StorageClass, err error)
	StorageClassExpansion
}

// storageClasses implements StorageClassInterface
type storageClasses struct {
	*gentype.ClientWithListAndApply[*storagev1beta1.StorageClass, *storagev1beta1.StorageClassList, *applyconfigurationsstoragev1beta1.StorageClassApplyConfiguration]
}

// newStorageClasses returns a StorageClasses
func newStorageClasses(c *StorageV1beta1Client) *storageClasses {
	return &storageClasses{
		gentype.NewClientWithListAndApply[*storagev1beta1.StorageClass, *storagev1beta1.StorageClassList, *applyconfigurationsstoragev1beta1.StorageClassApplyConfiguration](
			"storageclasses",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *storagev1beta1.StorageClass { return &storagev1beta1.StorageClass{} },
			func() *storagev1beta1.StorageClassList { return &storagev1beta1.StorageClassList{} }),
	}
}
