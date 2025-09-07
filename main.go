package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/tabwriter"
)

func main() {
	argHandler()
	fmt.Println()
	info()
}

func info() {
	// func NewWriter(output io.Writer, minwidth, tabwidth, padding int, padchar byte, flags uint) *Writer
	w := tabwriter.NewWriter(os.Stdout, 40, 0, 1, ' ', 0)

	fmt.Fprintln(w, "name: ezrieh"+
		"\tlocation: NYC")

	fmt.Fprintln(w, "www: https://ezri.pet"+
		"\temail: me@ezri.pet")

	fmt.Fprintln(w, "fedi: @ezri@starry.cafe"+
		"\tmatrix: @ezri:envs.net")

	w.Flush()

	fmt.Println()
	fmt.Println("Discord Status")
	discord()

	fmt.Println()
	fmt.Println("Workstation Status")
	ide0()

	fmt.Println()
	fmt.Println("EzriCloud Status")
	ezricloud()

	fmt.Println("Have a nice day~")
}

func ezricloud() {
	resp, err := http.Get("https://api.ezri.pet/ezricloud.text")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func ide0() {
	resp, err := http.Get("https://api.ezri.pet/ide0.text")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func discord() {
	resp, err := http.Get("https://api.ezri.pet/discord.text")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))
}

func argHandler() {
	/*
		https://manpages.ubuntu.com/manpages/trusty/man8/efingerd.8.html
		$1 - identity of remote user, (null) if his/her/its system is not running ident
		$2 - address  of  remote  machine  (IP  number  if it has not reverse DNS entry or you
			 specified -n)
		$3 - name of local user being fingered
	*/

	if len(os.Args) < 3 {
		fmt.Println("Greetings!")
		return
	}

	remoteUser := os.Args[1]
	remoteMachine := os.Args[2]
	localUser := os.Args[3]

	// log to file
	path := "finger.log"
	if hostname, err := os.Hostname(); err == nil {
		if hostname == "tilde.town" {
			path = "/home/ezri/finger.log"
		}
		// envs.net
		if hostname == "localhost" {
			path = "/home/ezri/finger.log"
		}
		// tilde.team
		if hostname == "tilde" {
			path = "/home/ezri/finger.log"
		}
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("error opening file: %v", err)
	}
	defer f.Close()

	log := fmt.Sprintf("%s\t%s\t%s\n", remoteUser, remoteMachine, localUser)
	if _, err = f.WriteString(log); err != nil {
		fmt.Println("error writing to file: %v", err)
	}

	switch {
	case remoteUser == "(null)" && remoteMachine == "(null)":
		fmt.Println("Greetings!")
	// remoteUser is null but remoteMachine is not null
	case remoteUser == "(null)" && remoteMachine != "(null)":
		fmt.Println("Greetings! " + remoteMachine)
	// both remoteUser and remoteMachine are not null
	default:
		fmt.Println("Greetings! " + remoteUser + "@" + remoteMachine)
	}
}
