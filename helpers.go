package helpers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"os"
	"time"
)

// GetDate ...
func GetDate() string {
	t := time.Now()
	return t.Format(time.RFC3339)
}

// GetMacAddr ...
func GetMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

// GetJSON ...
func GetJSON(v interface{}) []byte {
	res1B, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	return res1B
}

// GetEnv will give env value of default one
func GetEnv(key, fallback string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = fallback
	}
	return value
}
