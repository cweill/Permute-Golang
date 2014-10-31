package permute_test

import (
	. "permute"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Permute", func() {
  Describe("NextPermutation", func() {
    It("is `` after ``", func() {
      Expect(NextPermutation("")).To(Equal(""))
    })

    It("is `0` after `0`", func() {
      Expect(NextPermutation("0")).To(Equal("0"))
    })

    It("is 8344112666 after 8342666411", func() {
      Expect(NextPermutation("8342666411")).To(Equal("8344112666"))
    })

    It("permutes 012", func() {
      permutations := []string{"012", "021", "102", "120", "201", "210"}
      for i := 0; i < len(permutations); i++ {
        if i == len(permutations)-1 {
          Expect(NextPermutation(permutations[i])).To(Equal(permutations[i]))
        } else {
          Expect(NextPermutation(permutations[i])).To(Equal(permutations[i+1]))
        }
      }
    })
  })

  Describe("LexicographicPermutations", func() {
    It("ouputs `` for ``", func() {
      Expect(LexicographicPermutations("")).To(Equal([]string{""}))
    })

    It("ouputs `0` for `0`", func() {
      permutations := []string{"0"}
      Expect(LexicographicPermutations("0")).To(Equal(permutations))
    })

    It("outputs 210 in lexicographic order", func() {
      permutations := []string{"012", "021", "102", "120", "201", "210"}
      Expect(LexicographicPermutations("210")).To(Equal(permutations))
    })
  })
})
