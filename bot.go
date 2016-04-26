package main

import (
	"bufio"
	"fmt"
	"net"
	"net/textproto"
	"os"
	"regexp"
)

const HOST string = "irc.twitch.tv:6667"
const PASS string = ""
const NICK string = ""
const CHAN string = ""

func pong(conn net.Conn, msg string) {
	b, err := regexp.MatchString("PING :tmi.twitch.tv", msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		r := "PONG :tmi.twitch.tv"
		fmt.Fprintf(conn, "%s\r\n", r)
		fmt.Printf("%s\n", msg)
	}
}

func hello(conn net.Conn, msg string) {
	b, err := regexp.MatchString("hi bot", msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	if b {
		r := "Hello!"
		fmt.Fprintf(conn, "PRIVMSG %s  %s\r\n", CHAN, r)
		fmt.Printf("\t%s", r)
	}
}

func main() {
	conn, err := net.Dial("tcp", HOST)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Fprintf(conn, "PASS %s\r\n", PASS)
	fmt.Fprintf(conn, "NICK %s\r\n", NICK)
	fmt.Fprintf(conn, "JOIN %s\r\n", CHAN)
	fmt.Fprintf(conn, "CAP REQ :twitch.tv/membership\r\n")
	defer conn.Close()

	reader := bufio.NewReader(conn)
	tp := textproto.NewReader(reader)

	for {
		line, err := tp.ReadLine()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%s\n", line)

		pong(conn, line)
		hello(conn, line)
	}
}
