package test


import(
	"net/http"
	"net/http/httptest"
	"testing"
    "encoding/json"

    "net/url"
    "strings"
//     "bytes"
//     "mime/multipart"

    "github.com/stretchr/testify/assert"
    "crud/router"
    "crud/user"
)

type response struct{
    ErrorCode string `json:"errorCode"`
    Message string `json:"message"`
    Data user.User
}



//     multipart/form-data 版本
//     var b bytes.Buffer
//     formWrite := multipart.NewWriter(&b)
//     formWrite := multipart.NewWriter(&b)
//     formWrite.WriteField("PersonId", "1")
//     formWrite.WriteField("Unix", "1729477530")
//     formWrite.WriteField("Token", "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c")
//     req.Header.Set("Content-Type", formWrite.FormDataContentType())
// assert.Equal(t, http.StatusOK, w.Code) //此方法不須使用 assert.New(t)


var UUIDGlobal string

func TestSetupRouter(t *testing.T) {
    RouterUserCreate(t)
    RouterUserRead(t)
    RouterUserUpdate(t)
    RouterUserDelete(t)
}

func RouterUserCreate(t *testing.T){

    var tests = []struct {
        name string
        path string
        fields user.User
        unix string
        token string
        expected string
    }{
        {
            name: "create is empty Unix",
            path: "/user/create",
            fields: user.User{
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "4",
        },
        {
            name: "create is empty token",
            path: "/user/create",
            fields: user.User{
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "",
            expected: "4",
        },
        {
            name: "create is successful",
            path: "/user/create",
            fields: user.User{
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "0",
        },
        {
            name: "create is reply email",
            path: "/user/create",
            fields: user.User{
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "2",
        },
    }

    a := assert.New(t)

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            requestValues := url.Values{}
            requestValues.Set("FirstName", tt.fields.FirstName)
            requestValues.Set("LastName", tt.fields.LastName)
            requestValues.Set("Gender", tt.fields.Gender)
            requestValues.Set("Email", tt.fields.Email)
            requestValues.Set("Address", tt.fields.Address)
            requestValues.Set("City", tt.fields.City)
            requestValues.Set("Unix", tt.unix)
            requestValues.Set("Token", tt.token)

            req, _ := http.NewRequest(http.MethodPost, tt.path, strings.NewReader(requestValues.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

            w := httptest.NewRecorder()
            router.SetupRouter().ServeHTTP(w, req)

            data := new(response)

            json.Unmarshal([]byte(w.Body.String()), &data)

            a.Equal(http.StatusOK, w.Code, "It is not http code 200")
            a.Equal(tt.expected, data.ErrorCode, data.Message)

            if data.ErrorCode == "0" {
                UUIDGlobal = data.Data.UUID
            }
        })
    }
}

func RouterUserRead(t *testing.T){

    var tests = []struct {
        name string
        path string
        fields user.User
        unix string
        token string
        expected string
    }{
        {
            name: "read is successful",
            path: "/user/read",
            fields: user.User{
                UUID: UUIDGlobal,
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "0",
        },
        {
            name: "read is empty Unix",
            path: "/user/read",
            fields: user.User{
                UUID: UUIDGlobal,
            },
            unix: "",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "4",
        },
    }

    a := assert.New(t)

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            requestValues := url.Values{}
            requestValues.Set("UUID", tt.fields.UUID)
            requestValues.Set("Unix", tt.unix)
            requestValues.Set("Token", tt.token)

            req, _ := http.NewRequest(http.MethodPost, tt.path, strings.NewReader(requestValues.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

            w := httptest.NewRecorder()
            router.SetupRouter().ServeHTTP(w, req)

            data := new(response)

            json.Unmarshal([]byte(w.Body.String()), &data)

            a.Equal(http.StatusOK, w.Code, "It is not http code 200")
            a.Equal(tt.expected, data.ErrorCode, data.Message)
        })
    }
}


func RouterUserUpdate(t *testing.T){

    var tests = []struct {
        name string
        path string
        fields user.User
        unix string
        token string
        expected string
    }{
        {
            name: "update is empty Unix",
            path: "/user/update",
            fields: user.User{
                UUID: UUIDGlobal,
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "4",
        },
        {
            name: "update is empty token",
            path: "/user/update",
            fields: user.User{
                UUID: UUIDGlobal,
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "",
            expected: "4",
        },
        {
            name: "update is successful",
            path: "/user/update",
            fields: user.User{
                UUID: UUIDGlobal,
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第一大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "0",
        },
        {
            name: "update is lost uuid",
            path: "/user/update",
            fields: user.User{
                UUID: "",
                FirstName: "王",
                LastName: "大民",
                Gender: "man",
                Email: "wang@mail.com",
                Address: "第二大道",
                City: "Taipei",
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "2",
        },
    }

    a := assert.New(t)

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            requestValues := url.Values{}
            requestValues.Set("UUID", tt.fields.UUID)
            requestValues.Set("FirstName", tt.fields.FirstName)
            requestValues.Set("LastName", tt.fields.LastName)
            requestValues.Set("Gender", tt.fields.Gender)
            requestValues.Set("Email", tt.fields.Email)
            requestValues.Set("Address", tt.fields.Address)
            requestValues.Set("City", tt.fields.City)
            requestValues.Set("Unix", tt.unix)
            requestValues.Set("Token", tt.token)

            req, _ := http.NewRequest(http.MethodPost, tt.path, strings.NewReader(requestValues.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

            w := httptest.NewRecorder()
            router.SetupRouter().ServeHTTP(w, req)

            data := new(response)

            json.Unmarshal([]byte(w.Body.String()), &data)

            a.Equal(http.StatusOK, w.Code, "It is not http code 200")
            a.Equal(tt.expected, data.ErrorCode, data.Message)
        })
    }

}

func RouterUserDelete(t *testing.T){

    var tests = []struct {
        name string
        path string
        fields user.User
        unix string
        token string
        expected string
    }{
        {
            name: "delete is empty Unix",
            path: "/user/delete",
            fields: user.User{
                UUID: UUIDGlobal,
            },
            unix: "",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "4",
        },
        {
            name: "delete is successful",
            path: "/user/delete",
            fields: user.User{
                UUID: UUIDGlobal,
            },
            unix: "1729477530",
            token: "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c",
            expected: "0",
        },
    }

    a := assert.New(t)

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            requestValues := url.Values{}
            requestValues.Set("UUID", tt.fields.UUID)
            requestValues.Set("Unix", tt.unix)
            requestValues.Set("Token", tt.token)

            req, _ := http.NewRequest(http.MethodPost, tt.path, strings.NewReader(requestValues.Encode()))
            req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

            w := httptest.NewRecorder()
            router.SetupRouter().ServeHTTP(w, req)

            data := new(response)

            json.Unmarshal([]byte(w.Body.String()), &data)

            a.Equal(http.StatusOK, w.Code, "It is not http code 200")
            a.Equal(tt.expected, data.ErrorCode, data.Message)
        })
    }
}


// 單一情境測試範例
// func TestSetupRouter(t *testing.T) {
//
//     a := assert.New(t)
// 	r := router.SetupRouter()
// 	w := httptest.NewRecorder()
//
//     // URL => /user/read
//     formData := url.Values{}
//     formData.Set("PersonId", "1")
//     formData.Set("Unix", "1729477530")
//     formData.Set("Token", "f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c")
//
// 	req, _ := http.NewRequest(http.MethodPost, "/user/read", strings.NewReader(formData.Encode()))
//
//     req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
//
// 	r.ServeHTTP(w, req)
//
//     var data response
//
//     json.Unmarshal([]byte(w.Body.String()), &data)
//
//     a.Equal(http.StatusOK, w.Code, "It is not http code 200")
// 	a.Equal("0", data.ErrorCode, "ErrorCode code is not 0")
// 	a.Equal("successful", data.Message, "Message is not successful")
//
// }