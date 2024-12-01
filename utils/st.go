package utils

import (
	"fmt"
	"sync"
	"time"

	st "github.com/showwin/speedtest-go/speedtest"
)

func SpeedTestAll(stTime int) map[string][][]string {
	var comms = make(chan []string)
	var timeStamp = time.Now().Unix()
	var stopTime = timeStamp + int64(stTime)
	var rawDataFinal = make(map[string][][]string)
	var wg sync.WaitGroup
	go func() {
		for {
			if time.Now().Unix() >= stopTime {
				close(comms)
				comms = nil
				break
			}
			continue
		}
	}()
	var speedTester = st.New()
	serverLst, _ := speedTester.FetchServers()
	// update data
	go func() {
		for {
			if comms == nil {
				break
			}
			data := <-comms
			if len(data) > 0 {
				rawDataFinal[data[0]] = append(rawDataFinal[data[0]], data)
			}
			continue
		}
	}()
	// test against all servers in serverLst until time runs out
	for {
		if comms == nil {
			break
		}
		// ready_to_test_again_redundant is not used, so this assignment is removed
		for _, srvr := range serverLst {
			if comms == nil {
				break
			}
			wg.Add(1)
			go func(srvre *st.Server) {
				defer wg.Done()
				err := srvre.PingTest(nil)
				if err != nil {
					fmt.Println(err)
				}
				err = srvre.DownloadTest()
				if err != nil {
					fmt.Println(err)
				}
				err = srvre.UploadTest()
				if err != nil {
					fmt.Println(err)
				}
				stats := []string{
					srvre.Name,
					srvre.Host,
					srvre.Latency.String(),
					srvre.DLSpeed.String(),
					srvre.ULSpeed.String(),
				}
				if comms != nil {
					comms <- stats
				} else {
					return
				}
			}(srvr)
			wg.Wait()
		}
	}
	return rawDataFinal
}

// to fix: concurrent map writes error
