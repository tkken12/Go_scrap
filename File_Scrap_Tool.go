package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	Max_bw := os.Args[1]
	fopen(Max_bw)
}

func fopen(BW string) {

	parameter_to_int, err := strconv.Atoi(BW)
	if err != nil {
		log.Fatal(err)
	}

	var loop int = 10
	var str_read int = 0
	var tmp string
	var avg float64 = 0
	var tmp_str string
	var tmp_float float64
	var tmp_us_float float64
	var tmp_ms_float float64
	var tmp_sec_float float64

	//Read File
	for loop <= parameter_to_int {

		number_check := strconv.Itoa(loop)

		data, err := os.Open("qperf_test_" + number_check)

		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		scan := bufio.NewScanner(data)
		for scan.Scan() {
			tmp = scan.Text()
			tmp = strings.Replace(tmp, " ", "", -1)

			switch {
			case strings.Contains(tmp, "us") == true:
				tmp = strings.Trim(tmp, "latency")
				tmp = strings.Trim(tmp, "=")
				tmp = strings.Trim(tmp, "us")
				tmp_str = tmp
				tmp_float, err = strconv.ParseFloat(tmp_str, 64)
				tmp_us_float = (tmp_us_float + tmp_float) / 1000
				str_read = str_read + 1

			case strings.Contains(tmp, "ms") == true:
				tmp = strings.Trim(tmp, "latency")
				tmp = strings.Trim(tmp, "=")
				tmp = strings.Trim(tmp, "ms")
				tmp_str = tmp
				tmp_float, err = strconv.ParseFloat(tmp_str, 64)
				tmp_ms_float = tmp_ms_float + tmp_float
				str_read = str_read + 1

			case strings.Contains(tmp, "sec") == true:
				tmp = strings.Trim(tmp, "latency")
				tmp = strings.Trim(tmp, "=")
				tmp = strings.Trim(tmp, "sec")
				tmp_str = tmp
				tmp_float, err = strconv.ParseFloat(tmp_str, 64)
				tmp_float = tmp_float * 1000
				tmp_sec_float = (tmp_sec_float + tmp_float)
				str_read = str_read + 1
			}

			if str_read == 10 {
				avg = (tmp_us_float + tmp_ms_float + tmp_sec_float) / 10
				avg_str := strconv.FormatFloat(avg, 'f', 2, 64)
				tmp_us_float = 0
				tmp_ms_float = 0
				tmp_sec_float = 0
				loop_float := float64(loop)
				parameter_float := float64(parameter_to_int)
				band_per := float64(loop_float / parameter_float * 100)
				band_str := strconv.FormatFloat(band_per, 'f', 2, 64)

				value := map[string]string{
					"BW_per":  band_str,
					"Latency": avg_str,
				}

				value_json, _ := json.Marshal(value)
				fmt.Println(string(value_json))
				avg = 0
				str_read = 0
				break
			}
		}

		if loop > parameter_to_int {
			break
		}
		loop = loop + 10
	}
}
