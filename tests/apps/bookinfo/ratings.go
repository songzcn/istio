// Copyright 2017 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type rating struct {
	Stars int `json:"stars,omitempty"`
}

var ratings = map[string]*rating{
	"reviewer1": {
		Stars: 5,
	},
	"reviewer2": {
		Stars: 4,
	},
}

func main() {
	port := "9080"

	http.HandleFunc("/ratings", ratingsHandler)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func ratingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	bytes, _ := json.Marshal(ratings)
	w.Write(bytes)
}