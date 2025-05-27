//go:build mage
// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
	"github.com/magefile/mage/sh"
)

// Default target to run when none is specified
// If not set, running mage will list available targets
// var Default = Build

func Init() error {
	mg.Deps(CheckUv)
	fmt.Println("setting up go env...")
	if err := sh.Run("go", "mod", "download"); err != nil {
		return nil
	}
	fmt.Println("setting up python env...")

	lspgendir, err := lspgeneratorDir()

	if err != nil {
		return err
	}

	return sh.Run("uv", "sync", "--directory", lspgendir)
}

func Generate() error {
	mg.Deps(Init)
	lspgendir, err := lspgeneratorDir()

	if err != nil {
		return err
	}
	protocol, err := protocolDir()

	if err != nil {
		return err
	}

	fmt.Println("generating python files...")
	cmd := exec.Command("uv", "run", "--directory", lspgendir, "python", "-m", "generator", "--plugin", "go", "--output-dir", protocol)
	return cmd.Run()
}

func Test() error {
	mg.Deps(Generate)
	lspgendir, err := lspgeneratorDir()

	if err != nil {
		return err
	}

	protocol, err := protocolDir()

	if err != nil {
		return err
	}

	fmt.Println("generating test cases...")

	if err := sh.RunV("uv", "run", "--directory", lspgendir, "python", "-m", "generator", "--plugin", "testdata", "--output-dir", protocol+"/testdata"); err != nil {
		return err
	}

	fmt.Println("running tests...")

	return sh.RunV("go", "test", "-v", protocol)
}

func CheckUv() error {
	return sh.Run("uv", "--version")
}

func lspgeneratorDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	lspgendir := filepath.Join(wd, "lspgenerator")

	return lspgendir, nil
}

func protocolDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	lspgendir := filepath.Join(wd, "protocol")

	return lspgendir, nil
}
