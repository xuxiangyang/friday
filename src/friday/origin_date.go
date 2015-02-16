package friday

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type OriginDate struct {
	Year  int `json:year`
	Month int `json:month`
	Day   int `json:day`
}

type OriginDates map[string]OriginDate

func ReadFromFile(filePath string) (originDates OriginDates) {
	originDates = make(OriginDates)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	regex := regexp.MustCompile(`\d+\s+\d+\s+\d+`)
	for _, line := range regex.FindAllString(string(data), -1) {
		originDate, err := convertOriginDate(line)
		if err != nil {
			log.Fatal(err)
			continue
		}

		originDates[originDate.String()] = originDate
	}
	return originDates
}

func convertOriginDate(line string) (originDate OriginDate, err error) {
	ymd := strings.Split(line, " ")

	if len(ymd) != 3 {
		return originDate, errors.New("Bad Date Data with" + line)
	}

	yearStr, monthStr, dayStr := ymd[0], ymd[1], ymd[2]

	if year, err := strconv.Atoi(yearStr); err != nil {
		return originDate, errors.New("Bad Date Data with" + line)
	} else {
		originDate.Year = year
	}

	if month, err := strconv.Atoi(monthStr); err != nil && month > 0 && month < 12 {
		return originDate, errors.New("Bad Date Data with" + line)
	} else {
		originDate.Month = month
	}

	if day, err := strconv.Atoi(dayStr); err != nil && day > 0 && day < 32 {
		return originDate, errors.New("Bad Date Data with" + line)
	} else {
		originDate.Day = day
	}

	return originDate, nil
}

func (originDate OriginDate) String() string {
	return strconv.Itoa(originDate.Year) + strconv.Itoa(originDate.Month) + strconv.Itoa(originDate.Day)
}

func (a *OriginDate) Equal(b *OriginDate) bool {
	return a.Year == b.Year && a.Month == b.Month && a.Day == b.Day
}

func (originDate *OriginDate) IsIn(originDates OriginDates) bool {
	knownOriginDate, exists := originDates[originDate.String()]
	if exists && knownOriginDate.Equal(originDate) {
		return true
	}
	return false
}

//sunday is 0
func (originDate *OriginDate) WeekNum() int {
	return int(time.Date(originDate.Year, time.Month(originDate.Month), originDate.Day, 0, 0, 0, 0, time.UTC).Weekday())
}
