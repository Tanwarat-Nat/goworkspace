package report

var getReport = func(id string) string {
	//connect database หรือจะไปเรียก function ที่ connect database อีกทีก้ได้
	return "abcdef"
}

func generateReport(id string) string { //เอา id นี้ไป query ใน database
	data := getReport(id)        // ก่อนเอาข้อมูลมาต้องไปเรียกแล้ว getreport ก็จะไปต่อกับ internet
	return data + " : VERIFIED." // ให้ดูเหมือนกับว่า report ตรงนี้ถูก generate ไปเรียบร้อยเเล้ว
}

/*
คลิปที่ผ่านมาที่เทสกันเขาเรียกว่า black box testing หมายความว่า เราเทสโดยเราไม่ไปแตะต้อง
โค้ดที่อยู่ข้างใน เหมือนกับว่าเราไม่ใช่ owner ของ package เราเทส api ของเขาเฉยๆว่า
ส่งอะไรไปแล้วได้อะไรรีเทิร์นกลับมา

white box testing มีประโยชน์มากที่ว่า ในเคสที่มี implementation ที่ค่อนข้างซับซ้อน
เราเป็นคนเขียนตรงนั้นเอง เราเป็นเจ้าของโค้ดเอง อย่างการต่อ connection ทุกครั้งเวลาเราเทส
เราไม่ต้องการต่อ database บางทีเราไม่อยากจะต่อ restapi โดยตรงเราสามารถ mock ตรงนั้นได้


*/
