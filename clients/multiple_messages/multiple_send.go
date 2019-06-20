package main


import (
	"context"
        "time"
        "fmt"
        "os"
	"pack.ag/amqp"
       )



var fp = fmt.Fprintf 



func main ( ) {
  url     := "amqp://localhost:5672"
  address := "my_address"
  // 100-byte payload.
  content := []byte("0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789")

  client, _ := amqp.Dial ( url, amqp.ConnSASLAnonymous() )
  defer client.Close()

  session, _ := client.NewSession()
  ctx        := context.Background ( )
  sender,  _ := session.NewSender ( amqp.LinkTargetAddress ( address ) )
  defer sender.Close ( ctx )

  msg := amqp.NewMessage ( content )

  // BUGALERT -- make this an arg, so easier to keep it straight w/ receiver.
  n_messages := 1000 * 1000

  start_time := time.Now().UnixNano() / int64(time.Microsecond)


  for i := 0; i < n_messages; i ++ {
    sender.Send ( ctx, msg )
  }

  fp ( os.Stdout, "sender: %d messages sent. Start time was %d usec.\n", n_messages, start_time )
}





