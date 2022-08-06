package main  
// create a socket on a specific port
// listen to any attempt of connection to that port
// if the connection is successful, communication can begin between client and server
// the communication should be according to the protocol
// continue listening to port 
// once the connection is closed, the server stops listening and exits

//TODO: NEED TO FIX THE PROTOCOL
import (
	"fmt"
	"net"
	"os"
	"strings"
)

const (
	SERVER_HOST ="localhost"
	SERVER_PORT ="9988"
	SERVER_TYPE ="tcp"
	// INITIAL_STATE = 0
	TRANSACTION_STATE =1
	DEPOSIT_STATE=2
	WITHDRAW_STATE =3
	TRANSFER_STATE =4
	WRONG_INPUT_STATE=5
)

var state int = TRANSACTION_STATE

type User struct  {
	userName string
	accountBalance int
}

func main()  {
	fmt.Println("Server up and running")
	//ESTABLISHING a socket 
	//LISTENING to connection at specific port 
		//the listen function of net package enables this
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)

	if err != nil {
		fmt.Println("Error listening: ", err.Error())
		//quickly exit the execution
		os.Exit(1)
	}

	//at the end close the server so nothing ca
	defer server.Close()
	fmt.Println("Lisitening on "+ SERVER_HOST + ":"+ SERVER_PORT)
	fmt.Println("wating for client...")

	//if the connection is successful, communication can begin between client and server
	for{
		//creates net.conn (connection) by accepting the incoming connection 
		connection, err := server.Accept()
		if err != nil {
				fmt.Println("Error accepting: ", err.Error())
				os.Exit(1)
		}
		
		// the communication should be according to the protocol
		go serverSideProtocol(connection)
	}
}

//conn is an interface that allows to write and read data to and from the connection
func serverSideProtocol(connection net.Conn){
	//load up usernames from temp db
	// var userA = User{"UserA", 100}
	// var userB = User{"UserB", 100}
	// var userC= User{"UserC", 100}

	//Server talks first and informs client about connection establishment
	_, err := connection.Write([]byte("Hello client we are now connected"))
	if err!= nil {
		fmt.Println("Error reading: " , err.Error() )
	}

	for {
		//created an empty array called buffer
		buffer := make([]byte, 1024)
		//read from connection put the contents in buffer
		// the buffer will contain message or the error
		mLen, err:= connection.Read(buffer)
		if err!= nil {
			fmt.Println("Error reading: " , err.Error() )
		}

		var msg = string(buffer[:mLen])

		//if no errors then display the contents of message
		fmt.Println("Received: ", msg)
		if (msg=="bye"){
			_, err = connection.Write([]byte(string(buffer[:mLen])))
			break
		}
		var output = transactionStates(msg)
		_, err =connection.Write([]byte(output))

}
connection.Close()

}

func transactionStates(msg string) string  {
	//write the contents of buffer back to the connection
	var theOutput string 
	
	//check what user wants to do 
	if strings.Contains("Deposit", msg){
		//TODO: deposit money
		theOutput= "Deposited"
	} else if strings.Contains("Withdraw", msg ){
		//TODO: withdraw money
		theOutput ="Withdrawed"
	} else if strings.Contains("Transfer", msg ){
		//TODO: transfer money
		theOutput ="Transfered"
	}	else{
		theOutput ="wrong transaction option!"
	}

	
	return theOutput
}
