package handlers

import (
	"net/http"
	"html/template"
	"os/exec"
	"fmt"
)
func Watch(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("static/html/watcher.html"))
    comments := GetComments()

	t.Execute(w, comments)  // merge.
}

func Ping(w http.ResponseWriter, r *http.Request){
	r.ParseForm()

	command := fmt.Sprintf("ping %s", r.Form.Get("ip"))
	fmt.Println(command)
	
	cmd := exec.Command("/bin/sh", "-c", command)
	std_out, err := cmd.Output()
	
	if err != nil {
		fmt.Fprint(w, "Failed to execute command!")
	}

	std_out_str := string(std_out)

	fmt.Fprint(w, std_out_str)
}