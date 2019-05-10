package boxversion

// Package boxversion can be used to get version information from Vagrantfile
// (either ~/ktp/Vagrantfile or the one available in the cloud)

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"time"

	"naksu/constants"
	"naksu/log"
	"naksu/mebroutines"
	"naksu/network"
	"naksu/xlate"
)

// Cache for GetVagrantBoxAvailVersionDetails
type lastBoxAvail struct {
	boxString     string
	boxVersion    string
	boxTimestamp  int64
	updateStarted int64
}

// Global cache for GetVagrantBoxAvailVersionDetails()
var vagrantBoxAvailVersionDetailsCache lastBoxAvail

// GetVagrantFileVersion returns a human-readable localised version string
// for a given Vagrantfile (with "" defaults to ~/ktp/Vagrantfile)
//
// If you want to get the version of the current box consider using
// box.GetVersion() instead
func GetVagrantFileVersion(vagrantFilePath string) string {
	if vagrantFilePath == "" {
		vagrantFilePath = filepath.Join(mebroutines.GetVagrantDirectory(), "Vagrantfile")
	}

	boxString, boxVersion, err := GetVagrantFileVersionDetails(vagrantFilePath)
	if err != nil {
		log.Debug(fmt.Sprintf("Could not read from %s", vagrantFilePath))
		return ""
	}

	boxType := GetVagrantBoxType(boxString)

	versionString := fmt.Sprintf("%s (%s)", boxType, boxVersion)
	log.Debug(fmt.Sprintf("GetVagrantFileVersion returns: %s", versionString))

	return versionString
}

func getVagrantVersionDetails(vagrantFileStr string) (string, string, error) {
	boxRegexp := regexp.MustCompile(`config.vm.box = "(.+)"`)
	versionRegexp := regexp.MustCompile(`vb.name = "(.+)"`)

	boxMatches := boxRegexp.FindStringSubmatch(vagrantFileStr)
	versionMatches := versionRegexp.FindStringSubmatch(vagrantFileStr)

	if len(boxMatches) == 2 && len(versionMatches) == 2 {
		log.Debug(fmt.Sprintf("getVagrantVersionDetails returns: [%s] [%s]", boxMatches[1], versionMatches[1]))
		return boxMatches[1], versionMatches[1], nil
	}

	return "", "", errors.New("did not find values from vagrantstring")
}

// GetVagrantFileVersionDetails returns version string (e.g. "digabi/ktp-qa") and
// version string (e.g. "SERVER7108X") from the given vagrantFilePath
func GetVagrantFileVersionDetails(vagrantFilePath string) (string, string, error) {
	fileContent, err := ioutil.ReadFile(filepath.Clean(vagrantFilePath))
	if err != nil {
		log.Debug(fmt.Sprintf("Could not read from %s", vagrantFilePath))
		return "", "", err
	}

	return getVagrantVersionDetails(string(fileContent))
}

// GetVagrantBoxAvailVersion returns a human-readable localised version string
// for a vagrant box available with update
func GetVagrantBoxAvailVersion() string {
	boxString, boxVersion, err := GetVagrantBoxAvailVersionDetails()
	if err != nil {
		log.Debug("Could not get available version string")
		return ""
	}

	boxType := GetVagrantBoxType(boxString)

	versionString := fmt.Sprintf("%s (%s)", boxType, boxVersion)
	log.Debug(fmt.Sprintf("GetVagrantBoxAvailVersion returns: %s", versionString))

	return versionString
}

// GetVagrantBoxAvailVersionDetails gets info about available vagramt box
// from ReallyGetVagrantBoxAvailVersionDetails() or global vagrantBoxAvailVersionDetailsCache
func GetVagrantBoxAvailVersionDetails() (string, string, error) {
	boxString := ""
	boxVersion := ""
	var boxError error

	// There is a avail version fetch going on (break free after 240 loops)
	tryCounter := 0
	for vagrantBoxAvailVersionDetailsCache.updateStarted != 0 && tryCounter < 240 {
		time.Sleep(500)
		tryCounter++
	}

	if vagrantBoxAvailVersionDetailsCache.boxTimestamp < (time.Now().Unix() - constants.VagrantBoxAvailVersionDetailsCacheTimeout) {
		// We need to update the cache
		vagrantBoxAvailVersionDetailsCache.updateStarted = time.Now().Unix()

		boxString, boxVersion, boxError = reallyGetVagrantBoxAvailVersionDetails()
		if boxError == nil {
			vagrantBoxAvailVersionDetailsCache.boxString = boxString
			vagrantBoxAvailVersionDetailsCache.boxVersion = boxVersion
			vagrantBoxAvailVersionDetailsCache.boxTimestamp = time.Now().Unix()
		}

		vagrantBoxAvailVersionDetailsCache.updateStarted = 0
	} else {
		// Return data from the cache

		boxString = vagrantBoxAvailVersionDetailsCache.boxString
		boxVersion = vagrantBoxAvailVersionDetailsCache.boxVersion
	}

	log.Debug(fmt.Sprintf("GetVagrantBoxAvailVersionDetails returns: [%s] [%s]", boxString, boxVersion))
	return boxString, boxVersion, boxError
}

// reallyGetVagrantBoxAvailVersionDetails returns version string (e.g. "digabi/ktp-qa") and
// version string (e.g. "SERVER7108X") by getting Vagrantfile from the AbittiVagrantURL
func reallyGetVagrantBoxAvailVersionDetails() (string, string, error) {
	// Get Abitti Vagrantfile
	strVagrantfile, errVagrantfile := network.DownloadString(constants.AbittiVagrantURL)
	if errVagrantfile != nil {
		log.Debug(fmt.Sprintf("Could not download Abitti Vagrantfile from %s", constants.AbittiVagrantURL))
		return "", "", errors.New("could not download abitti vagrantfile")
	}

	return getVagrantVersionDetails(strVagrantfile)
}

// GetVagrantBoxType returns the type string (Abitti server or Matric Exam server) for vagrant box name
func GetVagrantBoxType(name string) string {
	if GetVagrantBoxTypeIsAbitti(name) {
		return xlate.Get("Abitti server")
	}

	if GetVagrantBoxTypeIsMatriculationExam(name) {
		return xlate.Get("Matric Exam server")
	}

	// Unknown box type
	log.Debug(fmt.Sprintf("Warning: We have a vagrant box type string '%s' which does not resolve to Abitti/Matriculation box type (GetVagrantBoxType)", name))
	return "-"
}

// GetVagrantBoxTypeIsAbitti returns true if given box name string
// belongs to an Abitti vagrant box
func GetVagrantBoxTypeIsAbitti(name string) bool {
	return (name == "digabi/ktp-qa")
}

// GetVagrantBoxTypeIsMatriculationExam returns true if given box name string
// belongs to a Matriculation Examination vagrant box
func GetVagrantBoxTypeIsMatriculationExam(name string) bool {
	re := regexp.MustCompile(`[ksKS]*\d\d\d\d[ksKS]*-\d+`)
	return re.MatchString(name)
}
