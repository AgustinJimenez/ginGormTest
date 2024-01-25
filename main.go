package main

import (
	"go_practice/setup"
)

func main() {
  app := setup.GetApp()
  app.Run()
  println("RUNNING!!!")
}