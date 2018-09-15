package netlearning_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestNetlearning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Netlearning Suite")
}
