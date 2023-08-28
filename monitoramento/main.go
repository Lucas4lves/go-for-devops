package main

import (
	"time"
	"fmt"
	"net/http"
	"os"
	"encoding/csv"
)

type Server struct {
	ServerName string
	ServerUrl string
	Runtime float64
	Status int
	DowntimeEvent string
}

func (s *Server) PrintInfo(){
	fmt.Println("Total response time for", s.ServerUrl, ":", s.Runtime, "seconds")
	fmt.Println("Response status:", s.Status)
}

func createServerList (serversList *os.File) []Server{
	csvReader := csv.NewReader(serversList)
	data, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	var servers []Server

	for i, line := range data{
		if i > 0 {
			server := Server{
				ServerName: line[0],
				ServerUrl: line[1],
			}
			servers = append(servers, server)
		}
	}

	return servers
}

func checkServers(servers[]Server){

	for _, server := range servers{

		now := time.Now()

		get, err := http.Get(server.ServerUrl)
		if err != nil {
			fmt.Println("ERRO:", err)
		} 

		server.Runtime = time.Since(now).Seconds()
		server.Status = get.StatusCode
		server.PrintInfo()
	}
}

func openFiles(serversList string, downtimeList string)(*os.File, *os.File){
	//Criando ponteiros para arquivos
	//Args: path, modo (leitura,escrita, apenas-leitura);
	serversList, err := os.OpenFile(arg1, os.O_RDONLY, 0666)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)	
	}

	downtimeList, err := os.OpenFile(arg2, os.O_APPEND |os.O_CREATE, 0666)
	if err != nil{
		fmt.Println(err)
		os.Exit(1)	
	}

	return serversList,downtimeList 
}

func main(){

	fmt.Println("Monitoramento HTTP")
	serverList, downtimeList := openFiles(os.Args[1], os.Args[2])

	defer serverList.Close()
	defer downtimeList.Close()

	//Abrindo o arquivo CSV
	file, err:= os.Open(os.Args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	//Fechando o arquivo	
	defer file.Close()
	
	//Instanciando leitor de CSV (ARG: caminho para o arquivo)
	csvReader := csv.NewReader(file)
	//Lendo o arquivo
	data, err:= csvReader.ReadAll()
	
	if err != nil {
		fmt.Println("ERROR", err)
		os.Exit(1)	
	}
	
	servers := createServerList(data)
	checkServers(servers)
}
