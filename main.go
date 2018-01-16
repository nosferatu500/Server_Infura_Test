package main

import (
	"log"
	"net/http"
	"fmt"
	"strings"
	"io/ioutil"
	"encoding/json"
	"strconv"
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

	if r.URL.Path == "/blockNumber" {
		getBlockNumber(w, r)
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

func getBlockNumber(w http.ResponseWriter, r *http.Request) {
	resp, _ := http.Get(`https://api.infura.io/v1/jsonrpc/mainnet/eth_blockNumber?token=dUQe3kE7aGgdnkBmEAny`)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)


	fmt.Println(resp.Body)
	fmt.Println(body)

	var res = new(Response)
	_ = json.Unmarshal(body, &res)

	data := &Response{}

	hex := strings.TrimPrefix(res.Result, "0x")
	x, _ := strconv.ParseInt(hex, 16, 64)
	s := strconv.Itoa(int(x))

	data.Result = s
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(data)
	//fmt.Fprintf(w, res.Result)
}

func main() {

	router := &Routes{} // set router
	err := http.ListenAndServe(":9090", router) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
