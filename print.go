package print

import (
	"fmt"
	"time"
)

// Print test message
func Print(cnt int) {
	fmt.Println("I'm directcd_test_module", cnt, time.Now())
}
