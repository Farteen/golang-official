package netlearning_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/Farteen/go-programming-book-learning/netlearning"
)

var _ = Describe("Net", func() {
	It("check net const", func() {
		Expect(DialFunc()).To(MatchRegexp("OK"))
	})
})
