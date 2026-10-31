package challenger

import (
	"context"

	"github.com/go-logr/logr"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/client-go/kubernetes/fake"
)

var _ = Describe("partition label on TOFU Secrets", func() {
	DescribeTable("the label value is one the API server accepts",
		func(label string) {
			Expect(validation.IsValidLabelValue(partitionLabelValue(label))).To(BeEmpty())
		},
		Entry("default label", "COS_PERSISTENT_LUKS"),
		Entry("plus sign", "data+1_LUKS"),
		Entry("leading underscore", "_data_LUKS"),
		Entry("colon", "db:1_LUKS"),
		Entry("at sign", "data@x_LUKS"),
		Entry("leading dot", ".data_LUKS"),
		Entry("too long", "this-partition-label-is-much-longer-than-the-sixty-three-character-limit"),
	)

	It("keeps a valid label unchanged", func() {
		Expect(partitionLabelValue("COS_PERSISTENT_LUKS")).To(Equal("COS_PERSISTENT_LUKS"))
	})

	It("records the raw label in an annotation", func() {
		kclient := fake.NewSimpleClientset()
		_, err := createOrReuseTOFUSecret(kclient, "default", "s", "passphrase", "pw", "abcd", "data+1_LUKS", logr.Discard())
		Expect(err).ToNot(HaveOccurred())
		s, err := kclient.CoreV1().Secrets("default").Get(context.TODO(), "s", metav1.GetOptions{})
		Expect(err).ToNot(HaveOccurred())
		Expect(s.Annotations).To(HaveKeyWithValue("kcrypt.kairos.io/partition", "data+1_LUKS"))
		Expect(validation.IsValidLabelValue(s.Labels["kcrypt.kairos.io/partition"])).To(BeEmpty())
	})
})
