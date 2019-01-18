package v1alpha1

import (
	"testing"

	"github.com/onsi/gomega"
	"golang.org/x/net/context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
)

func TestStorageNodeGroup(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "default",
	}
	created := &NodeGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
		Spec: NodeGroupSpec{
			Name: "node-foo",
		}}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &NodeGroup{}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))

	// Test Updating the Labels
	updated := fetched.DeepCopy()
	updated.Labels = map[string]string{"hello": "world"}
	g.Expect(c.Update(context.TODO(), updated)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(updated))

	// Test Delete
	g.Expect(c.Delete(context.TODO(), fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(c.Get(context.TODO(), key, fetched)).To(gomega.HaveOccurred())
}

func TestStorageNodeGroupWithIAMPolicies(t *testing.T) {
	key := types.NamespacedName{
		Name:      "foo",
		Namespace: "withpolicy",
	}
	created := &NodeGroup{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "withpolicy",
		},
		Spec: NodeGroupSpec{
			Name:        "node-foo",
			IAMPolicies: []Policy{},
		}}
	g := gomega.NewGomegaWithT(t)

	// Test Create
	fetched := &NodeGroup{
		Spec: NodeGroupSpec{
			Name:        "test",
			IAMPolicies: []Policy{},
		},
	}
	g.Expect(c.Create(context.TODO(), created)).NotTo(gomega.HaveOccurred())

	g.Expect(c.Get(context.TODO(), key, fetched)).NotTo(gomega.HaveOccurred())
	g.Expect(fetched).To(gomega.Equal(created))
}
