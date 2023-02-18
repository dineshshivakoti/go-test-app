package main
import (
	"net/http"
	"fmt"
	"io"
	"log"
	"io/ioutil"
	"os"
	"encoding/json"
)

func display(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("https://api.quotable.io/random")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		var res map[string]interface{}
		json.Unmarshal([]byte(body), &res)
		content := fmt.Sprint(res["content"])
		author := fmt.Sprint(res["author"])
		messiah_age := content + " - "+author + "\n"
		io.WriteString(w,messiah_age)

		log.Printf("Serving request: %s", r.URL.Path)
		host, _ := os.Hostname()
		log.Printf("From Hostname: %s\n", host)
		fmt.Fprint(w, "\n\n")
		fmt.Fprintf(w, "From Hostname: %s\n", host)
	}
}

func main() {
	http.HandleFunc("/messiah", display)
	fmt.Printf("Starting server at port 10101.\n")
	if err := http.ListenAndServe(":10101", nil); err != nil {
		log.Fatal(err)
	}

}
