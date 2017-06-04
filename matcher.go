package main

import (
	"os"

	"bufio"
	"fmt"
	"regexp"
	"strings"
	"sync"

	"github.com/matcher/utils"
	"github.com/urfave/cli"
)

var command = regexp.MustCompile(`^(\w+(,\w+)*)?$`)

func main() {
	var dictionary string
	reader := bufio.NewReader(os.Stdin)
	cli.AppHelpTemplate = utils.Template
	app := cli.NewApp()
	app.Version = utils.Version
	app.Name = utils.Name
	app.Usage = utils.Usage
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "dictionary",
			Usage:       "comma separated words",
			Destination: &dictionary,
		},
	}
	app.Action = func(c *cli.Context) error {
		if !command.Match([]byte(dictionary)) || len(dictionary) == 0 {
			fmt.Println("Invalid value formant for option --dictionary")
			fmt.Println("See 'matcher help'")
			return nil
		}

		fmt.Print("Dictionary:")
		fmt.Println(dictionary)
		fmt.Println("Please, enter lines of the single space separated words.")
		fmt.Println("A blank line would be considered as an end of the input.")

		r := utils.NewMatcher(dictionary)
		var wg sync.WaitGroup

		for {
			text, _ := reader.ReadString('\n')
			text = strings.Replace(text, "\n", "", -1)

			if strings.Compare("", text) == 0 {
				break
			}
			wg.Add(1)
			go r.Calculate(text, &wg)
		}

		wg.Wait()
		fmt.Println("Resulting output:")
		fmt.Println(r)

		return nil
	}

	app.Run(os.Args)
}
