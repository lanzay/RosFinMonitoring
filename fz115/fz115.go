package fz115

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var Data FinMonitoring

func RunWS() {

	addr := ":83"
	mux := http.NewServeMux()
	mux.HandleFunc("/", Index)
	mux.HandleFunc("/GetByAll", GetByAll)
	mux.HandleFunc("/GetByFIO", GetByFIO)
	mux.HandleFunc("/GetByPassportN", GetByPassportN)

	log.Println("Start on ", addr)
	log.Println(http.ListenAndServe(addr, mux))

}

func LoadData() {

	//TODO Загрузка файла
	//TODO Выбор последнего
	Data.FileName = "27.09.2018.xml"
	f, err := os.Open("DATA\\fedsfm.ru\\" + Data.FileName)
	if err != nil {
		log.Println("[ERR]", err)
	}

	body, err := ioutil.ReadAll(f)
	if err != nil {
		log.Println("[ERR]", err)
	}

	//var data FinMonitoring
	err = xml.Unmarshal(body, &Data)
	if err != nil {
		log.Println("[ERR]", err)
	}
	body = nil

	for i := 0; i < len(Data.TerroristList); i++ {
		Data.TerroristList[i].Name = cutCDATA(Data.TerroristList[i].Name)
		Data.TerroristList[i].BirthDate = cutCDATA(Data.TerroristList[i].BirthDate)
		Data.TerroristList[i].Description = cutCDATA(Data.TerroristList[i].Description)
		Data.TerroristList[i].Addres = cutCDATA(Data.TerroristList[i].Addres)
		Data.TerroristList[i].TerrorisrResolution = cutCDATA(Data.TerroristList[i].TerrorisrResolution)
		Data.TerroristList[i].BirthPlase = cutCDATA(Data.TerroristList[i].BirthPlase)
		Data.TerroristList[i].PassportStr = cutCDATA(Data.TerroristList[i].PassportStr)
		Data.TerroristList[i].ID = Data.TerroristList[i].ID[1:]
		Data.TerroristList[i].Passport = passport(Data.TerroristList[i].PassportStr)
	}

	log.Printf("File %s, load %d items from %s", Data.FileName, len(Data.TerroristList), Data.XMLDate)
}

func cutCDATA(s string) string {

	res := strings.Replace(s, "\n<![CDATA[", "", -1)
	res = strings.Replace(res, "]]>", "", -1)
	res = strings.Replace(res, "* ", "", -1)
	res = strings.TrimSpace(res)
	return res
}

func passport(p string) Passport {

	//"ПАСПОРТ РФ: 9610 246803 ВЫДАН ОУФМС РОССИИ ПО ЧР В ЗАВОДСКОМ"
	res := Passport{}
	t1 := strings.Split(p, ":")
	if len(t1) <= 1 {
		return res
	}
	switch t1[0] {
	case "ПАСПОРТ РФ":
		res.Type = "ПАСПОРТ РФ"
		pass := strings.Split(t1[1], " ")
		if len(pass) < 2 {
			return res
		}
		res.NumSer = pass[1]
		res.NumNum = pass[2]
	}
	return res
}
