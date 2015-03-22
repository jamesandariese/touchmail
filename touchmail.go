package main

import (
	"flag"
	"fmt"
	"net/mail"
	"os"
	"time"
)

var verboseFlag bool
var silentFlag bool

func getDateForFileName(fn string) (t time.Time, err error) {
	f, err := os.Open(fn)
	if err != nil {
		return
	}
	defer f.Close()
	msg, err := mail.ReadMessage(f)
	if err != nil {
		return
	}
	t, err = msg.Header.Date()
	return
}
func touchMailFile(fn string) (err error) {
	date, err := getDateForFileName(fn)
	if err != nil {
		return
	}
	if verboseFlag {
		fmt.Printf("%v: %v\n", fn, date)
	}
	err = os.Chtimes(fn, time.Now(), date)
	return
}

func init() {
	flag.BoolVar(&verboseFlag, "verbose", false, "Verbose or not verbose (default quiet)")
	flag.BoolVar(&verboseFlag, "v", false, "-verbose shorthand")
	flag.BoolVar(&silentFlag, "silent", false, "Silent -- don't show errors, does not cancel verbose")
	flag.BoolVar(&silentFlag, "s", false, "-silent shorthand")
}

func main() {
	flag.Parse()
	for _, x := range flag.Args() {

		if err := touchMailFile(x); err != nil && !silentFlag {
			fmt.Println(err)
		}
	}
}
