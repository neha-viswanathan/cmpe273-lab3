/*
* Lab 3 - RESTful Server Side Data Cache - Server Instance 2
* Neha Viswanathan
* 010029097
*/
package main

//import statements
import (
        "fmt"
        "net/http"
        "github.com/julienschmidt/httprouter"
        "strconv"
        "encoding/json"
 )

 type KVal struct {
   Key int `json:"key_id"`
   Value string `json:"value"`
 }

 type Collection struct {
   KVList []KVal `json:"kvlist"`
 }

 var kvMap map[int]KVa2

 func main() {
   kvMap = make(map[int]KVal)
   route2 := httprouter.New()
   route2.PUT("/keys/:key_id/:value", PutHandler)
   route2.GET("/keys/:key_id", GetHandler)
   route2.GET("/keys", GetAllHandler)
   server2 := http.Server{
     Addr: "localhost:3001",
     Handler: route2,
   }
   server2.ListenAndServe()
 }

 func PutHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
   var kv KVal
   k, err := strconv.Atoi(p.ByName("key_id"))
   if err != nil {
     panic(err)
   }
   v := p.ByName("value")

   kv.Key = k
   kv.Value = v
   kvMap[k] = kv

   rw.Header().Set("Content-Type", "application/json")
   rw.WriteHeader(http.StatusOK)
   //fmt.Fprintf(rw, "%s", kvMap)
  }

 func GetHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
   k, err := strconv.Atoi(p.ByName("key_id"))
   if err != nil {
     panic(err)
   }

   var kv KVal

   for key, value := range kvMap {
     if key == k {
       kv.Key = k
       kv.Value = value.Value
     }
   }
   resp, _ := json.Marshal(kv)
   rw.Header().Set("Content-Type", "application/json")
   rw.WriteHeader(http.StatusOK)
   fmt.Fprintf(rw, "%s", resp)
  }

 func GetAllHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
  	  var kvList []KVal
      for key, value := range kvMap {
        temp := KVal { key, value.Value, }
        kvList = append(kvList, temp)
    }

    var coll Collection
    coll.KVList = kvList

    resp, _ := json.Marshal(coll)
    rw.Header().Set("Content-Type", "application/json")
    rw.WriteHeader(http.StatusOK)
    fmt.Fprintf(rw, "%s", resp)
  }