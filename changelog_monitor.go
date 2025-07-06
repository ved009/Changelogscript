package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	file := flag.String("file", "", "file containing URLs to check")
	interval := flag.Duration("interval", time.Hour, "time between checks")
	flag.Parse()

	urls := []string{}

	if *file != "" {
		f, err := os.Open(*file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "unable to open file", err)
			return
		}
		defer f.Close()
		urls = append(urls, readLines(f)...)
	} else if flag.NArg() > 0 {
		urls = flag.Args()
	} else {
		urls = append(urls, readLines(os.Stdin)...)
	}

	if len(urls) == 0 {
		fmt.Fprintln(os.Stderr, "no URLs provided")
		return
	}

	for {
		today := time.Now()
		patterns := datePatterns(today)
		for _, u := range urls {
			if checkURL(u, patterns) {
				fmt.Println(u)
			}
		}
		time.Sleep(*interval)
	}
}

func readLines(r io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func datePatterns(t time.Time) []string {
	return []string{
		t.Format("January 2"),
		t.Format("Jan 2"),
		t.Format("2 January"),
		t.Format("2 Jan"),
	}
}

func checkURL(u string, patterns []string) bool {
	resp, err := http.Get(u)
	if err != nil {
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return false
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false
	}
	text := strings.ToLower(string(body))
	for _, p := range patterns {
		if strings.Contains(text, strings.ToLower(p)) {
			return true
		}
	}
	return false
}
