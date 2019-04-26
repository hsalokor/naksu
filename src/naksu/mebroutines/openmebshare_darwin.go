package mebroutines

import (
	"fmt"

	"naksu/log"
)

// OpenMebShare opens file explorer with meb share path
func OpenMebShare() {
	mebSharePath := GetMebshareDirectory()

	log.LogDebug(fmt.Sprintf("MEB share directory: %s", mebSharePath))

	if !ExistsDir(mebSharePath) {
		ShowWarningMessage("Cannot open MEB share directory since it does not exist")
		return
	}

	runParams := []string{"open", mebSharePath}

	output, err := RunAndGetOutput(runParams, false)

	if err != nil {
		ShowWarningMessage("Could not open MEB share directory")
	}

	log.LogDebug("MEB share directory open output:")
	log.LogDebug(output)
}
