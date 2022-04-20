package handlers

import (
	"net/http"
	"log"
	"os"
	"bufio"
	"fmt"
)

func Comment(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	
	comment := r.Form.Get("comment")
	fmt.Println(comment)
	
	if comment == ""{
		http.Redirect(w,r, "/watcher", http.StatusFound)
		return
	}

	comments_file, err := os.OpenFile("./comments.list", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil{
		fmt.Println(err)
	}
	defer comments_file.Close()
	
	_, err = comments_file.WriteString(fmt.Sprintf("%s\n",comment))
	if err != nil{
		fmt.Println(err)
	}
	http.Redirect(w,r, "/watcher", http.StatusFound)
}

func GetComments() []string {
	var comments []string 

	comments_file, err := os.Open("./comments.list")
	if err != nil {
		log.Fatal(err)
	}
	defer comments_file.Close()

	scanner := bufio.NewScanner(comments_file)
	for scanner.Scan(){
		comments = append(comments, scanner.Text())
	}

	return comments
}