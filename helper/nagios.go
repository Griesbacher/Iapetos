package helper

import (
	"regexp"
	"strings"

	"fmt"
	"strconv"

	"github.com/griesbacher/Iapetos/logging"
)

var (
	checkMulitRegex       = regexp.MustCompile(`^(.*::)(.*)`)
	regexPerformancelable = regexp.MustCompile(`([^=]+)=(U|[\d\.,\-]+)([\w\/%]*);?([\d\.,\-:~@]+)?;?([\d\.,\-:~@]+)?;?([\d\.,\-]+)?;?([\d\.,\-]+)?;?\s*`)
)

type PerformanceData struct {
	Label string
	Unit  string
	Data  map[string]float64
}

func IteratePerformanceData(input string) <-chan PerformanceData {
	ch := make(chan PerformanceData)
	go func() {
		perfSlice := regexPerformancelable.FindAllStringSubmatch(input, -1)
		currentCheckMultiLabel := ""
		//try to find a check_multi prefix
		if len(perfSlice) > 0 && len(perfSlice[0]) > 1 {
			currentCheckMultiLabel = getCheckMultiRegexMatch(perfSlice[0][1])
		}
	OuterLoop:
		for _, value := range perfSlice {
			perf := PerformanceData{
				Label: value[1],
				Unit:  value[3],
				Data:  map[string]float64{},
			}
			if currentCheckMultiLabel != "" {
				//if an check_multi prefix was found last time
				//test if the current one has also one
				if potentialNextOne := getCheckMultiRegexMatch(perf.Label); potentialNextOne == "" {
					// if not put the last one in front the current
					perf.Label = currentCheckMultiLabel + perf.Label
				} else {
					// else remember the current prefix for the next one
					currentCheckMultiLabel = potentialNextOne
				}
			}
			for i, data := range value {
				if data == "" {
					continue
				}
				data = strings.Replace(data, ",", ".", -1)
				var err error
				switch i {
				case 2:
					perf.Data["Value"], err = strconv.ParseFloat(data, 64)
				case 4:
					perf.Data["Warning"], _ = strconv.ParseFloat(data, 64)
				case 5:
					perf.Data["Critical"], _ = strconv.ParseFloat(data, 64)
				case 6:
					perf.Data["Minimum"], err = strconv.ParseFloat(data, 64)
				case 7:
					perf.Data["Maximum"], err = strconv.ParseFloat(data, 64)
				}
				if err != nil {
					logging.Flog("Skipping perfdata due to parse erros. Index: %d Data: %s\n", i, fmt.Sprint(value))
					continue OuterLoop
				}
			}
			ch <- perf
		}
		close(ch)
	}()
	return ch
}

func getCheckMultiRegexMatch(perfData string) string {
	regexResult := checkMulitRegex.FindAllStringSubmatch(perfData, -1)
	if len(regexResult) == 1 && len(regexResult[0]) == 3 {
		return regexResult[0][1]
	}
	return ""
}
