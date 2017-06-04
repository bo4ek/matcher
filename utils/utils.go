package utils

import (
	"bytes"
	"strconv"
	"strings"
	"sync"
)

const (
	//number of columns
	columns = 8
	// max column length
	maxLength = 9
	// space for the short values
	separator = " "
)

// This struct contains 2D slice of words count and dictionary of the words
type match struct {
	matrix     [][]int
	dictionary []string
	sync.Mutex
}

// Creates a new instance of match
func NewMatcher(args string) *match {
	r := new(match)
	// create map to avoid duplicates of the words in the dictionary
	dict := make(map[string]struct{})
	for _, v := range strings.Split(args, ",") {
		dict[v] = struct{}{}
	}
	//copy map to slice to avoid random output of results
	r.dictionary = make([]string, len(r.dictionary))
	for word := range dict {
		r.dictionary = append(r.dictionary, word)
	}

	// create 2D slice to store counts for dict words
	r.matrix = make([][]int, len(r.dictionary))
	for i := range r.matrix {
		r.matrix[i] = make([]int, columns+1)
	}
	return r
}

// Calculate function increment count of the word in the column and total count
func (r *match) Calculate(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	words := generateWordsList(line)
	dy := 0
	for _, dictRecord := range r.dictionary {
		for dx, word := range words {
			if dictRecord == word {
				r.incrementCount(dx, dy)
			}
		}
		dy++
	}
}

// increments count in the column and total count for specific word by 1
func (r *match) incrementCount(dx, dy int) {
	r.Lock()
	defer r.Unlock()
	r.matrix[dy][dx]++
	r.matrix[dy][columns]++
}

// generates slice of words from input
func generateWordsList(line string) []string {
	words := strings.Fields(line)
	if len(words) > columns {
		words = words[:columns]
	} else {
		words = append(words, make([]string, columns-len(words))...)
	}
	return words
}

// custom implementation of String() method for a fancy output
func (r *match) String() string {
	var (
		a       int
		element string
		buffer  bytes.Buffer
	)

	for n := 0; n < columns; n++ {
		element = "col " + strconv.Itoa(n+1)
		buffer.WriteString(element)
		for i := 0; i < maxLength-len(element); i++ {
			buffer.WriteString(separator)
		}
	}
	element = "total"
	buffer.WriteString(element)
	for i := 0; i < maxLength-len(element); i++ {
		buffer.WriteString(separator)
	}
	element = "word"
	buffer.WriteString(element)
	for i := 0; i < maxLength-len(element); i++ {
		buffer.WriteString(separator)
	}
	buffer.WriteString("\n")
	for _, word := range r.dictionary {
		for _, n := range r.matrix[a] {
			element = strconv.Itoa(n)
			buffer.WriteString(element)
			for i := 0; i < maxLength-len(element); i++ {
				buffer.WriteString(separator)
			}
		}
		buffer.WriteString(word)
		buffer.WriteString("\n")
		a++
	}
	return buffer.String()
}
