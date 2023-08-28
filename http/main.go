package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	auth "github.com/abbot/go-http-auth"
)


func Secret(user, realm string) string {
	if user == "Lucas"{
		return "$1$LfpFO/8U$NsakbsdAkbl/T9BpSv3pU0" 
	}

	return "" 
}

func main(){
	//Forçando o programa a rodar com 2 argumentos
	if len(os.Args) != 3 {
		fmt.Println("Not suficient arguments provided. It has to have 2")
		os.Exit(1)
	}
	//Argumentos de Linha de Comando
	PORT:= os.Args[1]
	PATH:= os.Args[2]

	//Criar variável do tipo File Server 
	//fs:= http.FileServer(http.Dir(PATH))

	authenticator:= auth.NewBasicAuthenticator("myserver.com", Secret)

	http.HandleFunc("/", authenticator.Wrap(func(w http.ResponseWriter, r *auth.AuthenticatedRequest){
		http.FileServer(http.Dir(PATH)).ServeHTTP(w, &r.Request)
	}))

	fmt.Println("Server running at PORT 3000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))
}
