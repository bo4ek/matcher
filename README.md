# Matcher

This is command line program which takes a dictionary of words and reads
input from stdin. This input is a series of lines consisting of single space
separated words. These words are compared with the dictionary. At the end a
statistic is printed out with the total number of matches for each
dictionary word and how many matches per "column" (defined by single spaces)
for the first 8 columns.

Example of Usage:

      $ glide install
      $ go build matcher.go
	  $ matcher --dictionary foo,bar,baz
	  $ Dictionary:foo,bar,baz
      $ Please, enter lines of the single space separated words.
      $ A blank line would be considered as an end of the input.
	  $ test foo bar baz
	  $ foo bar baz
	  $ foo bar bla bla bla baz
	  $ bar foo foo bla bla baz