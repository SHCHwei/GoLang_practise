## 聯絡我們API 

*data.json 存放輸入資料

### 新增

---
URL **[POST]** localhost:8001/create

| 參數欄位      |  參數key  | 型態  |
| -------------|:-------:| -----:|
| 姓名          |  Name   | string |
| 內容          | Content | string(50) |
| 手機          |  Phone  | string(10) |
| 信箱          |  Email  | string |

### 讀取

---
URL **[POST]** localhost:8001/read

| 參數欄位      |參數key        | 型態  |
|-----------|:-------------:| -----:|
| no parameter |


### 編輯

---
URL **[POST]** localhost:8001/update

| 參數欄位 |  參數key  |         型態 |
|------|:-------:|-----------:|
| ID   |   ID    |        INT | 
| 姓名   |  Name   |     string |
| 內容   | Content | string(50) |
| 手機   |  Phone  | string(10) |
| 信箱   |  Email  |     string |


### 刪除

---
URL **[POST]** localhost:8001/delete

| 參數欄位      |參數key        | 型態  |
|-----------|:-------------:| -----:|
| ID   |   ID    |        INT | 
