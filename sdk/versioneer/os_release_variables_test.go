package versioneer_test

import (
	"sort"
	"strings"

	"github.com/kairos-io/kairos/v4/sdk/versioneer"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("OSReleaseVariables", func() {
	It("returns the variables sorted and in the same order on every call", func() {
		artifact := versioneer.Artifact{
			Flavor:                "opensuse",
			FlavorRelease:         "leap-15.5",
			Variant:               "standard",
			Model:                 "generic",
			Arch:                  "amd64",
			Version:               "v2.4.2",
			SoftwareVersion:       "v1.26.9+k3s1",
			SoftwareVersionPrefix: "k3s",
		}

		first, err := artifact.OSReleaseVariables("quay.io/kairos", "kairos-io/kairos", "https://bug", "https://home")
		Expect(err).ToNot(HaveOccurred())

		lines := strings.Split(strings.TrimSuffix(first, "\n"), "\n")
		Expect(sort.StringsAreSorted(lines)).To(BeTrue(), first)

		for i := 0; i < 50; i++ {
			again, err := artifact.OSReleaseVariables("quay.io/kairos", "kairos-io/kairos", "https://bug", "https://home")
			Expect(err).ToNot(HaveOccurred())
			Expect(again).To(Equal(first))
		}
	})
})
