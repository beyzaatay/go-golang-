package main

import (
	"fmt"
	"time"
)

func main() {

	// şu an ki zaman
	zaman := time.Now()
	fmt.Println(zaman)

	fmt.Println(zaman.Day())
	fmt.Println(zaman.Month())

	fmt.Println()

	fmt.Println(zaman.Minute())
	fmt.Println(zaman.Hour())

	zamanstring := zaman.String()
	fmt.Println(zamanstring)

	//sleep -> bekleme süresi

	fmt.Println("merhaba")
	time.Sleep(time.Second * 5)
	fmt.Println("merhaba")

	//UTC -> evrensel saat dilimi

	zamanUTC := time.Now().UTC()
	fmt.Println(zamanUTC)

	// Date -> kendi tarihimizi oluşturabiliriz

	dogumgunu := time.Date(2003, 5, 12, 12, 30, 0, 0, time.Local)

	fmt.Println(dogumgunu)
	fmt.Printf("%v-%v-%v", dogumgunu.Day(), dogumgunu.Month(), dogumgunu.Year())

	baslangic := time.Date(1923, 10, 29, 0, 0, 0, 0, time.Local)
	bugun := time.Now().UTC()

	fark := bugun.Sub(baslangic).Hours() / 24
	fmt.Println(fark)

	var d time.Duration = 1000000000

	fmt.Println(d.Hours())

	// tarih ekleme

	suankizaman := time.Now()

	yarin := suankizaman.AddDate(0, 0, 1)
	fmt.Println(yarin)

	ikigunsonra := suankizaman.Add(time.Hour * 48)
	fmt.Println(ikigunsonra)

	// kendimize özel zaman biçimi

	hafta := time.Hour * 7 * 24
	fmt.Println(hafta)
}
