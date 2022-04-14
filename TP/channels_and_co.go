package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Response struct {
	respText string
	err      error
}

func main() {

	channelResp1 := make(chan Response)
	channelResp2 := make(chan Response)
	// http://localhost:4000/?id=id1
	// http://localhost:4000/

	go callServer("http://localhost:4000/?id=id1", channelResp1)
	go callServer("http://localhost:4000/?id=id2", channelResp2)

	// select {
	// case msg1 := <-channelResp1:
	// 	printResponse(msg1)
	// case msg2 := <-channelResp2:
	// 	printResponse(msg2)
	// }

	/**
	res1 := <-channelResp1
	res2 := <-channelResp2

	fmt.Println("res2 :")
	if res2.err == nil {
		fmt.Println(res2.respText)
	} else {
		fmt.Println(res2.err.Error())
	}
	**/
	println("Fin du programme")
}

func printResponse(r Response) {
	fmt.Println("res1 :")
	if r.err == nil {
		fmt.Println(r.respText)
	} else {
		fmt.Println(r.err.Error())
	}
}

func callServer(address string, c1 chan Response) {
	var r Response

	resp, err := http.Get(address)

	if err == nil {
		if resp.StatusCode != 200 {
			r.err = errors.New("Le code retourné par le serveur indique une erreur: " + strconv.Itoa(resp.StatusCode))
		}

	} else {
		r.err = err
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		r.err = err
	} else {
		r.respText = string(body)
	}

	c1 <- r

}
