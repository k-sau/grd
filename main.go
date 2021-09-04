package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/joeguo/tldextract"
)

func main() {
	cache := "/tmp/tld.cache"
	extract, _ := tldextract.New(cache, false)

	mm := make(map[string]bool)

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		domain := sc.Text()
		result := extract.Extract(domain)
		tmp := fmt.Sprintf("%s.%s", result.Root, result.Tld)
		if !mm[tmp] {
			fmt.Println(tmp)
			mm[tmp] = true
		}
	}
}
