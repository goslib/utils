package utlDate

import (
	"fmt"
	"testing"
	"time"
)

func TestGetUtcDayId(t *testing.T) {
	did := GetUtcDayId(time.Now())
	fmt.Println(did, FromUtcDayId(did))
}

func TestGetWeekId(t *testing.T) {
	did := GetWeekId(time.Now())
	fmt.Println(did, FromWeekId(did))
}
