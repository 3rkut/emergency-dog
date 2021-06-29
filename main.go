// Author: 3rkut

package main

////// imports. //////

import (
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/projectdiscovery/gologger"
)

///////////////////////

const banner = `
/ \__
(    @\___
/         O
/   (_____/
/_____/
`

// Banner Function.
func showBanner() {
	gologger.Print().Msgf("%s\n\n", color.MagentaString(banner))
	gologger.Print().Msg(color.RedString("[*] Welcome to emergency-dog!\n[*] Use with caution. You are responsible for your actions\n"))
}

func main() {
	showBanner()
	file, err := os.OpenFile("targets.txt", os.O_WRONLY, 0644) // Open our File.
	if err != nil {
		gologger.Fatal().Msgf("Error %s:\n", err)
	}
	defer file.Close()

	data, err := ioutil.ReadFile("targets.txt") // Read with ioutil.
	if err != nil {
		gologger.Fatal().Msgf("Error %s:\n", err)
	}
	if len(data) == 0 { // Check for empty targets.txt
		gologger.Info().Msg(color.RedString("Please enter a target per line!"))
	}
	for _, site := range strings.Split(string(data), "\n") { // Read line by line.
		if len(site) <= 1 {
			continue
		}
		resp, err := http.Get(site)
		if err != nil {
			gologger.Fatal().Msgf("Error %s:\n", err)
		}

		gologger.Info().Msgf("[*] Target: %s\n", site)
		result := "[+]" + " " + strconv.Itoa(resp.StatusCode) + " " + http.StatusText(resp.StatusCode)
		// TODO: if status code is not 200, program gets crash. FIX THIS.
		switch resp.StatusCode {
		case 200:
			gologger.Info().Msg(color.GreenString(result))
		case 300:
			gologger.Info().Msg(color.YellowString(result))
		case 301:
			gologger.Info().Msg(color.BlueString(result))
		case 302:
			gologger.Info().Msg(color.MagentaString(result))
		case 400:
			gologger.Info().Msg(color.RedString(result))
		case 403:
			gologger.Info().Msg(color.RedString(result))
		case 404:
			gologger.Info().Msg(color.MagentaString(result))
		case 500:
			gologger.Info().Msg(color.BlueString(result))
		case 502:
			gologger.Info().Msg(color.RedString(result))
		case 503:
			gologger.Info().Msg(color.YellowString(result))
		default:
			gologger.Info().Msg(color.RedString(err.Error()))
		}
	}
}
