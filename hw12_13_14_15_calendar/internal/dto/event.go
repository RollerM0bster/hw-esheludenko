package dto

import "time"

type Event struct {
	ID                     int64
	Title                  string
	Start                  time.Time
	End                    time.Time
	Description            string
	OwnerId                int64
	DaysAmountBeforeNotify int8
}
