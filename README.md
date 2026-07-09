# วิธีติดตั้งและรันระบบ

### สร้างไฟล์ .env
สร้างไฟล์ .env ใน root project แบบ .env.sample โดยตัวอย่างสามารถคัดลอกข้อมูลจาก .env.example มารันได้เลย (local)
### Start Services
```bash
docker compose up -d
```

### Route

```text
http://localhost:8080
```


### Stop Services

```bash
docker compose down
```

# วิธีการสร้าง JWT และใช้งาน

ให้ทำการ create user ผ่าน

```text
[POST] http://localhost:8080/auth/register
```

หรือ

```text
[POST] http://localhost:8080/users
```

โดยจะต้องส่ง body เป็น JSON

```json
{
    "name": string,
    "password": string,
    "email": string
}
```

และนำ email กับ password ไปล็อคอินในระบบผ่าน

```text
[POST] http://localhost:8080/auth/login
```

เพื่อนำ token ไปใช้เพื่อขอข้อมูลต่างๆในระบบ โดยจะต้องระบุ body แบบนี้

```json
{
    "password": string,
    "email": string
}
```

หลังจากล็อคอินจะได้ token มากับ response ให้นำ token นั่นไปใส่ใน header ใน api ส่วนอื่นๆ

# ตัวอย่าง Request / Response

### ระบบจะเป็น CRUD api สำหรับผู้ใช้งาน โดยจากมี request และ response ดังนี้

#### ค้นหาผู้ใช้ทั้งหมดของระบบ
Request
```text
[GET] http://localhost:8080/users
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": [
        {
            "id": string,
            "name": string,
            "email": string,
            "created_at": date
        }
    ]
}
```

#### ค้นหาผู้ใช้ด้วย id
Request
```text
[GET] http://localhost:8080/users/:id
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": {
        "id": string,
        "name": string,
        "email": string,
        "created_at": date
    }
}
```

#### สร้างผู้ใช้
Request
```text
[POST] http://localhost:8080/users
```
หรือ
```text
[POST] http://localhost:8080/auth/register
```
Body
```json
{
    "name": string,
    "password": string,
    "email": string
}
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": {
        "id": string,
        "name": string,
        "email": string,
        "created_at": date
    }
}
```

#### แก้ไขผู้ใช้
Request
```text
[PUT] http://localhost:8080/users/:id
```
Body
```json
{
    "name": string,
    "password": string,
    "email": string
}
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": {
        "id": string,
        "name": string,
        "email": string,
        "created_at": date
    }
}
```

#### ลบผู้ใช้
Request
```text
[DELETE] http://localhost:8080/users/:id
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": {
        "message": string
    }
}
```

#### เข้าสู่ระบบ
Request
```text
[POST] http://localhost:8080/auth/login
```
Response
```json
{
    "code": number,
    "status": boolean,
    "data": {
        "token": string
    }
}
```

# Design Decision

ดีไซน์ตามโจทย์ที่ได้รับมอบหมาย โดยใช้ภาษา GO ในการพัฒนา และเลือกใช้ GIN เพราะเคยพัฒนาระบบโดยใช้ GIN มาก่อนเพื่อความรวดเร็วและสามารถตรวจสอบโค้ดได้อย่างมีประสิทธิภาพ


