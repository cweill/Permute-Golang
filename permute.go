package permute

import (
  "sort"
  "strings"
)

// Reverses a string
func Reverse(s string) string {
  runes := []rune(s)
  for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
    runes[i], runes[j] = runes[j], runes[i]
  }
  return string(runes)
}

// Splits string into head and decreasing tail substrings
func splitHeadAndTail(str string) (string, string) {
  currentIndex := len(str) - 1
  currentChar := ""
  for currentIndex >= 0 {
    if str[currentIndex:currentIndex+1] < currentChar {
      currentIndex++
      break
    }
    currentChar = str[currentIndex : currentIndex+1]
    currentIndex--
  }

  if currentIndex < 0 {
    return "", str
  }
  return str[0:currentIndex], str[currentIndex:len(str)]
}

// Sorts characters of a string
func sortString(w string) string {
    s := strings.Split(w, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

// Get next lexicographic permutation in O(n) time where n is the length of the string
// http://wordaligned.org/articles/next-permutation
func NextPermutation(str string) string {
  if len(str) == 0 || len(str) == 1 {
    return str
  }
  head, tail := splitHeadAndTail(str)
  if head == "" {
    return str
  }
  headChars := []rune(head)
  tailChars := []rune(tail)
  currentIndex := len(tail) - 1
  for currentIndex >= 0 {
    if tail[currentIndex:currentIndex+1] > head[len(head)-1:len(head)] {
      temp := tailChars[currentIndex]
      tailChars[currentIndex] = headChars[len(head)-1]
      headChars[len(head)-1] = temp
      break
    }
    currentIndex--
  }
  return string(headChars) + Reverse(string(tailChars))
}

// Compute all lexicographic permutations of a string
func LexicographicPermutations(str string) []string {
  str = sortString(str)
  result := []string{str}
  last := str
  current := NextPermutation(str)
  for last != current {
    result = append(result, current)
    last = current
    current = NextPermutation(current)
  }
  return result
}
