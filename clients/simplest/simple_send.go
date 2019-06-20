package main

import (
	"context"
	"pack.ag/amqp"
       )


func main ( ) {
  url     := "amqp://localhost:5672"
  address := "my_address"
  content := []byte("Hello, World!")

  client, _ := amqp.Dial ( url, amqp.ConnSASLAnonymous() )
  defer client.Close()

  session, _ := client.NewSession()
  ctx        := context.Background ( )
  sender,  _ := session.NewSender ( amqp.LinkTargetAddress ( address ) )
  defer sender.Close ( ctx )

  msg := amqp.NewMessage ( content )
  sender.Send ( ctx, msg )
}


