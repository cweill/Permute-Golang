package permute_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestPermute(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Permute Suite")
}
