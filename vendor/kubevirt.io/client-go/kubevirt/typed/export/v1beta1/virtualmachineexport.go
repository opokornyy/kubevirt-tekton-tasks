/*
Copyright The KubeVirt Authors.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	v1beta1 "kubevirt.io/api/export/v1beta1"
	scheme "kubevirt.io/client-go/kubevirt/scheme"
)

// VirtualMachineExportsGetter has a method to return a VirtualMachineExportInterface.
// A group's client should implement this interface.
type VirtualMachineExportsGetter interface {
	VirtualMachineExports(namespace string) VirtualMachineExportInterface
}

// VirtualMachineExportInterface has methods to work with VirtualMachineExport resources.
type VirtualMachineExportInterface interface {
	Create(ctx context.Context, virtualMachineExport *v1beta1.VirtualMachineExport, opts v1.CreateOptions) (*v1beta1.VirtualMachineExport, error)
	Update(ctx context.Context, virtualMachineExport *v1beta1.VirtualMachineExport, opts v1.UpdateOptions) (*v1beta1.VirtualMachineExport, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, virtualMachineExport *v1beta1.VirtualMachineExport, opts v1.UpdateOptions) (*v1beta1.VirtualMachineExport, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.VirtualMachineExport, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.VirtualMachineExportList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.VirtualMachineExport, err error)
	VirtualMachineExportExpansion
}

// virtualMachineExports implements VirtualMachineExportInterface
type virtualMachineExports struct {
	*gentype.ClientWithList[*v1beta1.VirtualMachineExport, *v1beta1.VirtualMachineExportList]
}

// newVirtualMachineExports returns a VirtualMachineExports
func newVirtualMachineExports(c *ExportV1beta1Client, namespace string) *virtualMachineExports {
	return &virtualMachineExports{
		gentype.NewClientWithList[*v1beta1.VirtualMachineExport, *v1beta1.VirtualMachineExportList](
			"virtualmachineexports",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1beta1.VirtualMachineExport { return &v1beta1.VirtualMachineExport{} },
			func() *v1beta1.VirtualMachineExportList { return &v1beta1.VirtualMachineExportList{} }),
	}
}
