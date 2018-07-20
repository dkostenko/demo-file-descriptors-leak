package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
)

var pid string

func handler(w http.ResponseWriter, r *http.Request) {
	body, _ := ioutil.ReadAll(r.Body)

	// Case 1.
	os.Open(string(body))

	// Case 2.
	// resp, _ := http.Get(string(body))
	// ioutil.ReadAll(resp.Body)
	// defer resp.Body.Close()

	cmd := exec.Command("lsof", "-p", pid)
	lsofOutput, _ := cmd.Output()
	fmt.Println(string(lsofOutput))

	w.Write([]byte(fmt.Sprintf("resp_data: %s\n", body)))
}

func main() {
	pid = strconv.Itoa(os.Getpid())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
