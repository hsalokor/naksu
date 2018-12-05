// required by selfupdate (needs context)
// +build go1.7

package main

import (
	"fmt"
	"naksu/config"
	"naksu/mebroutines"
	"naksu/xlate"
	"os"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var isOutOfDate bool

// RunSelfUpdate executes self-update
func RunSelfUpdate() {
	// Run auto-update
	if config.GetReleaseChannel() == "release" {
		if doReleaseSelfUpdate() {
			mebroutines.ShowWarningMessage("naksu has been automatically updated. Please restart naksu.")
			os.Exit(0)
		}
	} else {
		if doChannelSelfUpdate(config.GetReleaseChannel()) {
			mebroutines.ShowWarningMessage("naksu has been automatically updated. Please restart naksu.")
			os.Exit(0)
		}
	}
	if WarnUserAboutStaleVersionIfUpdateDisabled() {
		mebroutines.ShowWarningMessage("naksu has update available, but your version of naksu has updates disabled. please update or ask your administrator to update naksu")
	}
}

func doChannelSelfUpdate(channel string) bool {
	v := semver.MustParse(version)

	if mebroutines.IsDebug() {
		selfupdate.EnableLog()
	}

	latest, err := selfupdate.UpdateSelf(v, "digabi/naksu")
	if err != nil {
		mebroutines.ShowWarningMessage(fmt.Sprintf(xlate.Get("Naksu update failed. Maybe you don't have network connection?\n\nError: %s"), err))
		return false
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		mebroutines.LogDebug(fmt.Sprintf("Current binary is the latest version: %s", version))
		return false
	}
	mebroutines.LogDebug(fmt.Sprintf("Successfully updated to version: %s", latest.Version))
	return true
	//log.Println("Release note:\n", latest.ReleaseNotes)
}

func doReleaseSelfUpdate() bool {
	v := semver.MustParse(version)

	if mebroutines.IsDebug() {
		selfupdate.EnableLog()
	}

	// If self-update is disabled, do a version check nevertheless and store information for user warning
	if config.IsSelfUpdateDisabled() {
		latest, found, err := selfupdate.DetectLatest("digabi/naksu")
		if err != nil {
			mebroutines.LogDebug(fmt.Sprintf("Version check failed: %s", err))
			return false
		}
		if found && latest.Version.GT(v) {
			isOutOfDate = true
		}
		return false
	}

	latest, err := selfupdate.UpdateSelf(v, "digabi/naksu")
	if err != nil {
		mebroutines.ShowWarningMessage(fmt.Sprintf(xlate.Get("Naksu update failed. Maybe you don't have network connection?\n\nError: %s"), err))
		return false
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		mebroutines.LogDebug(fmt.Sprintf("Current binary is the latest version: %s", version))
		return false
	}
	mebroutines.LogDebug(fmt.Sprintf("Successfully updated to version: %s", latest.Version))
	return true
	//log.Println("Release note:\n", latest.ReleaseNotes)
}

// WarnUserAboutStaleVersionIfUpdateDisabled tells us if we should warn user that they are running old version if self-update is disabled. This is very corner-case check
// for environments that prefer distributing naksu via AD or other centralized management environment
func WarnUserAboutStaleVersionIfUpdateDisabled() bool {
	return config.IsSelfUpdateDisabled() && isOutOfDate
}
