package option

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Options struct {
	ShowHelp     bool
	TargetUrl    string
	WordlistPath string
	Data         string
	FilterCode   int
	XSS          bool
}

func (o Options) SearchInput(v string) {
	urlStr := v

	resp, err := http.Get(urlStr)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var inputs []string
	doc.Find("input").Each(func(index int, input *goquery.Selection) {
		inputName, _ := input.Attr("name")
		inputs = append(inputs, inputName)
	})

	payloads, err := readPayloadsFromFile("payloads.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, inputName := range inputs {
		for _, payload := range payloads {
			result, err := sendPayloadAndGetResult(urlStr, inputName, payload)
			if err != nil {
				log.Printf("Error sending payload: %v", err)
				continue
			}
			fmt.Printf("Input: %s, Payload: %s\nResults:\n%s\n", inputName, payload, strings.Join(result, "\n"))
		}
	}
}

func readPayloadsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var payloads []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		payloads = append(payloads, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return payloads, nil
}

func sendPayloadAndGetResult(url, inputName, payload string) ([]string, error) {
	client := &http.Client{}
	payloadData := strings.NewReader(inputName + "=" + payload)
	req, err := http.NewRequest("POST", url, payloadData)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	var results []string
	doc.Find("a").Each(func(index int, link *goquery.Selection) {
		linkHref, _ := link.Attr("href")
		results = append(results, linkHref)
	})

	return results, nil
}

func (o Options) DisplayHelp() {

	if !o.ShowHelp {
		return
	}
	usage := `  
	                                                                         
 HHHHHHHHH     HHHHHHHHHEEEEEEEEEEEEEEEEEEEEEELLLLLLLLLLL             PPPPPPPPPPPPPPPPP   
 H:::::::H     H:::::::HE::::::::::::::::::::EL:::::::::L             P::::::::::::::::P  
 H:::::::H     H:::::::HE::::::::::::::::::::EL:::::::::L             P::::::PPPPPP:::::P 
 HH::::::H     H::::::HHEE::::::EEEEEEEEE::::ELL:::::::LL             PP:::::P     P:::::P
   H:::::H     H:::::H    E:::::E       EEEEEE  L:::::L                 P::::P     P:::::P
   H:::::H     H:::::H    E:::::E               L:::::L                 P::::P     P:::::P
   H::::::HHHHH::::::H    E::::::EEEEEEEEEE     L:::::L                 P::::PPPPPP:::::P 
   H:::::::::::::::::H    E:::::::::::::::E     L:::::L                 P:::::::::::::PP  
   H:::::::::::::::::H    E:::::::::::::::E     L:::::L                 P::::PPPPPPPPP    
   H::::::HHHHH::::::H    E::::::EEEEEEEEEE     L:::::L                 P::::P            
   H:::::H     H:::::H    E:::::E               L:::::L                 P::::P            
   H:::::H     H:::::H    E:::::E       EEEEEE  L:::::L         LLLLLL  P::::P            
 HH::::::H     H::::::HHEE::::::EEEEEEEE:::::ELL:::::::LLLLLLLLL:::::LPP::::::PP          
 H:::::::H     H:::::::HE::::::::::::::::::::EL::::::::::::::::::::::LP::::::::P          
 H:::::::H     H:::::::HE::::::::::::::::::::EL::::::::::::::::::::::LP::::::::P          
 HHHHHHHHH     HHHHHHHHHEEEEEEEEEEEEEEEEEEEEEELLLLLLLLLLLLLLLLLLLLLLLLPPPPPPPPPP   
  
 -url <url> : Target url
 -xss <path> : XSS wordlist path
 -help : Show usage details
 	


 `
	fmt.Printf("%s", usage)
	os.Exit(0)
}
