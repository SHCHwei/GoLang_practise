package main

import (
    "log"
    "os"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "strconv"
    "time"
    "github.com/go-playground/validator/v10"
)


type ContentData struct {
	ID int `json:"id"`
	Name string `json:"name" validate:"required"`
	Content string `json:"content" validate:"required,max=20,min=1"`
	Phone string `json:"phone" validate:"required"`
	Email string `json:"email" validate:"required,email"`
	Time int64 `json:"time"`
}


type message struct{
    Statue string `json:"statue"`
    Message string `json:"message"`
    Err error `json:"error"`
}


var validate *validator.Validate
var response message

var Rows []ContentData

func main(){

    validate = validator.New()

    http.HandleFunc("/create", create)
    http.HandleFunc("/read", read)
    http.HandleFunc("/update", update)
    http.HandleFunc("/delete", delete)

    logF("Server exec ", http.ListenAndServe(":8001", nil))
}



func Get() bool {

    jsonFile, err := os.OpenFile("data.json", os.O_CREATE, 0777)

    if err != nil {
        logF("JSON FILE error", err)
        return false
    }

    defer jsonFile.Close()

    dataList, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(dataList, &Rows)

    if len(Rows) == 0 {
        return false
    }

    return true
}


func create(w http.ResponseWriter, r *http.Request) {

    if r.Method == "POST" {

        id := 1

        if Get() {
            id = Rows[len(Rows) - 1].ID + 1
        }

        newRow := ContentData{
            ID: id,
            Name: r.FormValue("Name"),
            Content: r.FormValue("Content"),
            Phone: r.FormValue("Phone"),
            Email: r.FormValue("Email"),
            Time: time.Now().Unix(),
        }

        if err := validate.Struct(newRow); err != nil {
            response.Message = "validate error"
            response.Err = err
            response.writeMessage(w)

            return
        }

        response.Message = ""
        response.Err = enterFile(append(Rows, newRow))

    } else {
         response.Message = "Method not allow"
    }

    response.writeMessage(w)
}

func read(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        Get()
        jsonResp, _ := json.Marshal(Rows)
        w.Write(jsonResp)
    } else {
        response.Message = "Method not allow"
        response.writeMessage(w)
    }
}

func update(w http.ResponseWriter, r *http.Request) {

    if r.Method == "POST" {

        if Get() == false {
            response.Message = "Data null"
            response.writeMessage(w)
            return
        }

        id, _ := strconv.Atoi(r.FormValue("ID"))

        newRow := ContentData{
            ID: id,
            Name: r.FormValue("Name"),
            Content: r.FormValue("Content"),
            Phone: r.FormValue("Phone"),
            Email: r.FormValue("Email"),
            Time: time.Now().Unix(),
        }

        if err := validate.Struct(newRow); err != nil {
            response.Message = "validate error"
            response.Err = err
            response.writeMessage(w)

            return
        }

        for k, v := range Rows{
            if v.ID == newRow.ID {
                Rows[k] = newRow
            }
        }

        response.Message = ""
        response.Err = enterFile(Rows)

    } else {
        response.Message = "Method not allow"
    }

    response.writeMessage(w)
}

func delete(w http.ResponseWriter, r *http.Request) {

    if r.Method == "POST" {

        if Get() {
            var newRow []ContentData

            for _, v := range Rows{
                id := strconv.Itoa(v.ID)
                if id != r.FormValue("ID") {
                    newRow = append(newRow, v)
                }
            }

            response.Message = ""
            response.Err = enterFile(newRow)
        }

    } else {
        response.Message = "Method not allow"
    }

    response.writeMessage(w)
}


func(m *message) writeMessage(w http.ResponseWriter){

    if m.Err != nil {
        m.Statue = "false"
        m.Message = "失敗"
    }else if m.Message != ""{
        m.Statue = "false"
    }else{
        m.Statue = "true"
    }

    jsonResp, _ := json.Marshal(m)
    w.Write(jsonResp)
}


func enterFile(rows []ContentData) error {

    newTeamBytes, err := json.Marshal(rows)

    if err != nil {
        return err
    }

    return ioutil.WriteFile("data.json", newTeamBytes, 0644)

}


func logF(s string , err error){
    if err != nil {
		log.Fatal(s, err.Error())
	}else{
        log.Println(s)
    }
}
