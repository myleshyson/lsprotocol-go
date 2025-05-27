//go:build mage
// +build mage

package main

import (
	"fmt"
	"os/exec"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Init() error {
	mg.Deps(CheckUv)
	fmt.Println("setting up python env...")
	cmd := exec.Command("cd", "lspgenerator", "&&", "uv", "sync")
	return cmd.Run()
}

func Generate() error {
	mg.Deps(Init)
	fmt.Println("generating python files...")
	cmd := exec.Command("cd", "lspgenerator", "&&", "uv", "run", "python", "-m", "generator", "--plugin", "go", "--output-dir", "../")
	return cmd.Run()
}

func Test() error {
	mg.Deps(Generate)
	fmt.Println("generating test cases...")
	if err := sh.RunV("cd", "lspgenerator", "&&", "uv", "run", "python", "-m", "generator", "--plugin", "testdata", "--output-dir", "../"); err != nil {
		return err
	}

	fmt.Println("running tests...")
	return sh.RunV("go", "test", "-v", "./protocol")
}

func CheckUv() error {
	return sh.Run("uv", "--version")
}
