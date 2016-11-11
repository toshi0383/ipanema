package main

import (
    "os"
    "bytes"
    "flag"
    "fmt"
    "runtime"
    "os/exec"
    "log"
)

// Command line flags
var (
    embeddedMobileprovision bool
    infoPlist               bool
    all                     bool
    showVersion             bool

    version = "0.1.0" // for -v flag, updated during the release process with -ldflags=-X=main.version=...
)

var ipaPath string

func init() {
    flag.BoolVar(&embeddedMobileprovision, "E", false, "Find, Decrypt and Print all embedded.mobileprovision files.")
    flag.BoolVar(&infoPlist, "I", false, "Find and Print all Info.plist files.")
    flag.BoolVar(&all, "A", false, "Do both what -E and -I would do.")
    flag.BoolVar(&showVersion, "v", false, "print version number")
    flag.Usage = usage
}

func usage() {
    fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] path-to.ipa\n\n", os.Args[0])
    fmt.Fprintf(os.Stderr, "OPTIONS:\n")
    flag.PrintDefaults()
}

func main() {
    flag.Parse()

    if showVersion {
        fmt.Printf("%s %s (runtime: %s)\n", os.Args[0], version, runtime.Version())
        os.Exit(0)
    }

    args := flag.Args()
    if len(args) != 1 || !verifyOptions() {
        flag.Usage()
        os.Exit(2)
    }
    ipaPath = args[0]

    if all || embeddedMobileprovision {
        printEmbeddedMobileprovision()
    }
    if all || infoPlist {
        printInfoPlist()
    }
}

func verifyOptions() bool {
   return all || (embeddedMobileprovision || infoPlist)
}

func execAndPrintStdout(command string) {
    cmd := exec.Command(command, ipaPath)
	var stdout bytes.Buffer
	cmd.Stdout = &stdout

    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }
	fmt.Println(stdout.String())
}

func printEmbeddedMobileprovision() {
	execAndPrintStdout("./scripts/inspect_ipa_mobileprovision.sh")
}
func printInfoPlist (){
	execAndPrintStdout("./scripts/inspect_ipa_infoplist.sh")
}
