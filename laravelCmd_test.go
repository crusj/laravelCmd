package main

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestRoute(t *testing.T) {
	cmd := exec.Command("go", "run", "laravelCmd.go", "route", "--path", "test_Route.php")
	cmd.Dir = "/Users/crusj/Project/laravelCmd"
	output,err := cmd.Output()
	fmt.Printf("%s\n",output)
	fmt.Println(err)
}
