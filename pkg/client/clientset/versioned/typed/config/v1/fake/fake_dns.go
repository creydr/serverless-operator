// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "github.com/openshift/api/config/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDNSs implements DNSInterface
type FakeDNSs struct {
	Fake *FakeConfigV1
}

var dnssResource = v1.SchemeGroupVersion.WithResource("dnss")

var dnssKind = v1.SchemeGroupVersion.WithKind("DNS")

// Get takes name of the dNS, and returns the corresponding dNS object, and an error if there is any.
func (c *FakeDNSs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(dnssResource, name), &v1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DNS), err
}

// List takes label and field selectors, and returns the list of DNSs that match those selectors.
func (c *FakeDNSs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DNSList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(dnssResource, dnssKind, opts), &v1.DNSList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DNSList{ListMeta: obj.(*v1.DNSList).ListMeta}
	for _, item := range obj.(*v1.DNSList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dNSs.
func (c *FakeDNSs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(dnssResource, opts))
}

// Create takes the representation of a dNS and creates it.  Returns the server's representation of the dNS, and an error, if there is any.
func (c *FakeDNSs) Create(ctx context.Context, dNS *v1.DNS, opts metav1.CreateOptions) (result *v1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(dnssResource, dNS), &v1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DNS), err
}

// Update takes the representation of a dNS and updates it. Returns the server's representation of the dNS, and an error, if there is any.
func (c *FakeDNSs) Update(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (result *v1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(dnssResource, dNS), &v1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DNS), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDNSs) UpdateStatus(ctx context.Context, dNS *v1.DNS, opts metav1.UpdateOptions) (*v1.DNS, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(dnssResource, "status", dNS), &v1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DNS), err
}

// Delete takes name of the dNS and deletes it. Returns an error if one occurs.
func (c *FakeDNSs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(dnssResource, name, opts), &v1.DNS{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDNSs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(dnssResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DNSList{})
	return err
}

// Patch applies the patch and returns the patched dNS.
func (c *FakeDNSs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.DNS, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(dnssResource, name, pt, data, subresources...), &v1.DNS{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.DNS), err
}
