package main

import ("flag"
	"fmt"
	//"net"
	"pairbroker/stubs"
	"net/rpc")


type Factory struct {}

//TODO: Define a Multiply function to be accessed via RPC.
func (f *Factory) Multiply(pair stubs.Pair, res stubs.JobReport) (err error){
	 res.Result = (pair.X) *(pair.Y)
	 fmt.Println("Pair multiplication")
	 return
}
//Check the previous weeks' examples to figure out how to do this.


func main(){
	pAddr := flag.String("ip", "127.0.0.1:8050", "IP and port to listen on")
	brokerAddr := flag.String("broker","127.0.0.1:8030", "Address of broker instance")
	flag.Parse()


	client,_ := rpc.Dial("tcp",*brokerAddr)
	defer client.Close()

	subscriptionx := stubs.Subscription{Topic:"multiply",FactoryAddress:*pAddr,Callback:*brokerAddr}
	response:=stubs.JobReport{}
	client.Call(stubs.Subscribe,subscriptionx,response)
	//fmt.Println("Responded: "+response.Message)

	//TODO: You'll need to set up the RPC server, and subscribe to the running broker instance.
}
