package utils

var (
	Template = `NAME:
		{{.Name}}{{if .Usage}}

		{{.Usage}}{{end}}

		USAGE:
		   {{if .UsageText}}{{.UsageText}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[global options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}{{end}}{{if .Version}}{{if not .HideVersion}}

		VERSION:
		   {{.Version}}{{end}}{{end}}{{if .Description}}

		DESCRIPTION:
		   {{.Description}}{{end}}{{if len .Authors}}

		AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
		   {{range $index, $author := .Authors}}{{if $index}}
		   {{end}}{{$author}}{{end}}{{end}}{{if .VisibleCommands}}

		COMMANDS:{{range .VisibleCategories}}{{if .Name}}
		   {{.Name}}:{{end}}{{range .VisibleCommands}}
		     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}

		GLOBAL OPTIONS:
		   {{range $index, $option := .VisibleFlags}}{{if $index}}
		   {{end}}{{$option}}{{end}}{{end}}{{if .Copyright}}

		COPYRIGHT:
		   {{.Copyright}}{{end}}
		`
	Name  = "Dictionary Matcher"
	Usage = `Takes a series of lines consisting of single separator separated words
	  and compare them with the dictionary.

	  Example of Usage:

	  $ matcher --dictionary foo,bar,baz
	  $ --------
	  $ Please, enter lines of single separator separated words. A blank line would be considered as an end of input.
	  $ --------
	  $ test foo bar baz
	  $ foo bar baz
	  $ foo bar bla bla bla baz
	  $ bar foo foo bla bla baz`
	Version = "0.0.1"
)
