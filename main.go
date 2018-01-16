package main

import (
	"log"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
)

// token=dUQe3kE7aGgdnkBmEAny

// "https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny"

// Mainnet	test network	https://mainnet.infura.io/dUQe3kE7aGgdnkBmEAny
// Ropsten	test network	https://ropsten.infura.io/dUQe3kE7aGgdnkBmEAny
// INFURAnet	test network	https://infuranet.infura.io/dUQe3kE7aGgdnkBmEAny
// Kovan	test network	https://kovan.infura.io/dUQe3kE7aGgdnkBmEAny
// Rinkeby	test network	https://rinkeby.infura.io/dUQe3kE7aGgdnkBmEAny
// IPFS	gateway	https://mainnet.infura.io/dUQe3kE7aGgdnkBmEAny

type Routes struct {
}

type Response struct {
	Result string `json:"result"`
}

func (p *Routes)ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName2(w, r)
		return
	}

	if r.URL.Path == "/test" {
		sayhelloName2(w, r)
		return
	}

	http.NotFound(w, r)
	return
}

func sayhelloName2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  // parse arguments, you have to call this by yourself
	fmt.Println(r.Form)  // print form information in server side
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie2!") // send data to client side
}

func main() {
	resp, err := http.Get("https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny")

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.Body)
	fmt.Println(body)

	var r = new(Response)
	err = json.Unmarshal(body, &r)

	fmt.Println(r.Result)


	router := &Routes{} // set router
	err = http.ListenAndServe(":9090", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
