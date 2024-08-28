package utils

import (
	"fmt"
	//"strings"
	"time"
)

func Debug(source string, msg string) {
	//if strings.Contains(source, "seat") {
	fmt.Printf("[debug] %s %s %s\n", time.Now().Format(time.RFC3339), source, msg)
	//}
}
