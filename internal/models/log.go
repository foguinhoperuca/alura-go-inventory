package models

import (
	"fmt"
	"time"
)

type Log struct {
	Timestamp time.Time
	Action    string
	User      string
	ItemID    int
	Quantity  int
	Reason    string
}

func (l *Log) Info() string {
	return fmt.Sprintf("[%s] Action: %s | User: %s | ItemID: %d | Quantity: %d | Reason: %s", l.Timestamp.Format("01/02 15:04:05"), l.Action, l.User, l.ItemID, l.Quantity, l.Reason)
}
