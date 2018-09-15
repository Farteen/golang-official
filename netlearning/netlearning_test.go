package netlearning_test

import (
	// "fmt"
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

	It("check IPv6 len", func() {
		Expect(IPv6len).To(Equal(16))
	})

	It("join host and port", func() {
		result := JoinHostPort("github.com", "443")
		Expect(result).To(Equal("github.com:443"))
	})

	It("look up addr", func() {
		names, _ := LookupAddr("192.30.253.112")
		Expect(len(names)).NotTo(Equal(0))
	})
})
