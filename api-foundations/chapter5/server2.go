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
	fmt.Printf("[%.4f] Request with firstname\n", bootstrap.Now());
	time.Sleep(300 * time.Millisecond);
	value := r.FormValue("firstname")
	response := FirstName{ Firstname: value };
	response_json := json_encode(response);	
	fmt.Fprintf(w, response_json);
	fmt.Printf("[%.4f] Response with firstname\n", bootstrap.Now());
}

type LastName struct {
	Lastname string `json:"lastname"`
}
func getLastName(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("[%.4f] Request with lastname\n", bootstrap.Now());
	time.Sleep(200 * time.Millisecond);
	value := r.FormValue("lastname");
	response := LastName{ Lastname: value };
	response_json := json_encode(response);
	fmt.Fprintf(w, response_json);
	fmt.Printf("[%.4f] Response with lastname\n", bootstrap.Now());
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
	fn_chan := make(chan []byte, 1);
	go func() {
		fn_url := "http://localhost/firstname?" + data.Encode();
		fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), fn_url);
		fn_response, _ := http.Get(fn_url);
		contents, _ := ioutil.ReadAll(fn_response.Body);
		fn_chan <- contents;
	}()


	// fetch lastname
	ln_chan := make(chan []byte, 1);
	go func() {
		ln_url := "http://localhost/lastname?" + data.Encode();
		fmt.Printf("[%.4f] Fetching url: %s\n", bootstrap.Now(), ln_url);
		ln_response, _ := http.Get(ln_url);
		contents, _ := ioutil.ReadAll(ln_response.Body);
		ln_chan <- contents;
	}()

	// fetch response data

	fn_contents := <-fn_chan;
	_ = json.Unmarshal(fn_contents, &firstname);
	fullname.Firstname = firstname.Firstname;

	ln_contents := <-ln_chan;
	_ = json.Unmarshal(ln_contents, &lastname);
	fullname.Lastname = lastname.Lastname;

	fmt.Printf("[%.4f] Done fetching\n", bootstrap.Now());

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