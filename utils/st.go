package utils

import (
	"time"
	st "github.com/showwin/speedtest-go/speedtest"
)

func SpeedTestAll(st_time int) (map[string][][]string) {
	var comms chan []string = make(chan []string)
	var time_stamp = time.Now().Unix()
	var stop_time = time_stamp + int64(st_time)
	var raw_data_final map[string][][]string = make(map[string][][]string)
	go func() {
		for {
			if time.Now().Unix() >= stop_time {
				close(comms)
				comms = nil
				break
			}
			continue
		}
	}()
	var speed_tester = st.New()
	serverLst, _ := speed_tester.FetchServers()
	// update data
	go func() {
		for {
			if comms == nil {
				break
			}
			data := <-comms
			raw_data_final[data[0]] = append(raw_data_final[data[0]], data)
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
			go func(srvre *st.Server) {
				srvre.PingTest(nil)
				srvre.DownloadTest()
				srvre.UploadTest()
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
		}
	}
	return raw_data_final
}

// to fix: concurrent map writes error