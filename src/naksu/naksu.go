package main

import (
	"fmt"
	"naksu/config"
	"naksu/mebroutines"
	"naksu/xlate"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
	"github.com/kardianos/osext"
)

const version = "1.6.0"
const lowDiskLimit = 50 * 1024 * 1024 // 50 Gb

// URLTest is testing URL for checking network connection
const URLTest = "http://static.abitti.fi/usbimg/qa/latest.txt"

var isDebug bool

// Options contains command line options
type Options struct {
	IsDebug       bool   `short:"D" long:"debug" description:"Turn debugging on" optional:"true"`
	Version       bool   `short:"v" long:"version" description:"Print naksu version" optional:"true"`
	UpdateChannel string `long:"channel" choice:"release" choice:"beta" description:"Use given release channel. Values other than release are experimental versions" optional:"true"`
	SelfUpdate    string `long:"self-update" choice:"enabled" choice:"disabled" description:"Control self-update behaviour. Naksu will always warn if your version is out-of-date" optional:"true"`
}

var options Options

func handleOptionalArgument(longName string, parser *flags.Parser, function func(option *flags.Option)) {
	opt := parser.FindOptionByLongName(longName)
	if opt != nil && opt.IsSet() {
		function(opt)
	}
}

func main() {
	// Load configuration if it exists
	config.Load()

	// Set default UI language
	xlate.SetLanguage(config.GetLanguage())

	var parser = flags.NewParser(&options, flags.Default)
	_, parseErr := parser.Parse()

	if flags.WroteHelp(parseErr) {
		os.Exit(0)
	} else if parseErr != nil {
		panic(parseErr)
	}

	handleOptionalArgument("debug", parser, func(opt *flags.Option) {
		isDebug = true
	})

	handleOptionalArgument("version", parser, func(opt *flags.Option) {
		fmt.Printf("Naksu version is %v\n", version)
		os.Exit(0)
	})

	handleOptionalArgument("channel", parser, func(opt *flags.Option) {
		config.SetReleaseChannel(opt.String())
		config.Save()
	})

	handleOptionalArgument("self-update", parser, func(opt *flags.Option) {
		config.SetReleaseChannel(opt.String())
		config.Save()
	})

	RunSelfUpdate()

	mebroutines.SetDebug(isDebug)

	// Determine/set path for debug log
	mebroutines.SetDebugFilename(mebroutines.GetNewDebugFilename())

	mebroutines.LogDebug(fmt.Sprintf("This is Naksu %s. Hello world!", version))

	// Check whether we have a terminal (restart with x-terminal-emulator, if missing)
	if !mebroutines.ExistsStdin() {
		pathToMe, err := osext.Executable()
		if err != nil {
			mebroutines.LogDebug("Failed to get executable path")
		}
		commandArgs := []string{"x-terminal-emulator", "-e", pathToMe}

		mebroutines.LogDebug(fmt.Sprintf("No stdin, restarting with terminal: %s", strings.Join(commandArgs, " ")))
		_, err = mebroutines.RunAndGetOutput(commandArgs)
		if err != nil {
			mebroutines.LogDebug(fmt.Sprintf("Failed to restart with terminal: %d", err))
		}
		mebroutines.LogDebug(fmt.Sprintf("No stdin, returned from %s", strings.Join(commandArgs, " ")))

		// Normal termination
		os.Exit(0)
	}

	var err = RunUI()

	if err != nil {
		panic(err)
	}

	config.Save()

	mebroutines.LogDebug("Exiting GUI loop")
}
