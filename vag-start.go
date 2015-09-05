package main

import (
  "fmt"
  "io"
  "os"
  "os/exec"
  "time"
  )

const (
  RG_DIR = "Registration/"
  RG_FILE1 = "d1.bin"
  RG_FILE2 = "d2.bin"
  RG_FILE3 = "d3.bin"
)

func main() {
 go watchFiles();
 fmt.Println("Watching registration files for changes...")
 fmt.Println("Starting VAG-COM")
 cmd := exec.Command("VagCom.exe")
 cmd.Run()
}

func watchFiles() {
  for {
    if RegistrationFilesExist() == false {
      CopyRegistrationFile(RG_FILE1)
      CopyRegistrationFile(RG_FILE2)
      CopyRegistrationFile(RG_FILE3)
    }
    time.Sleep(500 * time.Millisecond)
  }
}
//This functions checks for the existence of the registration files.
//Returns false if they do not exist
//Returns true if they do
func RegistrationFilesExist() bool {
  if _, err := os.Stat(RG_FILE1); os.IsNotExist(err) {
    return false;
  }
  if _, err := os.Stat(RG_FILE2); os.IsNotExist(err) {
    return false;
  }
  if _, err := os.Stat(RG_FILE3); os.IsNotExist(err) {
    return false;
  }
  return true;
}

func CopyRegistrationFile(filename string) {
  file, err := os.Open(RG_DIR+filename)
  dest, err := os.Create(filename)
  defer file.Close()
  defer dest.Close()
  if err != nil {
    panic(err)
  }
  _, err = io.Copy(dest, file)
  if err != nil {
    panic(err)
  }
}
