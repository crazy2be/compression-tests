package main

import (
  "encoding/base64"
  "compress/flate"
  "fmt"
  "log"
  "os"
)

func main() {
  b64d := base64.NewEncoder(base64.StdEncoding, os.Stdout)
//   if err != nil {
//     log.Fatal(err)
//   }
  w, err := flate.NewWriter(b64d, 2)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Fprintln(w, "This is a test!");
  w.Flush();
  b64d.Close();
  fmt.Fprintln(os.Stdout);
}