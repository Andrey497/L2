package pkg

import (
	"flag"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type telnet struct {
	host    string
	port    string
	timeOut time.Duration
}

func NewTelnet() *telnet {
	return &telnet{}
}
func (t *telnet) Start() error {
	errorChan := make(chan error, 1)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	t.switchFlag()
	client := NewClient(t.host, t.port, t.timeOut, os.Stdin, os.Stdout)
	err := client.BuildConnect()
	if err != nil {
		return err
	}
	go Send(client, errorChan)
	go Receive(client, errorChan)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case err = <-errorChan:
				if err != nil {
					log.Print(err)
					return
				}
			case <-sigChan:
				client.Close()
				return
			default:
				continue
			}
		}
	}()
	wg.Wait()
	return err
}
func (t *telnet) switchFlag() {

	time := flag.Duration("timeout", 10*time.Second, "connectTimeOut")
	flag.Parse()
	args := flag.Args()
	if len(args) != 2 {
		println("Missing required params, usage: \"go-telnet host port\" or \"go-telnet --timeout=10s host port\"")
		println(args)
	}

	t.timeOut = *time
	t.host = args[0]
	t.port = args[1]
}

func Send(client *telnetClient, errorChan chan error) {
	for {
		if err := client.Send(); err != nil {
			errorChan <- err
			return
		}
	}
}
func Receive(client *telnetClient, errorChan chan error) {
	for {
		if err := client.Receive(); err != nil {
			errorChan <- err
			return
		}
	}
}
