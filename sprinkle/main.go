package main

import (
	"math/rand"
	"time"
	"os"
	"bufio"
	"fmt"
	"strings"
)

const otherWord = "*"
var transforms = []string{
	otherWord,
	otherWord + "app",
	otherWord + "site",
	otherWord + "time",
	otherWord + "hq",
	"get" + otherWord,
	"go" + otherWord,
	"lets" + otherWord,
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		t := transforms[rand.Intn(len(transforms))]
		fmt.Println(strings.Replace(t, otherWord, s.Text(), -1))

	}
}