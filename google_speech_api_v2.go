package main

import (
  "fmt"
  "io/ioutil"
  "bytes"
  "net/http"
)

func main(){
  fmt.Printf("Hello, Go!\n")

  file, err := ioutil.ReadFile("hello.wav")
  if err != nil {
    // 
  }
  s := string(file)
  // fmt.Printf(s)

  httpBody := bytes.NewBufferString(s)

  client := &http.Client{}
  key := "AIzaSyBnCpUVPqPQm6xMcINVVBOS0xYOuxcZ7Jg"
  url := "https://www.google.com/speech-api/v2/recognize?output=json&lang=ru-ru&key="+key
  req, err := http.NewRequest("POST", url, httpBody)

  req.Close = true

  req.Header.Set("Content-Type", "audio/l16; rate=16000")
  // req.SetBasicAuth("user", "pass")
  resp, err := client.Do(req)
  if err != nil {
	    // whatever
  }
  defer resp.Body.Close()

  response, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    // Whatever
  }
  fmt.Println(string(response))
  fmt.Println(resp.Status)

  fmt.Println(resp.Body)

}
