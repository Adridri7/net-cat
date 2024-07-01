package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"

	"colors"
	"netcat_user"
)

//var clients map[net.Conn]string

var (
	clientsArray []netcat_user.User
	orange       string
)

const (
	joinLeaveMsgColor = colors.FgBrightCyan
	serverMsgColor    = colors.FgBrightYellow
	clientPromptColor = colors.FgBrightGreen
	errorColor        = colors.FgBrightRed
)

func init() {
	//clients = make(map[net.Conn]string)
	clientsArray = make([]netcat_user.User, 0)

	errOrng := error(nil)
	orange, errOrng = colors.NewFGColorRGB(255, 127, 0)
	if errOrng != nil {
		fmt.Fprintln(os.Stderr, "Could not create orange:", errOrng.Error())
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	handledClient := netcat_user.NewUser(conn, "", colors.NewRandomFGColorRGB())

	fmt.Println(colors.SprintfANSI(fmt.Sprintf("New client connected: %s", handledClient.Connection.RemoteAddr().String()), serverMsgColor, colors.BgReset))
	displayWelcomeMessage(handledClient.Connection)

	//sendMessage(handledClient.Connection, colors.SprintfANSI("[ENTER YOUR MESSAGE]: ", clientPromptColor, colors.BgReset))

	reader := bufio.NewReader(handledClient.Connection)

	for {
		if handledClient.Name == "" {
			sendMessage(conn, colors.SprintfANSI("[ENTER YOUR NAME]: ", clientPromptColor, colors.BgReset))
			username, err := reader.ReadString('\n')
			if err != nil {
				clientsArray = removeClient(clientsArray, handledClient)
				msg := colors.SprintfANSI(fmt.Sprintf("%s has left the server.", handledClient.Connection.RemoteAddr().String()), serverMsgColor, colors.BgReset)
				fmt.Println(msg)
				broadcastMessage(msg, handledClient.Connection)
				return
			}

			username = strings.TrimSpace(username)
			handledClient.Name = username
			broadcastMessage(handledClient.ColoredUsername()+colors.SprintfANSI(" has joined the chat.", joinLeaveMsgColor, colors.BgReset), handledClient.Connection)
		} else {
			sendMessage(conn, "> ")
			msg, err := reader.ReadString('\n')
			if err != nil {
				//	fmt.Println(color.Yellow.Sprintf("%s has left the server.", clients[conn]))
				broadcastMessage(handledClient.ColoredUsername()+colors.SprintfANSI(" has left the chat.", joinLeaveMsgColor, colors.BgReset), handledClient.Connection)
				clientsArray = removeClient(clientsArray, handledClient)
				break
			}
			msg = strings.TrimSpace(msg)

			if msg != "" {
				if msg[0] == '/' {
					switch msg[1:] {
					case "quit":
						//broadcastMessage(color.Yellow.Sprintf(username, "has left the chat.", clients[conn]), conn)
						broadcastMessage(handledClient.ColoredUsername()+colors.SprintfANSI(" has left the chat.", joinLeaveMsgColor, colors.BgReset), handledClient.Connection)
						clientsArray = removeClient(clientsArray, handledClient)
						return
					case "color":
					case "clear":
					default:
						broadcastMessage("["+time.Now().Format("2006-01-02 15:04:05")+"]["+handledClient.ColoredUsername()+"]: "+colors.SprintfANSI(msg, colors.FgBrightWhite, colors.BgReset), handledClient.Connection)
					}

				} else {
					//broadcastMessage(color.LightRed.Sprintf("%s: %s", clients[conn], msg), conn)
					broadcastMessage("["+time.Now().Format("2006-01-02 15:04:05")+"]["+handledClient.ColoredUsername()+"]: "+colors.SprintfANSI(msg, colors.FgBrightWhite, colors.BgReset), handledClient.Connection)
				}
			}
		}
	}
}

func broadcastMessage(message string, senderConn net.Conn) {
	fmt.Println(message)
	for _, client := range clientsArray {
		if client.Connection != senderConn {
			_, err := client.Connection.Write([]byte(message + "\n"))
			if err != nil {
				fmt.Println(err.Error())
			}
		}
	}
}

func sendMessage(conn net.Conn, message string) {
	_, err := conn.Write([]byte(message))
	if err != nil {
		fmt.Println("Error sending message:", err.Error())
	}
}

func removeClient(arr []netcat_user.User, toRemove *netcat_user.User) []netcat_user.User {
	newArr := make([]netcat_user.User, 0)

	for _, client := range arr {
		if !(&client == toRemove) {
			newArr = append(newArr, client)
		}
	}

	return newArr
}

func displayWelcomeMessage(conn net.Conn) {
	fmt.Fprintln(conn, "Welcome to TCP-Chat!")
	/*
		fmt.Fprintln(conn, "         _nnnn_\n"+
			"        dGGGGMMb\n"+
			"       @p~qp~~qMb\n"+
			"       M|@||@) M|\n"+
			"       @,----.JM|\n"+
			"      JS^\\__/  qKL\n"+
			"     dZP        qKRb\n"+
			"    dZP          qKKb\n"+
			"   fZP            SMMb\n"+
			"   HZM            MMMM\n"+
			"   FqM            MMMM\n"+
			" __| \".        |\\dS\"qML\n"+
			" |    `.       | `' \\Zq\n"+
			"_)      \\.___.,|     .'\n"+
			"\\____   )MMMMMP|   .'\n"+
			"     `-'       `--'")
	*/
	fmt.Fprintln(conn, "         _nnnn_\n"+
		"        dGGGGMMb\n"+
		"       @p~qp~~qMb\n"+
		"       M|@||@) M|\n"+
		"       @,"+orange+"----."+colors.ResetFGColorTag+"JM|\n"+
		"      JS^"+orange+"\\__/"+colors.ResetFGColorTag+"  qKL\n"+
		"     dZP        qKRb\n"+
		"    dZP          qKKb\n"+
		"   fZP            SMMb\n"+
		"   HZM            MMMM\n"+
		"   FqM            MMMM\n"+
		" "+orange+"__| \"."+colors.ResetFGColorTag+"        "+orange+"|\\"+colors.ResetFGColorTag+"dS"+orange+"\""+colors.ResetFGColorTag+"qML\n"+
		" "+orange+"|    `."+colors.ResetFGColorTag+"       "+orange+"| `' \\"+colors.ResetFGColorTag+"Zq\n"+
		orange+"_)      \\"+colors.ResetFGColorTag+".___.,"+orange+"|     .'"+colors.ResetFGColorTag+"\n"+
		orange+"\\____   )"+colors.ResetFGColorTag+"MMMMMP"+orange+"|   .'"+colors.ResetFGColorTag+"\n"+
		orange+"     `-'       `--'"+colors.ResetFGColorTag)
	fmt.Fprintln(conn, "Enter :", colors.SprintfANSI("/quit", serverMsgColor, colors.BgReset), "to exit the server.")
	//fmt.Fprintln(conn, "Enter :", colors.SprintfANSI("/clear", serverMsgColor, colors.BgReset), "to clear the chat")
	//fmt.Fprintln(conn, "Enter :", colors.SprintfANSI("/color", serverMsgColor, colors.BgReset), "to change your message color")
	fmt.Fprintln(conn, "")
}
