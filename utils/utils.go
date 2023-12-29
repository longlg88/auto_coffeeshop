package utils

import (
	"fmt"
	"time"
)

var Loc *time.Location // Package-level variable to store location

func init() {
	var err error
	Loc, err = time.LoadLocation("Asia/Seoul")
	if err != nil {
		fmt.Println("Error:", err)
	}
}
