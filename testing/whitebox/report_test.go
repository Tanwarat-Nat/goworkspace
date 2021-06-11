package report

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicReport(t *testing.T) {
	getReport = func(id string) string {
		//connect database หรือจะไปเรียก function ที่ connect database อีกทีก้ได้
		return "xxxx"
	} //มันไม่ใช่สร้างตัวใหม่(:=) มันเป็นการ reassign (=) ใหม่

	report := generateReport("12DB")
	if !strings.Contains(report, ": VERIFIED") {
		t.Error("failed")
	}
	fmt.Println("got report : ", report)
}

/*
เวลาเราเทสแบบนี้มันจะไปเรียก database จริงๆ เราไม่ต้องการให้มันไปเรียก database
เราต้องการจะแทนที่ getReport() ตรงนี้จะทำยังไง

- over write โดย coppy มันมาวางในไฟล์เทส ก็ไม่ได้ มันฟ้องว่า redrclared ประกาศซ้ำกัน
- anonymous fun โดยตั้ง getReport := func() ก็ไม่ได้ มันฟ้องว่า declar but not used
แล้วถ้าเราใช้ getReport("sdf")  แล้วให้ใน func return "xxx" พอลองรันดูก็ยังไม่ได้
generateReport l6fmhkpก็ไปเรียก getReport เดิมในไฟล์ report.go

เราจะทำได้ยังไง จะ subtitute เราจะแทนที่ getReport() ในไฟล์ report.go ได้ยังไง
วิธีง่ายๆเลยคือเราจะทำกาา refactor ตรงนั้นออกมา เขาถึงเรียกว่า whitebox

คำว่า  whitebox ก็คือเราสามารถเห็นว่าข้างในคืออะไร อย่างที่นี้เราเห็น generateReport()
มันเขียนยังไง


*/
