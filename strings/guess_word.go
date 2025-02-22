package main

import "fmt"

// 843. Guess the Word
// Hard
// Topics
// Companies
// You are given an array of unique strings words where words[i] is six letters long. One word of words was chosen as a secret word.

// You are also given the helper object Master. You may call Master.guess(word) where word is a six-letter-long string, and it must be from words. Master.guess(word) returns:

// -1 if word is not from words, or
// an integer representing the number of exact matches (value and position) of your guess to the secret word.
// There is a parameter allowedGuesses for each test case where allowedGuesses is the maximum number of times you can call Master.guess(word).

// For each test case, you should call Master.guess with the secret word without exceeding the maximum number of allowed guesses. You will get:

// "Either you took too many guesses, or you did not find the secret word." if you called Master.guess more than allowedGuesses times or if you did not call Master.guess with the secret word, or
// "You guessed the secret word correctly." if you called Master.guess with the secret word with the number of calls to Master.guess less than or equal to allowedGuesses.
// The test cases are generated such that you can guess the secret word with a reasonable strategy (other than using the bruteforce method).

// Example 1:

// Input: secret = "acckzz", words = ["acckzz","ccbazz","eiowzz","abcczz"], allowedGuesses = 10
// Output: You guessed the secret word correctly.
// Explanation:
// master.guess("aaaaaa") returns -1, because "aaaaaa" is not in wordlist.
// master.guess("acckzz") returns 6, because "acckzz" is secret and has all 6 matches.
// master.guess("ccbazz") returns 3, because "ccbazz" has 3 matches.
// master.guess("eiowzz") returns 2, because "eiowzz" has 2 matches.
// master.guess("abcczz") returns 4, because "abcczz" has 4 matches.
// We made 5 calls to master.guess, and one of them was the secret, so we pass the test case.
// Example 2:

// Input: secret = "hamada", words = ["hamada","khaled"], allowedGuesses = 10
// Output: You guessed the secret word correctly.
// Explanation: Since there are two words, you can guess both.

// Constraints:

// 1 <= words.length <= 100
// words[i].length == 6
// words[i] consist of lowercase English letters.
// All the strings of wordlist are unique.
// secret exists in words.
// 10 <= allowedGuesses <= 30

/**
 * // This is the Master's API interface.
 * // You should not implement it, or speculate about its implementation
 * type Master struct {
 * }
 *
 * func (this *Master) Guess(word string) int {}
 */

type Master struct {
	allowedGuesses int
	secretWord     string
	words          map[string]bool
	secretWordFreq map[string]int
}

func (this *Master) Guess(word string) int {
	this.allowedGuesses--
	matchedSecretWords := map[int]int{}
	if ok := this.words[word]; !ok {
		return -1
	}
	if this.secretWord == word {
		return 6
	}
	count := 0
	for i := 0; i < len(word); i++ {
		count1, ok := this.secretWordFreq[string(word[i])]
		if _, exist := matchedSecretWords[count1]; !exist && ok {
			count++
			matchedSecretWords[count1] = count1
		}

	}

	return count
}

func NewMaster(allowedGuesses int, secretWord string) *Master {
	return &Master{
		allowedGuesses: allowedGuesses,
		secretWord:     secretWord,
		words:          map[string]bool{},
		secretWordFreq: map[string]int{},
	}
}

func (this *Master) initializeMap(words []string) {
	for _, word := range words {
		this.words[word] = true
	}
	for i := 0; i < len(this.secretWord); i++ {
		this.secretWordFreq[string(this.secretWord[i])] = i
	}
}

func findSecretWord(words []string, master *Master) {
	master.initializeMap(words)
	found := false
	for _, word := range words {
		if master.allowedGuesses == 0 {
			fmt.Println("Either you took too many guesses, or you did not find the secret word.")
			return
		}
		result := master.Guess(word)
		if result == 6 {
			found = true
		}
	}
	if found {
		fmt.Println("You guessed the secret word correctly.")
	} else {
		fmt.Println("Either you took too many guesses, or you did not find the secret word.")
	}
}

func main() {
	findSecretWord([]string{"acckzz", "ccbazz", "eiowzz", "abcczz"}, NewMaster(10, "acckzz"))
	findSecretWord([]string{"hamada", "khaled"}, NewMaster(10, "hamada"))
	fmt.Println()
	findSecretWord([]string{"acckzze", "ccbazz", "eiowzz", "abcczz"}, NewMaster(10, "acckzz"))
	findSecretWord([]string{"hamadaw", "khaled"}, NewMaster(10, "hamada"))
}
