package main

import (
  "encoding/base64"
  "compress/flate"
  "log"
  "os"
  "io"
)

func main() {
  b64d := base64.NewEncoder(base64.StdEncoding, os.Stdout)
  w, err := flate.NewWriter(b64d, 2)
  if err != nil {
    log.Fatal(err)
  }
  io.Copy(w, os.Stdin)
  w.Flush();
  b64d.Close();
}