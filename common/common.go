package common

import (
	"strconv"
	"strings"
	"time"

	"github.com/satori/go.uuid"
)

const RFC3339Nano = "2006-01-02 15:04:05"

type Time time.Time

func BuilderPrimaryid() string {
	return strings.Replace(uuid.NewV4().String(), "-", "", -1)
}

func TableName(t string) string {
	return "m_" + t
}

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+RFC3339Nano+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(RFC3339Nano)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, RFC3339Nano)
	b = append(b, '"')
	return b, nil
}

func Now() Time {
	return Time(time.Now())
}

func Split(str string, cut string) []string {
	if str == "" {
		return []string{}
	}
	return strings.Split(str, cut)
}

func SplitInt(str string, cut string) []int {
	if str == "" {
		return []int{}
	}

	strArr := strings.Split(str, cut)
	var intArr []int
	for _, v := range strArr {
		int, _ := strconv.Atoi(v)
		intArr = append(intArr, int)
	}

	return intArr
}
