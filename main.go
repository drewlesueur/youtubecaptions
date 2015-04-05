package main

//http://grahamc.com/blog/openssl-madness-how-to-create-keys-certificate-signing-requests-authorities-and-pem-files/
import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	fmt.Println("yo")
	var port string
	if len(os.Args) >= 2 {
		fmt.Println(os.Args[1])
		port = os.Args[1]
	} else {
		port = "8003"
	}
	code = strings.Replace(code, "{PORT}", port, 1)
	fmt.Println("running on port " + port)

	http.HandleFunc("/", mainHanlder)
	http.HandleFunc("/send", sendHandler)
	//log.Fatal(http.ListenAndServe(":"+port, nil))
	//log.Fatal(http.ListenAndServeTLS(":"+port, "cert.pem", "key.pem", nil))
	//log.Fatal(http.ListenAndServeTLS(":"+port, "lcc.drewles.com.pem", "lcc.drewles.com.key", nil))
	log.Fatal(http.ListenAndServeTLS(":"+port, "lcc.drewles.com.pem", "lcc.drewles.com.key", nil))
}

var code = `

//Run this code in the javascript console of any youtube video with captions (disable flash too)
//==============================================================================================

function sendCaptions() {
	window.sendingCaptions = true
	var lastText = ""
	var name = prompt("give this a short name")
	if (!name) {
		name = "default"	
	}
	setInterval(function () {
		var caption = document.querySelector('.captions-text')	
		if (caption) {
			var text = caption.innerText		
			if (text != lastText) {
				var script = document.createElement('script')
				var jsonp = Date.now()
				window[jsonp] = function () {
					// clean up
					document.body.removeChild(script)	
					delete window[jsonp]
				}
				script.src = "https://lcc.drewles.com:{PORT}/send?name="+name+"&text=" + encodeURIComponent(text) + "&jsonp=" + jsonp 
				document.body.appendChild(script)
				lastText = text
			}
		}
	}, 500)
}
if (!window.sendingCaptions) {
	sendCaptions();	
}
`

func mainHanlder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	if r.URL.Path == "favicon.ico" {
		//w.Write([]byte(""))
		return
	} else if r.URL.Path == "/" {
		fmt.Fprintf(w, code)
	} else {
		fmt.Fprintf(w, r.URL.Path)
	}
}

func sendHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	text := r.Form.Get("text")
	name := r.Form.Get("name")
	_ = name
	fmt.Println(text)
}

//technology and revelation combine
