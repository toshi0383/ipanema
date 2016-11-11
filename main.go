package main

import (
    "os"
    "flag"
    "fmt"
    "runtime"
    "os/exec"
    "log"
)

// Command line flags
var (
    embeddedMobileprovision bool
    showVersion             bool

    version = "devel" // for -v flag, updated during the release process with -ldflags=-X=main.version=...
)

var ipaPath string

func init() {
    flag.BoolVar(&embeddedMobileprovision, "E", false, "Find, Decrypt and Print all embedded.mobileprovision files.")
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

    if embeddedMobileprovision {
        printEmbeddedMobileprovision()
    }
}

func verifyOptions() bool {
    return embeddedMobileprovision
}

func printEmbeddedMobileprovision() {
    bytes, err := exec.Command("./scripts/inspect_ipa_mobileprovision.sh", ipaPath).Output()
    if err != nil {
        log.Fatal(err)
        os.Exit(2)
    }
    for i := range bytes {
        fmt.Printf("%c", bytes[i])
    }
}
