package p2p

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/csuito/block/core"
	"github.com/csuito/block/core/types"
	"github.com/davecgh/go-spew/spew"
	net "github.com/libp2p/go-libp2p-net"
)

var (
	mux = &sync.Mutex{}
	bc  = core.Get()
)

func HandleStream(s net.Stream) {
	log.Println("Got a new stream!")
	rw := bufio.NewReadWriter(bufio.NewReader(s), bufio.NewWriter(s))

	go ReadData(rw)
	go WriteData(rw)
}

func ReadData(rw *bufio.ReadWriter) {
	// We create an infinite loop to stay open to peer connections and
	// handle data streams
	for {
		str, err := rw.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if str == "" {
			return
		}
		if str != "\n" {
			chain := make(core.Blockchain, 0)
			if err := json.Unmarshal([]byte(str), &chain); err != nil {
				log.Fatal(err)
			}

			// Following block will execute in mutual exclusion to avoid conflicts
			mux.Lock()
			if len(chain) > len(*bc) {
				*bc = chain
				bytes, err := json.MarshalIndent(bc, "", "	")
				if err != nil {
					log.Fatal(err)
				}
				// Green console color: 	\x1b[32m
				// Reset console color: 	\x1b[0m
				fmt.Printf("\x1b[32m%s\x1b[0m> ", string(bytes))
			}
			mux.Unlock()
		}
	}
}

func WriteData(rw *bufio.ReadWriter) {
	go func() {
		for {
			// Every five seconds we broadcast the latest state of our blockchain to our peers
			time.Sleep(5 * time.Second)
			mux.Lock()
			bytes, err := json.Marshal(bc)
			if err != nil {
				log.Fatal(err)
			}
			mux.Unlock()

			mux.Lock()
			rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
			rw.Flush()
			mux.Unlock()
		}
	}()

	stdReader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		sendData, err := stdReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		sendData = strings.Replace(sendData, "\n", "", -1)

		fmt.Println(sendData)
		bpm, err := strconv.Atoi(sendData)
		if err != nil {
			fmt.Println(err)
		}

		lb := (*bc)[len(*bc)-1]
		b := types.NewBlock(lb, bpm)

		if err := b.Validate(lb); err != nil {
			fmt.Println(err)
		} else {
			mux.Lock()
			bc.AddBlock(b)
			mux.Unlock()
		}

		bytes, err := json.Marshal(bc)
		if err != nil {
			log.Println(err)
		}
		spew.Dump(bc)

		mux.Lock()
		rw.WriteString(fmt.Sprintf("%s\n", string(bytes)))
		rw.Flush()
		mux.Unlock()
	}
}
