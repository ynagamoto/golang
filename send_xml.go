package main

import (
  "encoding/base64"
  "io/ioutil"
  "os"
  "github.com/ant0ine/go-json-rest/rest"
  "log"
  "net/http"
)

type Res struct {
  Version string  `json:"version"`
  Data []string   `json:data`
}

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/get_img", getVersion),
  )

  if err != nil {
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getVersion(w rest.ResponseWriter, r *rest.Request) {
  f1, err := os.Open("./activity_main2.xml")
  if err != nil {
    panic(err)
  }
  defer f1.Close()

  data1, _ := ioutil.ReadAll(f1)
  data1_base64 := base64.StdEncoding.EncodeToString(data1)

  res := Res{
    Version: "0.2.2",
    Data: []string{
      data1_base64,
    },
  }

  w.WriteJson(&res)
}
