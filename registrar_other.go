// +build !windows

package main
import (
  "encoding/json"
  "os"
  "log"
)

func WriteRegistry(state map[string]*FileState) {
  var newTempRegistrarFile string = appconfig.RegistrarFile + ".new"
  // Open tmp file, write, flush, rename
  log.Printf("Saving registrar state.\n")

  //read the current state file and overwrite the current state vaules
  historical_state := make(map[string]*FileState)
  history, err := os.Open(appconfig.RegistrarFile)
  if err != nil {
    log.Printf("Registar was unable to read privous states. Error: %s\n", err)
    return
  }
  decoder := json.NewDecoder(history)
  decoder.Decode(&historical_state)
  history.Close()


  //loop though the sate for the file, should be a map of lenght 1
  for path, new_state := range state {
    historical_state[path] = new_state
  }

  file, err := os.Create(newTempRegistrarFile)
  if err != nil {
    log.Printf("Failed to open %s for writing new state: %s\n", newTempRegistrarFile, err)
    return
  }

  encoder := json.NewEncoder(file)
  encoder.Encode(historical_state)
  file.Close()

  os.Rename(newTempRegistrarFile, appconfig.RegistrarFile)
}