package main


import (
	"fmt"
	"time"
	"log"
	"strconv"
	"strings"
	"encoding/csv"
	"os"
)

func main(){
    timeDiff()
}



type CompanyTime struct {
	workStart time.Time
	workOff   time.Time
	workOver  time.Time
}


func timeDiff() {

	const today string = "20180101"
	const tomorrow string = "20180102"
	workTime := CompanyTime{
		time.Date(2018, 01, 01, 8, 30, 0, 0, time.UTC),
		time.Date(2018, 01, 01, 17, 29, 59, 59, time.UTC),
		time.Date(2018, 01, 01, 17, 45, 0, 0, time.UTC),
	}


	records := [][]string{
		{"上班打卡", "下班打卡", "遲到", "早退", "加班", "是否有滿正常8小時"},
	}

	for {

		var startTime, endTime string = "", ""

		fmt.Println("輸入你的時間 格式(HH-mm HH-mm) .....")
		fmt.Scan(&startTime, &endTime)

		if startTime == "x"{
			csvFile, err := os.Create("test.csv")

			if err != nil {
            	log.Fatalln("建立csv檔錯誤 : ", err)
            }

			defer csvFile.Close()


			w := csv.NewWriter(csvFile)
			defer w.Flush()

            w.WriteAll(records)

			if writeErr := w.Error(); writeErr != nil {
				log.Fatalln("寫入錯誤 : ", writeErr)
			}

            break
        }


		// 格式錯誤
		if _, errStart := time.Parse("1504", startTime) ; errStart != nil {
			log.Printf("上班打卡時間-格式錯誤")
			continue
		} else if _, errEnd := time.Parse("1504", endTime) ; errEnd != nil{
			log.Printf("下班打卡時間-格式錯誤")
            continue
		}

		startTime = today+startTime
		if h, _ := strconv.ParseInt(endTime[:2], 10, 0) ; h < 6{
			endTime = tomorrow+endTime
		}else{
			endTime = today+endTime
		}

		checkIn, _ := time.Parse("200601021504", startTime)
		checkOut, _ := time.Parse("200601021504", endTime)

		var early, over, late string
		var isAllDay string = "N"

		fmt.Printf("上班 %s, 下班 %s \n", checkIn.Format("15:04:05"), checkOut.Format("15:04:05"))


		if checkIn.After(workTime.workStart) {
			late = timeFormat(workTime.workStart, checkIn)
			fmt.Printf("遲到 %s \n", late)
		}

		if checkOut.Before(workTime.workOff) {
	        early = timeFormat(checkOut,workTime.workOff)
	        fmt.Printf("早退 %s \n", early)
	    }else if checkOut.After(workTime.workOver) {
	        over = timeFormat(workTime.workOver, checkOut)
	        fmt.Printf("加班 %s \n", over)
	    }

		if checkIn.Before(workTime.workStart) && checkOut.After(workTime.workOff) && checkOut.Before(workTime.workOver)  {
			isAllDay = "Y"
			fmt.Printf("標準8小時 %s \n", isAllDay)
		}

		row := []string{
			checkIn.Format("15:04:05"),
			checkOut.Format("15:04:05"),
			late,
			early,
			over,
			isAllDay,
		}

		records = append(records, row)
	}
}


func timeFormat(timeA time.Time, timeB time.Time) string {

	overTemp, err := time.ParseDuration(timeB.Sub(timeA).String())
	if err != nil {
		return ""
	}

	res := strings.SplitAfter(overTemp.Round(time.Minute).String(), "m")

    return strings.Replace(strings.Replace(res[0], "h", "小時", 1), "m", "分", 1)
}