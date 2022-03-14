package main

import (
	"flag"
	"fmt"
	"github.com/goccy/go-yaml"
	"io"
	"log"
	"os"
)

type Env struct {
	Name  string `yaml:"name"`
	Value string `yaml:"value"`
}

func main() {
	var export bool
	flag.BoolVar(&export, "e", false, "export for shell")
	flag.Parse()

	args := flag.Args()
	var r io.Reader
	if len(args) < 1 {
		r = os.Stdin
	} else {
		f, err := os.OpenFile(args[0], os.O_RDONLY, 0)
		if err != nil {
			log.Fatal(err)
		}
		r = f
	}

	var envs []*Env
	if err := yaml.NewDecoder(r).Decode(&envs); err != nil {
		log.Fatal(err)
	}

	if export {
		fmt.Printf("export ")
		for _, env := range envs {
			fmt.Printf("%v=\"%v\" ", env.Name, env.Value)
		}
	} else {
		for i, env := range envs {
			fmt.Printf("%v=%v", env.Name, env.Value)
			if i != len(envs)-1 {
				fmt.Printf(",")
			}
		}
	}
}
