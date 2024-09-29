# content us

建立簡易填入表單 API，MongoDB存放

---

### 查詢 API

URL **[GET]** localhost:8002/api/user/search

| 參數欄位 |  參數key  |  型態 |  範例 |
|------|:-------:|-------:|-------:|
| 開始日期 |  startDate   | string | 2023-09-06T13:40:17.862Z |
| 結束日期 | endDate | string | 2023-09-13T23:59:17.862Z |


----

## 新增 API

URL **[POST]** localhost:8002/api/user/create


| 參數欄位 |  參數key  |  型態 |  範例 |
|------|:-------:|-------:|-------:|
| 姓名   |  Name   | string | 王小明 |
| 主內容  | Content | string | 內容 |
| 電話   | Phone | string | 0911111111 |
| 信箱   | Email | string | human@mai.com |