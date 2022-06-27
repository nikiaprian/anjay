<br />
<p align="center">
  <a href="https://github.com/rg-km/final-project-engineering-15">
    <img src="./frontend/src/Assets/logo2-01.png" alt="Logo" width="200" height="100">
  </a>  
  
  <h3 align="center">Find Your Solution Here!</h3>

  <p align="center">
    Semakin bertumbuhnya teknologi, maka semakin bertambah pengguna dari teknologi tersebut, juga disertai dengan kondisi dimana tidak memungkinkan kita dalam melakukan interaksi secara langsung terkhususnya untuk bertanya dalam dunia pemrograman oleh sebab itu  kami melakukan pembuatan website yang bernama <strong> Codein </strong>. <br /> 
    <br />
    <strong> CodeIn </strong> merupakan platform yang dibuat untuk memberikan kemudahan bagi Programmer Indonesia untuk berinteraksi menyelesaikan permasalahan ataupun memberikan informasi yang bermanfaat bagi programmer lainnya
    <br />
    <a href="https://docs.google.com/presentation/d/1zMoOK3sKTzY6oitYUy-qlqJskLaAwqnW/edit?usp=sharing&ouid=115204411593671225594&rtpof=true&sd=true"><strong>Explore the docs »</strong></a>
    <br />
    <br />
  </p>
</p>

## Table of Contents

| [Tech Stack](#tech-stack) | [Getting Started](#getting-started) | [Wesbite Page](#website-page) | [Mockup Design](#mockup-design) | [Entity Relationship Diagram](#entity-relationship-diagram) | [Architectural Pattern](#architectural-pattern) | [Swagger OpenAPI Docs](#swagger-openapi-docs) | [Server URL](#server-url) |
| :-----------------------: | :---------------------------------: | :---------------------------: | :-----------------------------: | :---------------------------------------------------------: | :---------------------------------------------: | --------------------------------------------- | ------------------------- |


## Tech Stack

[`^ kembali ke atas ^`](#table-of-contents)

### Frontend Engineering

- **Language:** [Javascript](https://www.javascript.com/)
- **Framework:** [ReactJs](https://reactjs.org/)
- **Styling:** [TailwindCSS](https://tailwindcss.com/)
- **State Management:** [Zuztand](https://github.com/pmndrs/zustand)
- **Size App:** [Web and Mobile](https://github.com/rg-km/final-project-engineering-15)
- **Deployment:** [EC2](https://aws.amazon.com/ec2/)
- **Object Cloud:** [S3](https://aws.amazon.com/s3/)
- **Code Editor:** [Visual Studio Code](https://code.visualstudio.com/)

### Backend Engineering

- **Language:** [Go](https://golang.org/)
- **Framework:** [Gin](https://gin-gonic.com/)
- **Database:** [SQLite](https://www.sqlite.org), [RDS](https://aws.amazon.com/rds/)
- **Unit Testing:** [Ginkgo](https://github.com/onsi/ginkgo)
- **API Testing:** [Postman](https://www.getpostman.com/)
- **API Docs:** [Swagger](https://swagger.io/)
- **Deployment:** [EC2](https://aws.amazon.com/ec2/)
- **Object Cloud:** [S3](https://aws.amazon.com/s3/)
- **Code Editor:** [Visual Studio Code](https://code.visualstudio.com/)

## Getting Started

[`^ kembali ke atas ^`](#table-of-contents)

### Frontend

1. Clone Repository
```sh
git clone https://github.com/rg-km/final-project-engineering-15.git
```
2. Go to Frontend Folder
```sh
cd frontend
```
3. Install Module Packages
```sh
npm install
``` 
4. Run Code
```sh
npm start
```

### Backend

1. Clone Repository
  ```sh
  git clone https://github.com/rg-km/final-project-engineering-15.git
  ```
2. Go to Backend Folder
```sh
cd backend
```
3. Build Module
```sh
go get
```
4. Migrate Database
```sh
sql-migrate up
```
5. Run Code
```sh
go run main.go
```

## Website Page

[`^ kembali ke atas ^`](#table-of-contents)

- **Home Page**
- **Forum Page**
  - Create Forum Page
  - Detail Forum Page
- **Blog Page**
  - Create Blog Page
  - Detail Blog Page
- **Dashboard Profile Page**
- **FAQ Page**
- **About Page**

## Mockup Design

[`^ kembali ke atas ^`](#table-of-contents)

https://www.figma.com/file/YggOu5kNAMGF7tLF3hNeIW/Design-RG?node-id=70%3A11867

## Entity Relationship Diagram

[`^ kembali ke atas ^`](#table-of-contents)

![ERD](./frontend/src/Assets/CodeIn-ERD.png)

## Architectural Pattern

[`^ kembali ke atas ^`](#table-of-contents)

Architectural pattern yang digunakan adalah Clean Architecture, dimana aplikasi terbagi atas 4 layer, antara lain: **Domain/Entity**, **Use Case**, **Controller**, dan **Repository**. Dengan pola seperti ini, semua komponen aplikasi dapat dibuat secara independen sehingga mengurangi dependensi antar komponen dan dapat dikembangkan secara berkelanjutan.

![Clean Architecture](./frontend/src/Assets/CleanArch.png)

## Swagger OpenAPI Docs

[`^ kembali ke atas ^`](#table-of-contents)

https://app.swaggerhub.com/apis/rchmatagung/CodeIn/1.0.0

## Server URL

[`^ kembali ke atas ^`](#table-of-contents)
### Frontend

https://fe.codein.studio/

### Backend
https://be.codein.studio/
