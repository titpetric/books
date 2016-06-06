package main

import (
	"log"
	"fmt"
	"time"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
	"foundations/bootstrap"
)

func json_encode(r interface{}) string {
	jsonString, _ := json.MarshalIndent(r, "", "\t");
	return string(jsonString[:]);
}

type FirstName struct {
	Firstname string `json:"firstname"`
}
func getFirstName(w http.ResponseWriter, r *http.Request) {
	time.Sleep(200 * time.Millisecond);
	value := r.FormValue("firstname")
	response := FirstName{ Firstname: value };
	response_json := json_encode(response);	
	fmt.Fprintf(w, response_json);
}

type LastName struct {
	Lastname string `json:"lastname"`
}
func getLastName(w http.ResponseWriter, r *http.Request) {
	time.Sleep(300 * time.Millisecond);
	value := r.FormValue("lastname");
	response := LastName{ Lastname: value };
	response_json := json_encode(response);
	fmt.Fprintf(w, response_json);
}

type FullName struct {
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}
func getFullName(w http.ResponseWriter, r *http.Request) {
	bootstrap.StartTime = 0;

	firstname_value := r.FormValue("firstname");
	lastname_value := r.FormValue("lastname");

	var firstname FirstName;
	var lastname LastName;
	var fullname FullName;

	data := url.Values{};
	data.Add("firstname", firstname_value);
	data.Add("lastname", lastname_value);

	// fetch firstname

	fn_url := "http://localhost/firstname?" + data.Encode();
	fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), fn_url);
	fn_response, _ := http.Get(fn_url);

	fn_contents, _ := ioutil.ReadAll(fn_response.Body);

	_ = json.Unmarshal(fn_contents, &firstname);
	fullname.Firstname = firstname.Firstname;

	// fetch lastname

	ln_url := "http://localhost/lastname?" + data.Encode();
	fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), ln_url);
	ln_response, _ := http.Get(ln_url);

	ln_contents, _ := ioutil.ReadAll(ln_response.Body);

	fmt.Printf("[%.4f] Done fetching\n", bootstrap.Now());


	_ = json.Unmarshal(ln_contents, &lastname);
	fullname.Lastname = lastname.Lastname;

	// return fullname response
	response_json := json_encode(fullname);
	fmt.Fprintf(w, response_json);
	fmt.Printf("[%.4f] Done with response: %#v\n", bootstrap.Now(), fullname);
}

func main() {
	fmt.Printf("Starting server on port :80\n");
	http.HandleFunc("/fullname", getFullName);
	http.HandleFunc("/firstname", getFirstName);
	http.HandleFunc("/lastname", getLastName);
	err := http.ListenAndServe(":80", nil);
	if (err != nil) {
		log.Fatal("ListenAndServe: ", err);
	}
}