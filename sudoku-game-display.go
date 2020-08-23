package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/urfave/negroni"
)

var (
	dummyGrid matrix
	result = make(chan bool)			//channel define
	upgrader = websocket.Upgrader{}		//websocket connection
)

//genrate the dummy grid to copy anser from user
func copyDummyGrid(grid matrix){

	for i := 0; i < 9; i+=1{
        for j := 0; j < 9; j+=1{
            dummyGrid[i][j] = grid[i][j];
        }
    }
}

//check the userinput grid and anserGrid
func checkWin() bool {
	
	for i:= 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if ansGrid[i][j] != dummyGrid[i][j] {
				return false
			} 
		}
	}
	return true
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "game.html")
}

//send webBrowser data in the string format 
func getStringArray() string {
	
	var str string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			str = str + strconv.Itoa(dummyGrid[i][j])
		}
	}
	return str
}


/* func compareGrid(dummyGrid matrix) bool{

	for i:= 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if(dummyGrid[i][j] != ansGrid[i][j]){
				return false
			}
		}
	}
	return true
} */

/* func notZeroEntry(dummyGrid matrix) bool{
	
	for i:= 0; i < 9; i++{
		for j := 0; j < 9; j++{
			if(dummyGrid[i][j] == 0){
				return false
			}
		}
	}
	return true	

}
 */

//comunications between the server and webBrowser
func newGameHandler(rw http.ResponseWriter, req *http.Request) {
	

	c, err := upgrader.Upgrade(rw, req, nil)
	if err != nil {
		log.Print("Upgrade : ", err)
	}

	str := getStringArray()
	//pass the string array
	c.WriteMessage(websocket.TextMessage, []byte(str))

	for {
		//recive data from the webBrowser
		_, recvData, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			break
		}

		//Extracting data from UI
		data := string(recvData)
		split := strings.Split(data, ",")
		value, _ := strconv.Atoi(split[0])
		row, _ := strconv.Atoi(split[1])
		col, _ := strconv.Atoi(split[2])
		fmt.Println("display----->",data)

		dummyGrid[row][col] = value
		

		if ansGrid[row][col] != dummyGrid[row][col] {
			c.WriteMessage(websocket.TextMessage, []byte("violation"))
		} else {
			if checkWin() {
				c.WriteMessage(websocket.TextMessage, []byte("win"))
				result <- true
				return
			}
		}
	}
}

func InitRouter() (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/", homeHandler).Methods(http.MethodGet)
	router.HandleFunc("/ws", newGameHandler).Methods(http.MethodGet)

	return
}

func serverStart(){

	router := InitRouter()
	server := negroni.Classic()
	server.UseHandler(router)
	server.Run(":9009")
}

func userDisplay(grid matrix)  {

	//fmt.Println("grid----------->",grid)
	t := time.NewTimer(3 *time.Minute)
	copyDummyGrid(grid)
	
	go serverStart()

	select {
	case <- result :
	case <- t.C :
		fmt.Println("Time out")
	}
}