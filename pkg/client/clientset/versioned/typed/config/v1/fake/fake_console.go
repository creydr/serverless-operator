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

// FakeConsoles implements ConsoleInterface
type FakeConsoles struct {
	Fake *FakeConfigV1
}

var consolesResource = v1.SchemeGroupVersion.WithResource("consoles")

var consolesKind = v1.SchemeGroupVersion.WithKind("Console")

// Get takes name of the console, and returns the corresponding console object, and an error if there is any.
func (c *FakeConsoles) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(consolesResource, name), &v1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Console), err
}

// List takes label and field selectors, and returns the list of Consoles that match those selectors.
func (c *FakeConsoles) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ConsoleList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(consolesResource, consolesKind, opts), &v1.ConsoleList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ConsoleList{ListMeta: obj.(*v1.ConsoleList).ListMeta}
	for _, item := range obj.(*v1.ConsoleList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested consoles.
func (c *FakeConsoles) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(consolesResource, opts))
}

// Create takes the representation of a console and creates it.  Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Create(ctx context.Context, console *v1.Console, opts metav1.CreateOptions) (result *v1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(consolesResource, console), &v1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Console), err
}

// Update takes the representation of a console and updates it. Returns the server's representation of the console, and an error, if there is any.
func (c *FakeConsoles) Update(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (result *v1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(consolesResource, console), &v1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Console), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeConsoles) UpdateStatus(ctx context.Context, console *v1.Console, opts metav1.UpdateOptions) (*v1.Console, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(consolesResource, "status", console), &v1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Console), err
}

// Delete takes name of the console and deletes it. Returns an error if one occurs.
func (c *FakeConsoles) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(consolesResource, name, opts), &v1.Console{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeConsoles) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(consolesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ConsoleList{})
	return err
}

// Patch applies the patch and returns the patched console.
func (c *FakeConsoles) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.Console, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(consolesResource, name, pt, data, subresources...), &v1.Console{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Console), err
}
