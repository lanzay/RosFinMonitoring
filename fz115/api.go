package fz115

import (
	"encoding/json"
	"net/http"
	"strings"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("/GetByAll?FIO=\"\"&PASSPORT_SER=\"\"&PASSPORT_NUM=\"\"&BD=\"\""))
}

func respond(w http.ResponseWriter, r *http.Request, res []Terrorist) {

	if len(res) > 0 {
		body, _ := json.Marshal(res)
		w.WriteHeader(http.StatusFound)
		w.Write(body)
	} else {
		w.WriteHeader(http.StatusOK)
	}

}

func GetByAll(w http.ResponseWriter, r *http.Request) {

	res := []Terrorist{}
	r.ParseForm()
	fio := strings.ToUpper(r.FormValue("FIO"))
	passSer := strings.ToUpper(r.FormValue("PASSPORT_SER"))
	passNum := strings.ToUpper(r.FormValue("PASSPORT_NUM"))
	bd := strings.ToUpper(r.FormValue("BD"))

	for i := 0; i < len(Data.TerroristList); i++ {
		if strings.EqualFold(Data.TerroristList[i].Passport.NumSer, passSer) &&
			strings.EqualFold(Data.TerroristList[i].Passport.NumNum, passNum) &&
			strings.EqualFold(Data.TerroristList[i].Name, fio) &&
			strings.EqualFold(Data.TerroristList[i].BirthDate, bd) {
			res = append(res, Data.TerroristList[i])
		}
	}
	respond(w, r, res)
}

func GetByFIO(w http.ResponseWriter, r *http.Request) {

	res := []Terrorist{}
	r.ParseForm()
	fio := r.FormValue("FIO")

	for i := 0; i < len(Data.TerroristList); i++ {
		if strings.EqualFold(Data.TerroristList[i].Name, fio) {
			res = append(res, Data.TerroristList[i])
		}
	}
	respond(w, r, res)
}

func GetByPassportN(w http.ResponseWriter, r *http.Request) {

	res := []Terrorist{}
	r.ParseForm()
	passSer := strings.TrimSpace(r.FormValue("PASSPORT_SER"))
	passNum := strings.TrimSpace(r.FormValue("PASSPORT_NUM"))

	if len(passSer+passNum) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	for i := 0; i < len(Data.TerroristList); i++ {
		if strings.EqualFold(Data.TerroristList[i].Passport.NumSer, passSer) &&
			strings.EqualFold(Data.TerroristList[i].Passport.NumNum, passNum) {
			res = append(res, Data.TerroristList[i])
		}
	}
	respond(w, r, res)
}
