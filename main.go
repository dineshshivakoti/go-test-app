package main
import (
	"net/http"
	"fmt"
	"io"
	"log"
	"io/ioutil"
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
	}
}

func main() {
	http.HandleFunc("/", display)
	fmt.Printf("Starting server at port 8081.\n")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal(err)
	}

}
