# Gin CRUD API

* 建立 CRUD API
* GIN middleware
* Mariadb連接
* Config loading File
* Router
* 測試功能

---

### 新增 API

URL **[POST]** localhost:8004/user/create

| 參數欄位  |   參數key    |  型態 |                                   範例 |
|-------|:----------:|-------:|-------------------------------------:|
| 姓     | FirstName  | string |                                    王 |
| 名     |  LastName  | string |                                   大民 |
| 性別    |   Gender   | string |                                  man |
| 信箱    |   Email    | string |                        wang@mail.com |
| 地址    |  Address   | string |                                 第一大道 |
| 城市    |    City    | string |                               Taipei |
| 時間    |    Unix    | string |                      Timestamp(ex: 1729477530) |
| TOKEN |    Token    | string |f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c|


### 查詢 API

URL **[POST]** localhost:8004/user/read

| 參數欄位  |   參數key    |  型態 |                                   範例 |
|-------|:----------:|-------:|-------------------------------------:|
| UUID  |    UUID    | string | 065c9e21-19ef-4e9e-b257-43ccbd5be369 |
| 時間    |    Unix    | string |                      Timestamp(ex: 1729477530) |
| TOKEN |    Token    | string |f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c|


### 更新 API

URL **[POST]** localhost:8004/user/update

| 參數欄位  |   參數key    |  型態 |                                   範例 |
|-------|:----------:|-------:|-------------------------------------:|
| UUID  |    UUID    | string | 065c9e21-19ef-4e9e-b257-43ccbd5be369 |
| 姓     | FirstName  | string |                                    王 |
| 名     |  LastName  | string |                                   大民 |
| 性別    |   Gender   | string |                                  man |
| 信箱    |   Email    | string |                        wang@mail.com |
| 地址    |  Address   | string |                                 第一大道 |
| 城市    |    City    | string |                               Taipei |
| 時間    |    Unix    | string |                      Timestamp(ex: 1729477530) |
| TOKEN |    Token    | string |f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c|


### 刪除 API

URL **[POST]** localhost:8004/user/delete

| 參數欄位  |   參數key    |  型態 |                                   範例 |
|-------|:----------:|-------:|-------------------------------------:|
| UUID  |    UUID    | string | 065c9e21-19ef-4e9e-b257-43ccbd5be369 |
| 時間    |    Unix    | string |                      Timestamp(ex: 1729477530) |
| TOKEN |    Token    | string |f4785491bb0c9fd444b05fbc1ee110cb4962cdc733a9b62ed4bb4351f00f037c|
