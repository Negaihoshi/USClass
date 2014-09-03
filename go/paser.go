package main

import (
	"fmt"
	"github.com/opesun/goquery"
	"io/ioutil"
	"net/http"
	"os"
	// "github.com/franela/goreq"
)

func main() {

	// type Item struct {
	// 	optSclYer string
	// 	optSclTrm string
	// 	optSLSID  string
	// 	optDEPID  string
	// 	optSEL    string
	// 	optSEARCH string
	// 	btnSend   string
	// }
	//
	// item := Item{
	// 	optSclYer: "103",
	// 	optSclTrm: "1",
	// 	optSLSID:  "U",
	// 	optDEPID:  "$",
	// 	optSEL:    "0",
	// 	optSEARCH: "",
	// 	btnSend:   "½Òµ{·j´M",
	// }

	// response, err := http.Get("http://studentsystem.usc.edu.tw/CourseSystem/result_NewNew.asp")
	response, err := http.Get("http://studentsystem.usc.edu.tw/CourseSystem/result_NewNew.asp?optSclYer=103&optSclYer=103&optSclTrm=1&optSLSID=U&optDEPID=%24&optSEL=0&optSEARCH&btnSend=%C2%BD%C3%92%C2%B5{%C2%B7j%C2%B4M")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents))

	x, err := goquery.ParseUrl("http://studentsystem.usc.edu.tw/CourseSystem/result_NewNew.asp?optSclYer=103&optSclYer=103&optSclTrm=1&optSLSID=U&optDEPID=%24&optSEL=0&optSEARCH&btnSend=%C2%BD%C3%92%C2%B5{%C2%B7j%C2%B4M")
	if err != nil {
		panic(err)
	}
	x.Find("#eow-title").Print()
}
