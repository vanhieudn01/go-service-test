package main

import "github.com/vanhieudn01/service-test/cmd"

var revision string

func main() {
	cmd.SetRevision(revision)
	cmd.Execute()
}
