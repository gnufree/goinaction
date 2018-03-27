package main

import (
	"log"
	"os"
	_ "github.com/gnufree/goinaction/code/chapter2/sample/matchers"
	"github.com/gnufree/goinaction/code/chapter2/sample/search"
)

func init()  {
	log.SetOutput(os.Stdout)
}

func main()  {
	search.Run("president")
}