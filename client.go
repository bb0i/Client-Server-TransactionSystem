package main 
//Client Goal
	//Establish connection to remote host
	//check if any connection has occurred or not
	//send and receive bytes of information
	//finally, close the connection
import (
	//net package provides the necessary API to implement socket communication
	"net"
	"fmt"
	
)

const (
	SERVER_HOST ="localhost"
	SERVER_PORT ="9988"
	SERVER_TYPE ="tcp"
)

func main()  {
	//ESTABLISHING connection
		//Dial requires defining: type of socket and socket address (i.e. servername: portno)
		//Dial returns either with connection or with the error log
	connection, err := net.Dial(SERVER_TYPE,SERVER_HOST+":"+SERVER_PORT)

	//CHECKING if connection has occured 
	if err!=nil{
		//if error is returned we abort execution and display the error message
		panic(err)
	}

	clientSideProtocol(connection)
	//CLOSE connection
	defer connection.Close()

}

func clientSideProtocol(connection net.Conn)  {
	//create buffer where incoming messages can be saved
	buffer := make([]byte, 1024)
	//READ from the buffer and store the contents either in mLen or within error
	mLen, err := connection.Read(buffer)
	// if there is error display the error
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	//Gets message from server regarding connection establishment
	fmt.Println("Received: ", string(buffer[:mLen]))
	
	//SEND some data 
	// _, err = connection.Write([]byte("Hello server! Greetings."))
	fmt.Println("Your account balance is : £100 ")
	fmt.Println("Enter the following TRANSACTION options with the value example:\n 1) Deposit £---\n 2) Withdraw £--- \n 3) Transfer userX £--- ")
	
	//user enters an input
	for {
		var userInput string
		//takes in the input in double quotes, allows to accept input like deposit 100 as 1 input rather than 2
		fmt.Scanf("%q", &userInput)
		if(userInput !=""){
			fmt.Println("SENDING: " +userInput  )
			_, err =connection.Write([]byte(userInput))
			mLen, err = connection.Read(buffer)	
			fmt.Println("Received: ", string(buffer[:mLen]))
		}
		if(string(buffer[:mLen])=="bye"){
			break
		}
		
	}

}