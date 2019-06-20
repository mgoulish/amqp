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

  var stop_time int64
  // BUGALERT -- make this an arg, so easier to keep it straight w/ sender.
  expected_messages := 1000 * 1000
  message_count     := 0

  var data []byte

  for {
    msg, err := receiver.Receive ( ctx )
    if err != nil {
      fp ( os.Stdout, "uh-oh: |%s|\n", err.Error() )
    } 

    msg.Accept()
    // I want to get the data here, even though I don't use it,
    // because I am doing timing with this client, and I want to
    // include the time that it takes to make the data available 
    // to the application.
    data = msg.GetData ( )
    message_count ++

    // See if anything's actually happening.
    if 0 == message_count % 100000 {
      fp ( os.Stdout, "  %d\n", message_count )
    }

    if message_count >= expected_messages {
      stop_time = time.Now().UnixNano() / int64(time.Microsecond)
      break
    }
    // fp ( os.Stdout, "Received message : |%s|\n", msg.GetData ( ) )
  }

  fp ( os.Stdout, "recv %d messages received, stop time was %d usec.\n", message_count, stop_time )
  fp ( os.Stdout, "    last message data was: |%s|.\n", data )
}





