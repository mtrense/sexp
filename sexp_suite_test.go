package sexp_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSexp(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Sexp Suite")
}
