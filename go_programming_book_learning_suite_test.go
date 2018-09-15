package go_programming_book_learning_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGoProgrammingBookLearning(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GoProgrammingBookLearning Suite")
}
