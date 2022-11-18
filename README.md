# Balance-Manager
<a id="readme-top"></a>

<br />
<div align="center">
  <a href="https://github.com/ssergomol/Realtime-Chat">
</a>

<h3 align="center">Real-time chat application</h3>
</div>





<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#technologies">Technologies</a></li>
      </ul>
    </li>
    <li><a href="#features-and-used-implementations">Features</a></li>
    <li>
      <a href="#installation">Installation</a>
      <ul>
        <li><a href="#server-setup">Server setup</a></li>
        <li><a href="#client-setup">Client setup</a></li>
      </ul>
    </li>
    <li><a href="#contributing">Contributing</a></li>
    <li><a href="#contact">Contact</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About the project

[![Product Name Screen Shot][product-screenshot]](https://example.com)

Service for managing users' balance which allows the money operations such as crediting funds, debiting funds, transferring funds from user to user as well as  obtaining the user's balance

<!-- TECHNOLOGIES -->
### Technologies

* [![GoLang-logo]][GoLang-url]
* [![React-logo]][React-url]
* [![PostgreSQL-logo]][PostgreSQL-url]
* [![JavaScript-logo]][JavaScript-url]

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- Features -->
## Features and implementations

* Real‐time messaging for multiple users (Used WebSocket protocol for connection, goroutines and channel for asynchronous operation)
* Authorization and authentication (By storing JSON Web Tokens in Cookies)
* Users' Database (Implemented using GORM ORM library and PostgreSQL)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Installation

### Server setup

1. Go version 1.16 or greater installed. To set this up, visit the https://go.dev/doc/install to install it for your operation system

2. In case you are setting up the server, it will be required to use PostgreSQL database server, visit https://www.postgresql.org/download/ and follow the installation instructions

3. Clone the repo
```sh
git clone https://github.com/ssergomol/Realtime-Chat
```

4. By default server tries to connect to database ```realtimechat``` using your system username. To change connection parametres, go through the ```backend/pkg/database/connection.go``` file

4. Install dependencies
```sh
go mod download
```

5. Start server
```sh
# cd backend/cmd
go run main.go
```

### Client setup

1. Install node package manager
```sh
npm install npm@latest -g
```

2. Start React server
```sh
# cd frontend
npm start
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>


<!-- CONTRIBUTING -->
## Contributing

If you have any intentions that would make this project better, fork the repo and create pull request

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<p align="right">(<a href="#readme-top">back to top</a>)</p>



<!-- CONTACT -->
## Contact

Sergey Molchanov - @ssergomol - ssergomoll@gmail.com

Project Link: [https://github.com/ssergomol/Realtime-Chat](https://github.com/ssergomol/Realtime-Chat)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

[React-logo]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[GoLang-url]: https://go.dev
[GoLang-logo]: https://img.shields.io/badge/GoLang-ffffff?style=for-the-badge&logo=Go&logoColor=7bccec
[product-screenshot]: images/home_page.png
[PostgreSQL-url]: https://www.postgresql.org/
[PostgreSQL-logo]: https://img.shields.io/badge/PostgreSQL-ffffff?style=for-the-badge&logo=PostgreSQL&logoColor=008bb9
[JavaScript-url]: https://javascript.com
[JavaScript-logo]: https://img.shields.io/badge/JavaScript-323330?style=for-the-badge&logo=javascript&logoColor=f0db4f
