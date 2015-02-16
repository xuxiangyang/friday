package friday

import (
	"time"
)

var Holidays, Exchanges OriginDates

type Date struct {
	OriginDate
	NeedWork bool `json:need_work`
}

func NewDate(t time.Time) *Date {
	date := Date{OriginDate: OriginDate{Year: t.Year(), Month: int(t.Month()), Day: t.Day()}, NeedWork: true}
	date.parseDescribe()
	return &date
}

func (date *Date) parseDescribe() {
	date.parseNeedWork()
}

func (date *Date) parseNeedWork() {
	if date.IsIn(Holidays) {
		date.NeedWork = false
	} else if date.IsIn(Exchanges) {
		date.NeedWork = true
	} else if date.WeekNum() > 0 && date.WeekNum() < 6 {
		date.NeedWork = true
	} else {
		date.NeedWork = false
	}
}

func init() {
	Holidays = ReadFromFile("./data/holiday")
	Exchanges = ReadFromFile("./data/exchange")
}
