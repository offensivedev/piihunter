package cmd

import (
  "os/exec"
  "testing"
  "fmt"
  "strings"
  "github.com/stretchr/testify/assert"
)

var mainGo = "../main.go"
var dir = "/tmp/"
var gitRepo = "https://github.com/offensivedev/urldozer"

func checkErr(err error) {
  if err != nil {
    fmt.Println(err)
  }
}

func TestMain(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo).CombinedOutput()
  matchString := "Scans the filesystem and/or Git repository for any unencrypted Personally Identifiable Information (PII)."
  if !strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}

func TestScanCmd(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo, "scan").CombinedOutput()
  matchString := "Please supply either directory or github repository using --dir & --git flags"
  if !strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}

func TestScanCmdWithoutArgValueDir(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo, "scan", "--dir").CombinedOutput()
  matchString := "Error: flag needs an argument: --dir"
  if !strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}

func TestScanCmdWithoutArgValueGit(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo, "scan", "--git").CombinedOutput()
  matchString := "Error: flag needs an argument: --git"
  if !strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}

func TestScanCmdWrongDirectory(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo, "scan", "--dir=test").CombinedOutput()
  matchString := "Directory does not exist"
  if !strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}

func TestScanCmdInvalidRepository(t *testing.T) {
  out, err := exec.Command("go", "run", mainGo, "scan", "--git=https://github.com/404").CombinedOutput()
  matchString := "Unable to fetch the git repository"
  if strings.Contains(string(out), matchString) {
    assert.Error(t, err)
  }
}