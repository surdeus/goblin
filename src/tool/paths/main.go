package paths

import (
	"fmt"
	"flag"
	"bufio"
	"os"
	"github.com/surdeus/goblin/src/pathx"
	"path"
	"log"
	"errors"
)

var (
	part string
	handlers = map[string] func(string) string {
		"base" : path.Base,
		"ext" : path.Ext,
		"dir" : path.Dir,
		"all" : func(v string) string {return v},
	}
	handler func(string) string
	r bool
	noPartErr = errors.New("no such part")
)



func handlePath(p string) {
	if handler != nil {
		p = handler(p)
	}

	if r {
		pth := pathx.From(p)
		fmt.Println(pth.Real())
	} else {
		fmt.Println(p)
	}
}

func Run(args []string) {
	arg0 := args[0]
	args = args[1:]
	flags := flag.NewFlagSet(arg0, flag.ExitOnError)
	flags.StringVar(&part, "p", "all", "part of path you want to print")
	flags.BoolVar(&r, "r", false, "print real OS dependent paths")

	flags.Parse(args)
	args = flags.Args()
	lhandler, ok := handlers[part]
	if !ok {
		log.Fatal(noPartErr)
	}
	handler = lhandler

	for _, p := range args {
		handlePath(p)
	}

	rd := bufio.NewScanner(os.Stdin)
	for rd.Scan() {
		handlePath(rd.Text())
	}
}

