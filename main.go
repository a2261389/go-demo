package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type httpServer struct{}

func main() {
	var server httpServer
	http.Handle("/users", server)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (hs httpServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	params, ok := r.URL.Query()["id"]

	if !ok || len(params[0]) < 1 {
		log.Println("Url Param is missing")
		w.Write([]byte("No type a [ID] param"))
		return
	}
	param := params[0]
	log.Println("Url Param 'key' is: " + param)
	result, err := hs.findUser(param)
	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Write([]byte(result))
}

func (hs httpServer) findUser(id string) (string, error) {

	content, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return "", err
	}

	userFindID := id
	users := []map[string]interface{}{}
	json.Unmarshal(content, &users)

	result := map[string]interface{}{}

	// search user
	for _, user := range users {
		id := strconv.Itoa(int(user["id"].(float64)))
		if id == userFindID {
			result = user
		}
	}
	if _, ok := result["id"]; !ok {
		return "", errors.New("ID " + userFindID + " User No found")
	}

	response, _ := json.Marshal(result)
	return string(response), nil
}
