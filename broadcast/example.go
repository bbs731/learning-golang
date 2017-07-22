package main


import (
"log"
"os"
"os/signal"
"sync"
"time"
)

// Extend the sync.Cond struct to add Kill, and Payload
type MySyncCond struct {
	*sync.Cond
	Kill    bool
	Payload int
}

// Add an Init() and set Kill to false (declaritive)
func (m *MySyncCond) Init() {
	m.Kill = false
}

// Extend sync.Cond.Broadcast() to allow a payload
func (m *MySyncCond) Broadcast(payload int) {
	// set the new payload before broadcasting
	m.Payload = payload
	// call sync.Cond.Broadcast()
	m.Cond.Broadcast()
}

// first worker, generates message to be broadcast out
func pinger(c *MySyncCond, wg *sync.WaitGroup) {

	defer wg.Done() // alert the waiting group we are done
	defer c.L.Unlock() // make sure we unlock the condition
	defer log.Println("Terminated pinger") // log a message that we are done here

	// Loop, and send out pings
	count := 0
	for {
		count += 1
		log.Printf("Pinger Pings: %d\n", count)
		// here I use the extended MySyncCond.Broadcast() that allows a payload to be passed
		c.Broadcast(count)
		time.Sleep(5 * time.Second)
	}
}

// our second worker, checks if the ping count is divisible by 1 (true)
// Note: The infinite loop only breaks when the sync.Cond has the field Kill == true
func divByOne(c *MySyncCond, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.L.Unlock() // make sure we unlock
	defer log.Printf("Terminated divByOne")

	for {
		c.L.Lock()
		c.Wait()
		// here I evaluate the Kill field of the MySyncCond struct
		if c.Kill {
			break
		}
		// here I evaluate the Payload field of the MySyncCond struct
		if c.Payload%1 == 0 {
			log.Printf("One\n")
		}
		c.L.Unlock()
	}
}

// basically the same as divByOne, but by 2
func divByTwo(c *MySyncCond, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.L.Unlock() // make sure we unlock
	defer log.Println("Terminated divByTwo")

	for {
		c.L.Lock()
		c.Wait()
		if c.Kill {
			break
		}
		if c.Payload%2 == 0 {
			log.Printf("Two\n")
		}
		c.L.Unlock()
	}
}

// yet another worker, this one divides by 3
func divByThree(c *MySyncCond, wg *sync.WaitGroup) {
	defer wg.Done()
	defer c.L.Unlock()
	defer log.Println("Terminated divByThree")
	for {
		c.L.Lock()
		c.Wait()
		if c.Kill {
			break
		}
		if c.Payload%3 == 0 {
			log.Printf("Three\n")
		}
		c.L.Unlock()
	}
}

// Where the magic starts
func main() {
	// write all the logs to a known file
	logfile, err := os.OpenFile("/tmp/condfun.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	//write stdout to logfile
	//log.SetOutput(logfile)

	// make sure we close the file handler
	defer logfile.Close()

	// Create a new Lock (Mutext), pass it to a new sync.Cond, and the push the sync.Cond into MySyncCond (thus extending it)
	c := MySyncCond{sync.NewCond(new(sync.Mutex)),  false, 0}

	// our first channel, this is used to tell Main that the Waiter is done.
	done := make(chan bool)

	// Kill channel, used to catch Unix Signals
	kill := make(chan os.Signal)
	defer close(kill) // let's make sure to close the channel when we are done!
	signal.Notify(kill) // the magic here assigns Unix Signals to be passed to our channel

	var wg sync.WaitGroup // create a WaitGroup

	wg.Add(4) // tell our WaitGroup how many workers we have
	// fire off all our workers
	// each of these must call wg.Done() or else this whole thing falls apart, hence the defer wg.Done() they all call. To further enforce this we could have used and interface, but that seemed like real work
	go pinger(&c, &wg)
	go divByOne(&c, &wg)
	go divByTwo(&c, &wg)
	go divByThree(&c, &wg)

	// out anonymous worker, this waits for all things in the WaitGroup to complete, if that happens it will pass true to the done channel
	go func(wg *sync.WaitGroup, done chan bool) {
		defer close(done)
		wg.Wait()
		done <- true
	}(&wg, done)

	select {
	case <-kill: // handle the kill signal
		log.Println("Kill Seen")
		c.Kill = true
		c.Cond.Broadcast()
	case <-done: // handle the Waiter finishing (signals all workers stopped)
		log.Println("All workers have reported done")
	}

	log.Println("Terminated Main")
}
