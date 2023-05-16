package main

import (
	"adventure/model"
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// entrance - name of the first page
var entrance string

// StoryBlocks - map associated with story blocks and it's names
var StoryBlocks = make(map[string]struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
})

type Handler struct {
}

// ServeHTTP - servers http handler
func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var block model.Block
	switch r.URL.Query().Get("block") {
	case "":
		block = StoryBlocks[entrance]
	default:
		block = StoryBlocks[r.URL.Query().Get("block")]
	}
	renderTemplate(w, block)
}

// renderTemplate - rendering specific html template
func renderTemplate(w http.ResponseWriter, Data interface{}) {
	tmpl, err := template.New("adventure.gohtml").ParseFiles("./templates/adventure.gohtml")
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, Data)
	if err != nil {
		log.Printf("Problem executing template: %v", err)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		fmt.Println(err)
	}
}

// jsonStoryToGoStruct - parse json file from story directory and converts it to go struct
func jsonStoryToGoStruct(path string, story *model.Adventure) error {
	f, _ := os.Open(path)

	scaner := bufio.NewScanner(f)
	tmp := ""
	for scaner.Scan() {
		tmp += scaner.Text()
	}

	r := strings.NewReader(tmp)

	return json.NewDecoder(r).Decode(story)
}

// CreateStoryCache - creating map associated with story block data and it's name
func CreateStoryCache(b model.Adventure) {
	val := reflect.ValueOf(b)
	for i := 0; i < val.Type().NumField(); i++ {
		if i == 0 {
			entrance = val.Type().Field(i).Tag.Get("json")
		}
		StoryBlocks[val.Type().Field(i).Tag.Get("json")] = val.Field(i).Interface().(struct {
			Title   string   `json:"title"`
			Story   []string `json:"story"`
			Options []struct {
				Text string `json:"text"`
				Arc  string `json:"arc"`
			} `json:"options"`
		})
	}
}

// TerminalStory - run terminal version of story
func TerminalStory() {
	var (
		block = entrance
		input string
	)
	for {
		fmt.Fprintln(os.Stdout, StoryBlocks[block].Title)
		for _, paragraph := range StoryBlocks[block].Story {
			fmt.Fprintln(os.Stdout, paragraph)
		}
		for i, options := range StoryBlocks[block].Options {
			fmt.Fprintf(os.Stdout, "%s %s-%d\n", options.Text, options.Arc, i+1)
		}
		fmt.Fprintln(os.Stdout, "type exit if you want to quit")
		if _, err := fmt.Fscan(os.Stdin, &input); err != nil {
			fmt.Println(err)
			continue
		}
		if input == "exit" {
			return
		}
		cnt, err := strconv.Atoi(strings.TrimSpace(input))
		if err != nil {
			fmt.Println(err)
			continue
		}
		block = StoryBlocks[block].Options[cnt-1].Arc
	}
}

func main() {
	var (
		handler Handler
		Story   model.Adventure
	)
	if err := jsonStoryToGoStruct("./story/gopher.json", &Story); err != nil {
		panic(err)
	}
	CreateStoryCache(Story)
	terminal := flag.Bool("cmd", false, "specifies if 3story runs on terminal")
	flag.Parse()
	if *terminal {
		TerminalStory()
	} else {
		_ = http.ListenAndServe("localhost:8080", handler)
	}

}
