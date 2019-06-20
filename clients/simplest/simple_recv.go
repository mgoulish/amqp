package main

import (
	"context"
        "time"
        "fmt"
        "os"
	"pack.ag/amqp"
       )


var fp = fmt.Fprintf


func main() {
  url     := "amqp://localhost:5672"
  address := "my_address"

  client, _ := amqp.Dial ( url, amqp.ConnSASLAnonymous() )
  defer client.Close()

  session, _  := client.NewSession()
  ctx         := context.Background ( )
  receiver, _ := session.NewReceiver ( amqp.LinkSourceAddress ( address ),
                                       amqp.LinkCredit ( 10 ) )
  defer func () {
          ctx, cancel := context.WithTimeout ( ctx, 1 * time.Second )
          receiver.Close ( ctx )
          cancel ( )
        } ( )

  msg, err := receiver.Receive ( ctx )
  if err != nil {
    fp ( os.Stdout, "uh-oh: |%s|\n", err.Error() )
  } 

  msg.Accept()
  fp ( os.Stdout, "Received message : |%s|\n", msg.GetData ( ) )
}



