package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	log.SetFlags(0)          //จะไม่เอาดีเทลต่างๆเข้ามา
	log.Println(os.Getpid()) //ก่อนโปรแกรมนี้จะรันเราจะlog process id ตัวนี้ออกมา
	http.ListenAndServe("", nil)
}

/*

This site can't be reached เราอยากรู้ว่าตรงนี้รันที่พอร์ทไหนสามารถทำได้ง่ายๆ แต่ก่อนอื่น
ก่อนที่จะรันโปรแกรมนี้เราอยากให้โปรแกรมนี้ log process id ออกมาก่อน
พอรันแล้ว เราจะได้ pid 544(เลขเปลี่ยนได้) เลข pid เราสามารถเอาไปหาได้ว่ามันรันบนพอร์มไหน
ถ้าเกิดว่าลง ConEmu เราสามารถใช้คำสั่งนี้ได้ netstat ถ้ายังไม่มีก็ลงจาก cygwin ได้
comand ที่จะใช้คือ netstat -ano | grep 544 กด enterg เข้าไปจะเห็นว่า

TCP		0.0.0.0:80	0.0.0.0:0	LISTENING 12544
TCP 	[::]:80		[::]:0		LISTENING 12544

web server ถูก run บน TCP protocal และ run ที่ address by 0.0.0.0(ทุกพอร์ทเลย)
ซึ่งเราสามารถเข้าถึงได้จากเครื่องของเราโดยใช้ localhost แล้วพอร์ท run อยู่คือ 80
เราก็ไป check ว่าพอร์ท 80 ของเราทำงานได้จริงรึเปล่า localhost:80 / localhost
ถ้าพอร์ท 80 ไม่ต้องระบุพอร์ทก็ได้ ใส่แค่ localhost เฉยๆก็ได้ พอรันเราจะได้
404 page not found เพราะเรายังไม่ได้ใส่อะไรเข้าไปเลย

ทำไมเราพาร์ท " "(emtpy string)เข้าไป แล้วมันต้องเป็นพอร์ท 80 ....
ให้เข้าไปดูที่ ListenAndServe -->
 if addr == "" -->  addr = ":http" --> net.Listen


http.ListenAndServe(":8082", nil)
$ netstat -ano | grep 1556
  TCP    0.0.0.0:8082           0.0.0.0:0              LISTENING       1556
  TCP    [::]:8082              [::]:0                 LISTENING       155


*/
