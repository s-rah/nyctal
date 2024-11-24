package utils

import (
	"fmt"
	//"strings"
	"time"
)

func Error(client int, source string, msg string) {
	fmt.Printf("[error] %s [client#%d] %s %s\n", time.Now().Format(time.RFC3339), client, source, msg)
}

func Debug(client int, source string, msg string) {
	fmt.Printf("[debug] %s [client#%d] %s %s\n", time.Now().Format(time.RFC3339), client, source, msg)
}
