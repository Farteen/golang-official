package netlearning_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "net"
	. "github.com/Farteen/go-programming-book-learning/netlearning"
)

var _ = Describe("Netlearning", func() {
	It("check Dial", func() {
		Expect(DialFunc()).To(MatchRegexp("200"))
	})

	It("check IPv4 len", func() {
		Expect(IPv4len).To(Equal(4))
	})
})
