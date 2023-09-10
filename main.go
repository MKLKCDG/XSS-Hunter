package main

import (
	"flag"
	"fmt"

	"src/option"
)

func main() {
	var options option.Options

	flag.BoolVar(&options.ShowHelp, "help", false, "shows usage details")
	flag.StringVar(&options.TargetUrl, "url", "", "target url")
	flag.BoolVar(&options.XSS, "xss", false, "xss wordlist path")
	flag.Parse()

	options.DisplayHelp()

	usage := `

		  
╭━╮╭━┳━━━┳━━━╮╭╮╱╭┳╮╱╭┳━╮╱╭┳━━━━┳━━━┳━━━╮
╰╮╰╯╭┫╭━╮┃╭━╮┃┃┃╱┃┃┃╱┃┃┃╰╮┃┃╭╮╭╮┃╭━━┫╭━╮┃
╱╰╮╭╯┃╰━━┫╰━━╮┃╰━╯┃┃╱┃┃╭╮╰╯┣╯┃┃╰┫╰━━┫╰━╯┃
╱╭╯╰╮╰━━╮┣━━╮┃┃╭━╮┃┃╱┃┃┃╰╮┃┃╱┃┃╱┃╭━━┫╭╮╭╯
╭╯╭╮╰┫╰━╯┃╰━╯┃┃┃╱┃┃╰━╯┃┃╱┃┃┃╱┃┃╱┃╰━━┫┃┃╰╮
╰━╯╰━┻━━━┻━━━╯╰╯╱╰┻━━━┻╯╱╰━╯╱╰╯╱╰━━━┻╯╰━╯
  
   `
	go fmt.Printf("%s", usage)
	options.SearchInput(options.TargetUrl)

}
