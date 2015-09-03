package command

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"os"
	"reflect"
	"strings"
	"time"
)

// ChannelCommand is a Command implementation for specifying the command...
type ChannelCommand struct {
	UI cli.Ui
}

//Help returns the help for the command
func (channelCommand *ChannelCommand) Help() string {
	helpText := `
Usage: effective-go channels [options]

  A description of the channels is available here
  https://golang.org/doc/effective_go.html#concurrency

Options:
  -code-snipet = [1..]
`
	return strings.TrimSpace(helpText)
}

//CodeSnipet1 ...
func (channelCommand *ChannelCommand) CodeSnipet1() {

	code := `
--------------------------------------------------------------------------------
    ci := make(chan int)            // unbuffered channel of integers
    cj := make(chan int, 0)         // unbuffered channel of integers
    cs := make(chan *os.File, 100)  // buffered channel of pointers to Files
--------------------------------------------------------------------------------
    `
	fmt.Print(code)

	fmt.Print("nothing to do, but that's how you make channels!\n\n")

}

func doSomethingForAWhile() {
	fmt.Print("\ndoing something for a while, this is gonna take at least 3secs...\n\n")
	time.Sleep(3 * time.Second)
	fmt.Print("\nthat's that done!\n\n")
}

func pretendSort() {
	fmt.Print("\nphew... sorting some stuff, be right back...\n\n")
	time.Sleep(2 * time.Second)
	fmt.Print("\nok, that's that sorted!\n\n")

}

//Request ...
type Request struct {
	args       []int
	f          func([]int) int
	resultChan chan int
}

var sem = make(chan int, 4)

func process(r *Request) {
	doSomethingForAWhile()
}
func handle(r *Request) {
	sem <- 1   // Wait for active queue to drain.
	process(r) // May take a long time.
	<-sem      // Done; enable next request to run.
}

func serve(queue chan *Request) {
	for {
		req := <-queue
		go handle(req) // Don't wait for handle to finish.
	}
}

func serve2(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func() {
			process(req) // Buggy; see explanation below.
			<-sem
		}()
	}
}

func serve3(queue chan *Request) {
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}

func serve4(queue chan *Request) {
	for req := range queue {
		req := req // Create new instance of req for the goroutine.
		sem <- 1
		go func() {
			process(req)
			<-sem
		}()
	}
}

//CodeSnipet3 ...
func (channelCommand *ChannelCommand) CodeSnipet3() {

	code := `
--------------------------------------------------------------------------------
  var sem = make(chan int, MaxOutstanding)

  func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
  }

  func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
  }
--------------------------------------------------------------------------------
`
	fmt.Print(code)
	fmt.Print("\nnow running it...\n\n")
	queue := make(chan *Request, 4)
	go serve(queue)
	r := Request{}

	fmt.Print("\nputting 4 requests on the queue...\n\n")
	queue <- &r
	queue <- &r
	queue <- &r
	queue <- &r
}

//CodeSnipet4 ...
func (channelCommand *ChannelCommand) CodeSnipet4() {

	code := `
--------------------------------------------------------------------------------
  func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below.
            <-sem
        }()
    }
  }
--------------------------------------------------------------------------------
`
	fmt.Print(code)
	fmt.Print("\nnow running it...\n\n")
	queue := make(chan *Request, 4)
	go serve2(queue)
	r := Request{}

	fmt.Print("\nputting 8 requests on the queue, but we have a limit of 4...\n\n")
	for x := 0; x < 8; x++ {
		queue <- &r
	}

}

//CodeSnipet5 ...
func (channelCommand *ChannelCommand) CodeSnipet5() {

	code := `
--------------------------------------------------------------------------------
  func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
  }
--------------------------------------------------------------------------------
`
	fmt.Print(code)
	fmt.Print("\nnow running it...\n\n")
	queue := make(chan *Request, 4)
	go serve3(queue)
	r := Request{}

	fmt.Print("\nputting 8 requests on the queue, but we have a limit of 4...\n\n")
	for x := 0; x < 8; x++ {
		queue <- &r
	}
}

//CodeSnipet6 ...
func (channelCommand *ChannelCommand) CodeSnipet6() {

	code := `
--------------------------------------------------------------------------------
  func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
  }
--------------------------------------------------------------------------------
`
	fmt.Print(code)
	fmt.Print("\nnow running it...\n\n")
	queue := make(chan *Request, 4)
	go serve4(queue)
	r := Request{}

	fmt.Print("\nputting 8 requests on the queue, but we have a limit of 4...\n\n")
	for x := 0; x < 8; x++ {
		queue <- &r
	}

}

//CodeSnipet2 ...
func (channelCommand *ChannelCommand) CodeSnipet2() {

	code := `
--------------------------------------------------------------------------------
  c := make(chan int)  // Allocate a channel.
  // Start the sort in a goroutine; when it completes, signal on the channel.
  go func() {
      list.Sort()
      c <- 1  // Send a signal; value does not matter.
  }()
  doSomethingForAWhile()
  <-c   // Wait for sort to finish; discard sent value.
--------------------------------------------------------------------------------
  `
	fmt.Print(code)
	fmt.Print("\nnow running it...\n\n")

	c := make(chan int) // Allocate a channel.
	// Start the sort in a goroutine; when it completes, signal on the channel.
	go func() {
		pretendSort()
		c <- 1 // Send a signal; value does not matter.
	}()
	doSomethingForAWhile()
	<-c // Wait for sort to finish; discard sent value.

}

//Run runs the command
func (channelCommand *ChannelCommand) Run(args []string) int {

	var codeSnipet string
	cmdFlags := flag.NewFlagSet("channelcommand", flag.ContinueOnError)
	cmdFlags.Usage = func() { channelCommand.UI.Output(channelCommand.Help()) }
	cmdFlags.StringVar(&codeSnipet, "code-snipet", "", "")
	if err := cmdFlags.Parse(args); err != nil {
		return 1
	}

	methodName := "CodeSnipet" + codeSnipet
	fmt.Printf("calling %s\n\n", methodName)

	//Call the corresponding method by reflection
	var t ChannelCommand
	reflect.ValueOf(&t).MethodByName(methodName).Call([]reflect.Value{})

	fmt.Print("\npress enter to exit\n")
	reader := bufio.NewReader(os.Stdin)
	reader.ReadString('\n')

	channelCommand.UI.Output("ChannelCommand Complete")
	return 0
}

//Synopsis resturns the synopsis for the command
func (channelCommand *ChannelCommand) Synopsis() string {
	return "Run code snipets and get link for channels in effective go."
}
