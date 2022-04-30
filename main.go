package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var regex = regexp.MustCompile("^-v+$")

func main() {
	help := flag.NewFlagSet("help", flag.ExitOnError)
	var verbose int
	help.IntVar(&verbose, "verbose", 0, "Verbosity")

	escalation := flag.NewFlagSet("escalation", flag.ExitOnError)

	favoritism := flag.NewFlagSet("favoritism", flag.ExitOnError)
	var index int
	favoritism.IntVar(&index, "n", 0, "output at most COUNT lines")

	flag.Usage = usage

	switch os.Args[1] {
	case "help":
		v, others := verbosity(os.Args[2:])
		help.Parse(others)
		help.Set("verbose", strconv.Itoa(v))
	case "escalation":
		escalation.Parse(os.Args[2:])
	case "favoritism":
		favoritism.Parse(os.Args[2:])
	default:
		fmt.Printf("%q is not a valid subcommand.\n", os.Args[1])
		flag.Usage()
		os.Exit(2)
	}

	if help.Parsed() {
		switch verbose {
		case 0:
			flag.Usage()
		case 1:
			flag.Usage()
			fmt.Println("There are no Easter Eggs in this program.")
		case 2:
			fmt.Println("There really are no Easter Eggs in this program.")
		case 3:
			fmt.Println("Didn't I already tell you that there are no Easter Eggs in this program?")
		case 4:
			fmt.Println("Okay, okay, if I give you an Easter Egg, will you go away?")
		default:
			fmt.Println("All right, you win.")
			fmt.Println(aa)
		}
	}

	if escalation.Parsed() {
		if args := escalation.Args(); len(args) != 0 {
			escalationSelect(args)
		}
	}

	if favoritism.Parsed() {
		if args := favoritism.Args(); len(args) != 0 {
			favoritismSelect(args, index)
		}
	}

}

func verbosity(xs []string) (int, []string) {
	var v float64
	others := make([]string, 0, len(xs))
	for _, x := range xs {
		if regex.MatchString(x) {
			v = math.Max(v, float64(strings.Count(x, "v")))
		} else {
			others = append(others, x)
		}
	}
	return int(v), others
}

func usage() {
	fmt.Println(`
Usage: kaiji [OPTION]... [ARG]
Write a RANDOM selection of the input lines to standard output.

Following subcommands are available.
  help:        display this help and exit
  escalation:  choose one from given arguments
  favoritism:
    -n=COUNT   output at most COUNT lines
	`)
}

func escalationSelect(args []string) {
	cheat := make([]string, 0)
	cheatlen := len(args) * (len(args) + 1) / 2
	for k, v := range args {
		for i := 0; i <= k; i++ {
			cheat = append(cheat, v)
		}
	}
	rand.Seed(time.Now().UnixNano())
	fmt.Println(cheat[rand.Intn(cheatlen)])
}

func favoritismSelect(args []string, count int) {
	if count >= len(args) {
		fmt.Printf("%d is out of range.\n", count)
		os.Exit(2)
	}
	cheat := append(args[:count-1], args[count:]...)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		target := rand.Intn(len(cheat))
		fmt.Println(cheat[target])
		cheat = append(cheat[:target], cheat[target+1:]...)
	}
}

var aa = `
ざわ‥ ざわ‥
　　＿＿　　＿
　 ＞　＼／ /"￣＼
　∠　　　　　 ＜￣
　/　 ／/ 人＼　＞
 /　 /メ＜　ﾚ ヘヽ
｜　ｲ∠二ｿ <∠_| ヽ|
｜/||＼・/〈･／|
｜L|| u￣　 ＼ |
｜ﾋ||　　 ^ _ >|
/　 |、(三三ヲ ﾊ
／| | ＼u ―　/｜
　| |＼ ＼　 /|｜
　レ|　ヽ ＼/ |｜
　  |　 |ニﾆ| |/＼
	`
