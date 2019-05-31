package thp

import (
	"strconv"
	"strings"
)

type SensorData struct {
	Date     string  `json:"date"`
	Thermal  float64 `json:"thermal"`
	Humidity float64 `json:"humidity"`
	Pressure float64 `json:"pressure"`
}

func ParseJson(in string) []SensorData {
	s := strings.Replace(in, " ", "", -1)
	s = strings.Replace(s, "\t", "", -1)

	/* remove first [ and last ] */
	d := ""
	for k, v := range strings.Split(s, "\n") {
		if k != 0 && k < len(strings.Split(s, "\n"))-1 {
			d += v
		}
	}

	sd := []SensorData{}
	for _, v := range strings.Split(d, ",[") {
		y := strings.Replace(v, "[", "", -1)
		y = strings.Replace(y, "]", "", -1)
		y = strings.Replace(y, "\"", "", -1)
		q := strings.Split(y, ",")
		tsd := SensorData{}
		for qk, qv := range q {
			switch qk {
			case 0:
				tsd.Date = qv
			case 1:
				f, _ := strconv.ParseFloat(qv, 64)
				tsd.Thermal = f
			case 2:
				f, _ := strconv.ParseFloat(qv, 64)
				tsd.Humidity = f
			case 3:
				f, _ := strconv.ParseFloat(qv, 64)
				tsd.Pressure = f
			}
		}
		sd = append(sd, tsd)
	}
	return sd
}
