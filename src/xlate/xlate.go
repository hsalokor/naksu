package xlate


var current_language string
var xlate_strings map[string]string

func SetLanguage(new_language string) {
  current_language = new_language

  if current_language == "fi" {
    xlate_strings = map[string]string {
      // Main window, groups
      "Language": "Kieli",
      "Basic Functions": "Perustoiminnot",
      "Abitti": "Abitti-koe",
      "Matriculation Exam": "Ylioppilaskoe",
      "Status": "Tila",

      // Main window, buttons
      "Install or update Abitti Stickless Exam Server": "Asenna tai päivitä Abitin tikuton palvelin",
      "Install or update Stickless Matriculation Exam Server": "Asenna tai päivitä yo-kokeen tikuton palvelin",
      "Make Stickless Exam Server Backup": "Tee tikuttomasta palvelimesta varmuuskopio",
      "Exit": "Poistu",

      // Main window, other
      "Current version: %s": "Asennettu versio: %s",
      "Start Stickless Exam Server": "Käynnistä tikuton koetilan palvelin",
      "Naksu update failed. Maybe you don't have network connection?\n\nError: %s": "Naksun päivitys epäonnistui. Ehkä sinulla ei ole juuri nyt verkkoyhteyttä?\n\nVirhe: %s",
      "Did not get a path for a new Vagrantfile": "Uuden Vagrantfile-tiedoston sijainti on annettava",
      "Could not execute vagrant. Are you sure you have installed HashiCorp Vagrant?": "Ohjelman Vagrant käynnistys epäonnistui. Oletko varma, että koneeseen on asennettu HashiCorp Vagrant?",
      "Could not execute VBoxManage. Are you sure you have installed Oracle VirtualBox?": "Ohjelman VBoxManage käynnistys epäonnistui. Oletko varma, että koneeseen on asennettu Oracle VirtualBox?",

      // Backup dialog
      "naksu: SaveTo": "naksu: Tallennuspaikka",
      "Please select target path": "Valitse tallennuspaikka",
      "Save": "Tallenna",
      "Cancel": "Peruuta",

      // mebroutines
      "Abitti server": "Abitti-palvelin",
      "Matric Exam server": "Yo-palvelin",
      "command failed: %s": "komento epäonnistui: %s",
      "Failed to execute %s": "Komennon suorittaminen epäonnistui: %s",
      "Could not chdir to %s": "Hakemistoon %s siirtyminen epäonnistui",
      "Server failed to start. This is typical in Windows after an update. Please try again.": "Palvelimen käynnistys epäonnistui. Tämä on tyypillista Windows-koneissa päivityksen jälkeen. Yritä uudelleen.",
      "Error": "Virhe",
      "Warning": "Varoitus",
      "Info": "Tiedoksi",

      // backup
      "File %s already exists": "Tiedosto %s on jo olemassa",
      "Backup has been made to %s": "Varmuuskopio on talletettu tiedostoon %s",
      "Could not get vagrantbox ID: %d": "Vagrantboxin ID:tä ei voitu lukea: %d",
      "Could not make backup: failed to get disk UUID": "Varmuuskopion ottaminen epäonnistui: levyn UUID:tä ei löytynyt",
      "Could not back up disk %s to %s": "Varmuuskopion ottaminen levystä %s tiedostoon %s epäonnistui",

      // backup, getmediapath
      "Home directory": "Kotihakemisto",
      "Temporary files": "Tilapäishakemisto",
      "Profile directory": "Profiilihakemisto",

      // install
      "Could not change to vagrant directory ~/ktp": "Vagrant-hakemistoon ~/ktp siirtyminen epäonnistui",
      "Error while copying new Vagrantfile: %d": "Uuden Vagrantfile-tiedoston kopiointi epäonnistui: %d",
      "Could not create ~/ktp to %s": "Hakemiston ~/ktp luominen sijaintiin %s epäonnistui",
      "Could not create ~/ktp-jako to %s": "Hakemiston ~/ktp-jako luominen sijaintiin %s epäonnistui",
      "Failed to delete %s": "Tiedoston %s poistaminen epäonnistui",
      "Failed to rename %s to %s": "Tiedoston %s nimeäminen tiedostoksi %s epäonnistui",
      "Failed to create file %s": "Tiedoston %s luominen epäonnistui",
      "Failed to retrieve %s": "Sijainnista %s lataaminen epäonnistui",
      "Could not copy body from %s to %s": "Sisällön %s kopioint sijaintiin %s epäonnistui",

      // start
      // Already defined: "Could not change to vagrant directory ~/ktp"
    }
  } else if current_language == "sv" {
    xlate_strings = map[string]string {
      // Main window, groups
      "Language": "Språk",
      "Basic Functions": "Grundfunktionaliteter",
      "Abitti": "Abitti-prov",
      "Matriculation Exam": "Studentprov",
      "Status": "Status",

      // Main window, buttons
      "Start Stickless Exam Server": "Starta sticklös provlokalsserver",
      "Install or update Abitti Stickless Exam Server": "Installera eller uppdatera sticklös server för Abitti",
      "Install or update Stickless Matriculation Exam Server": "Installera eller uppdatera sticklös server för studentexamen",
      "Make Stickless Exam Server Backup": "Säkerhetskopiera den sticklösa servern",
      "Exit": "Stäng",

      // Main window, other
      "Current version: %s": "Installerad version: %s",
      "Naksu update failed. Maybe you don't have network connection?\n\nError: %s": "Uppdateringen av Naksu misslyckades. Du saknar möjligtvis nätförbindelse för tillfället?\n\nFel: %s",
      "Did not get a path for a new Vagrantfile": "Ge sökvägen för den nya Vagrantfile-filen",
      "Could not execute vagrant. Are you sure you have installed HashiCorp Vagrant?": "Utförandet av programmet Vagrant misslyckades. Är du säker, att HashiCorp Vagrant har installerats på datorn?",
      "Could not execute VBoxManage. Are you sure you have installed Oracle VirtualBox?": "Utförandet av programmet VBoxManage misslyckades. Är du säker, att Oracle VirtualBox har installerats på datorn?",

      // Backup dialog
      "naksu: SaveTo": "naksu: Spara till ",
      "Please select target path": "Välj sökväg",
      "Save": "Spara",
      "Cancel": "Avbryt",

      // mebroutines
      "Abitti server": "Abitti-server",
      "Matric Exam server": "Examensserver",
      "command failed: %s": "Komandot misslyckades: %s",
      "Failed to execute %s": "Utförning av komandot misslyckades: %s",
      "Could not chdir to %s": "Förflyttning till katalogen %s misslyckades",
      "Server failed to start. This is typical in Windows after an update. Please try again.": "Startandet av servern misslyckades. Detta är typiskt i Windows efter en uppdatering. Pröva på nytt.",
      "Error": "Fel",
      "Warning": "Varning",
      "Info": "För information",

      // backup
      "File %s already exists": "Filen %s existerar redan",
      "Backup has been made to %s": "Säkerhetskopian har sparats i filen %s",
      "Could not get vagrantbox ID: %d": "Det gick inte att läsa ID:n på Vagrantboxen: %d",
      "Could not make backup: failed to get disk UUID": "Säkerhetskopieringen misslyckades: skivans UUID hittades inte",
      "Could not back up disk %s to %s": "Säkerhetskopieringen av skivan %s i filen %s misslyckades",

      // backup, getmediapath
      "Home directory": "Hemkatalog",
      "Temporary files": "Tillfällig katalog",
      "Profile directory": "Profilkatalog",

      // install
      "Could not change to vagrant directory ~/ktp": "Förflyttningen till Vagrant-katalogen ~/ktp misslyckades",
      "Error while copying new Vagrantfile: %d": "Kopieringen av en ny Vagrantfile-fil misslyckades: %d",
      "Could not create ~/ktp to %s": "Det gick inte att skapa katalogen ~/ktp i sökvägen %s",
      "Could not create ~/ktp-jako to %s": "Det gick inte att skapa katalogen ~/ktp-jako i sökvägen %s",
      "Failed to delete %s": "Det gick inte att avlägsna filen %s",
      "Failed to rename %s to %s": "Det gick inte att namnge filen %s som %s",
      "Failed to create file %s": "Det gick inte att skapa filen %s",
      "Failed to retrieve %s": "Det gick inte att ladda ner från sökvägen %s",
      "Could not copy body from %s to %s": "Det gick inte att kopiera från sökvägen %s till %s",

      // start
      // Already defined: "Could not change to vagrant directory ~/ktp"
    }
  } else {
    xlate_strings = nil
  }
}

func Get(key string) string {
  if xlate_strings == nil {
    return key
  }

  new_string := xlate_strings[key]

  if new_string == "" {
    return key
  }

  return new_string
}
