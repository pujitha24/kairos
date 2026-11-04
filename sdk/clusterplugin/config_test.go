package clusterplugin

import (
	"encoding/json"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Role JSON encoding", func() {
	It("encodes a Cluster through a pointer as a plain string", func() {
		out, err := json.Marshal(&Cluster{Role: RoleWorker})
		Expect(err).ToNot(HaveOccurred())
		Expect(string(out)).To(Equal(`{"role":"worker"}`))
	})

	It("decodes the role without keeping the quotes", func() {
		var c Cluster
		Expect(json.Unmarshal([]byte(`{"role":"worker"}`), &c)).To(Succeed())
		Expect(c.Role).To(Equal(Role(RoleWorker)))
	})
})
