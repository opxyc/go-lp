package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"text/tabwriter"

	"github.com/opxyc/link"
)

func main() {
	url := flag.String("url", "", "url from which links are to be parsed")
	v := flag.Bool("v", false, "print html content wrapped by link")
	flag.Parse()
	if *url == "" {
		flag.Usage()
		return
	}

	res, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}

	links, err := link.Parse(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
	for _, link := range links {
		if *v {
			fmt.Fprintln(w, link.Href, "\t", link.Text)
		} else {
			fmt.Fprintln(w, link.Href)
		}
	}
	w.Flush()
}
