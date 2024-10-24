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

package v1

import (
	context "context"

	certificatesv1 "k8s.io/api/certificates/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	applyconfigurationscertificatesv1 "k8s.io/client-go/applyconfigurations/certificates/v1"
	gentype "k8s.io/client-go/gentype"
	scheme "k8s.io/client-go/kubernetes/scheme"
)

// CertificateSigningRequestsGetter has a method to return a CertificateSigningRequestInterface.
// A group's client should implement this interface.
type CertificateSigningRequestsGetter interface {
	CertificateSigningRequests() CertificateSigningRequestInterface
}

// CertificateSigningRequestInterface has methods to work with CertificateSigningRequest resources.
type CertificateSigningRequestInterface interface {
	Create(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.CreateOptions) (*certificatesv1.CertificateSigningRequest, error)
	Update(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*certificatesv1.CertificateSigningRequest, error)
	List(ctx context.Context, opts metav1.ListOptions) (*certificatesv1.CertificateSigningRequestList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *certificatesv1.CertificateSigningRequest, err error)
	Apply(ctx context.Context, certificateSigningRequest *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *certificatesv1.CertificateSigningRequest, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, certificateSigningRequest *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration, opts metav1.ApplyOptions) (result *certificatesv1.CertificateSigningRequest, err error)
	UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (*certificatesv1.CertificateSigningRequest, error)

	CertificateSigningRequestExpansion
}

// certificateSigningRequests implements CertificateSigningRequestInterface
type certificateSigningRequests struct {
	*gentype.ClientWithListAndApply[*certificatesv1.CertificateSigningRequest, *certificatesv1.CertificateSigningRequestList, *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration]
}

// newCertificateSigningRequests returns a CertificateSigningRequests
func newCertificateSigningRequests(c *CertificatesV1Client) *certificateSigningRequests {
	return &certificateSigningRequests{
		gentype.NewClientWithListAndApply[*certificatesv1.CertificateSigningRequest, *certificatesv1.CertificateSigningRequestList, *applyconfigurationscertificatesv1.CertificateSigningRequestApplyConfiguration](
			"certificatesigningrequests",
			c.RESTClient(),
			scheme.ParameterCodec,
			"",
			func() *certificatesv1.CertificateSigningRequest { return &certificatesv1.CertificateSigningRequest{} },
			func() *certificatesv1.CertificateSigningRequestList {
				return &certificatesv1.CertificateSigningRequestList{}
			}),
	}
}

// UpdateApproval takes the top resource name and the representation of a certificateSigningRequest and updates it. Returns the server's representation of the certificateSigningRequest, and an error, if there is any.
func (c *certificateSigningRequests) UpdateApproval(ctx context.Context, certificateSigningRequestName string, certificateSigningRequest *certificatesv1.CertificateSigningRequest, opts metav1.UpdateOptions) (result *certificatesv1.CertificateSigningRequest, err error) {
	result = &certificatesv1.CertificateSigningRequest{}
	err = c.GetClient().Put().
		Resource("certificatesigningrequests").
		Name(certificateSigningRequestName).
		SubResource("approval").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(certificateSigningRequest).
		Do(ctx).
		Into(result)
	return
}
