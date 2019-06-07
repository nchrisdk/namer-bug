package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	category := os.Args[1]

	switch category {
	case "help":
		getHelp()
	case "list":
		getList()
	case "pony":
		fmt.Println(getTwoPartName(category))
	case "mix":
		getMixedNames(os.Args[2:])
	default:
		fmt.Println(getOnePartName(category))
	}

}

func getHelp() {
	fmt.Printf("Usage %s [command] <category> [category]\n", os.Args[0])
	fmt.Println("Where command can be list or mix. List will show a list of name sources which can be used either on their own or combined using the mix command.")
	fmt.Printf("Examples:\n %s pony\n  # outputs: River-Blueberry\n %s mix adjectives pony\n  # outputs: Mundane-Lila\n", os.Args[0], os.Args[0])
}

func getList() {
	files, err := ioutil.ReadDir("./names")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Current naming sources available: ")
	for _, f := range files {
		fmt.Println(f.Name())
	}
}

func getOnePartName(category string) string {
	names := getNames(category)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	return strings.Title(names[r.Intn(len(names))])
}

func getTwoPartName(category string) string {
	names := getNames(category)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	length := len(names)
	return fmt.Sprintf("%s-%s", strings.Title(names[r.Intn(length)]), strings.Title(names[r.Intn(length)]))
}

func getMixedNames(categories []string) {
	var parts []string
	for _, cat := range categories {
		name := getOnePartName(cat)
		parts = append(parts, name)
	}
	fmt.Println(strings.Join(parts, "-"))
}

func getNames(category string) (names []string) {
	file, err := os.Open(fmt.Sprintf("./names/%s", category))
	if err != nil {
		fmt.Println("unable to read names for category: ", category)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		name := scanner.Text()
		names = append(names, name)
	}
	if scanner.Err() != nil {
		log.Fatal(err)
	}
	return names
}
