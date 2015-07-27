package main

import (
	// "antscript.com/utils/log"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {

	// log.LogInfo("domain start !")

	pDomain := flag.String("domain", "", "")
	pHost := (flag.String("host", "127.0.0.1", "host ?"))
	pAdd := (flag.Bool("add", true, "add ?"))

	flag.Parse()

	domain := strings.TrimSpace(*pDomain)
	host := strings.TrimSpace(*pHost)
	add := *pAdd

	if domain == "" {
		fmt.Print("Missing domain !")
		// log.LogError("domain missing !")
		os.Exit(0)
	}

	// fmt.Println(host, add)

	fileName := "/etc/hosts"
	fileOut, err := os.OpenFile(fileName, os.O_RDWR, 0644)

	if err != nil {
		fmt.Print("open /etc/hosts failed !\n\nsee readme !")
		os.Exit(0)
	}

	fileReader := bufio.NewReader(fileOut)

	fileStr, _ := fileReader.ReadString('~')

	if domain == "localhost" {
		fmt.Print(fileStr)
		os.Exit(0)
	}

	offset := strings.Index(fileStr, host+" "+domain)

	res := ""
	if add {
		if offset > -1 {
			fileStr = strings.Replace(fileStr, "# "+host+" "+domain, host+" "+domain, -1)
			res = host + " " + domain + "\n\nAdd to hosts !"
		} else {
			fileStr += ("\n" + host + " " + domain)
			res = host + " " + domain + "\n\nAdd to hosts !"
		}
	} else {
		if offset > -1 {
			if strings.Index(fileStr, "# "+host+" "+domain) == -1 {
				fileStr = strings.Replace(fileStr, host+" "+domain, "# "+host+" "+domain, -1)
			}
		}
		res = host + " " + domain + "\n\nRemove from hosts !"
	}

	fileOut.Close()

	fileOut, _ = os.OpenFile(fileName, os.O_RDWR|os.O_TRUNC, 0644)
	defer fileOut.Close()
	_, err = fileOut.WriteString(fileStr)
	if err != nil {
		fmt.Print("write /etc/hosts failed !\n\nsee readme !")
		os.Exit(0)
	}

	fmt.Print(res)
}
