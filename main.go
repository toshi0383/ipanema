package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/progrium/go-basher"
)

func assert(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Command line flags
var (
	embeddedMobileprovision bool
	infoPlist               bool
	appInfoPlist            bool
	all                     bool
	showVersion             bool

	version = "0.3.0"
)

var ipaPath string

func init() {
	flag.BoolVar(&embeddedMobileprovision, "E", false, "Find, Decrypt and Print all embedded.mobileprovision files.")
	flag.BoolVar(&infoPlist, "I", false, "Find and Print all Info.plist files.")
	flag.BoolVar(&appInfoPlist, "i", false, "Find and Print only an app's main Info.plist file.")
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
	if all || infoPlist || appInfoPlist {
		printInfoPlist()
	}
}

func verifyOptions() bool {
	return all || (embeddedMobileprovision || (infoPlist || appInfoPlist))
}

func _execute(scriptFilePath string, appOnly bool) {
	bash, _ := basher.NewContext("/bin/bash", false)
	bash.Source(scriptFilePath, Asset)
	if appOnly {
		status, err := bash.Run("main", []string{ipaPath})
		assert(err)
		os.Exit(status)
	} else {
		status, err := bash.Run("main", []string{ipaPath, "true"})
		assert(err)
		os.Exit(status)
	}
}

func printEmbeddedMobileprovision() {
	_execute("bash/inspect_ipa_mobileprovision.sh", false)
}
func printInfoPlist() {
	if appInfoPlist {
		_execute("bash/inspect_ipa_infoplist.sh", true)
	} else {
		_execute("bash/inspect_ipa_infoplist.sh", false)
	}
}
