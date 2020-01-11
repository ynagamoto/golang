package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "log"
    "net/http"
)

func main() {
  api := rest.NewApi()
  api.Use(rest.DefaultDevStack...)

  router, err := rest.MakeRouter(
    rest.Get("/version", getVersion),
  )

  if err != nil {
    log.Fatal(err)
  }

  api.SetApp(router)
  log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func getVersion(w rest.ResponseWriter, r *rest.Request) {
  w.WriteJson(map[string]string{
    "map version": "0.1.2",
    "manual version": "0.1.1",
  })
}
