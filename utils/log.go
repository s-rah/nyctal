package utils

import (
	"fmt"
	//"strings"
	"time"
)

func Debug(client int, source string, msg string) {
	//if strings.Contains(source, "seat") {
	fmt.Printf("[debug] %s [client#%d] %s %s\n", time.Now().Format(time.RFC3339), client, source, msg)
	//}
}
