<!-- PROJECT LOGO -->
<div align="center">
  <a href="https://github.com/github_username/repo_name">
    <img src="https://media1.giphy.com/media/v1.Y2lkPTc5MGI3NjExdm9hMmlqMTZzODB2eng2NWhxenN0aHNsNXVybjhxanF3MWU5cXFyciZlcD12MV9pbnRlcm5hbF9naWZfYnlfaWQmY3Q9cw/bqtNyuyMwo0bdNk1gt/giphy.gif" alt="Logo" height="100">
  </a>

<h3 align="center">Gin Auth API Example</h3>

  <p align="center">
    RESTful API with JWT Token Authentication. Secure stateless authentication with access/refresh tokens, Redis session management, CSRF protection, and hashed credentials.
    <br />
    <br />
    <a href="https://github.com/darkfat123/gin-auth-api-example/issues">🚨 Report Bug</a>
    ·
    <a href="https://github.com/darkfat123/gin-auth-api-example/issues">✉️ Request Feature</a>
    .
    <a href="https://github.com/darkfat123/gin-auth-api-example?tab=readme-ov-file#-getting-started-for-development-only">🚀 Getting Started</a>
  </p>
</div>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">✨ Features:</h3>

  * Implemented session management using access and refresh tokens, following best practices for stateless authentication.
  * Integrated Redis to validate and manage token sessions efficiently.
  * Secured all /api endpoints by enforcing access token authentication
  * Applied password hashing and salting during user registration before storing credentials in the database
  * Implemented CSRF protection using CSRF tokens and securely managed both tokens with HTTP-only cookies

</br>
<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">🖥️ Programming languages and tools:</h3>

- Backend
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=go" />
  </a>
</p>

- Tools
<p align="left">
  <a href="https://skillicons.dev">
    <img src="https://skillicons.dev/icons?i=git,github,vscode,docker,redis,postgres" />
  </a>
</p>

<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

<h3 align="left">🖲️ Endpoint:</h3>

```bash
POST /api/auth/register
POST /api/auth/login
POST /api/auth/refresh
POST /api/auth/logout

GET /api/users/:id (Bearer Token is required)
```


<img src="https://i.imgur.com/dBaSKWF.gif" height="30" width="100%">

### 🚀 Getting Started (for development only)

#### 1. Clone the project
```bash
git clone https://github.com/darkfat123/gin-auth-api-example.git
cd gin-auth-api-example
```

#### 2. Backend
```bash
go mod tidy
go run main.go
```

#### 3. Environment Variables
```bash
DB_HOST={YOUR_DBHOST}
DB_USER={YOUR_DBUSER}
DB_PORT={YOUR_DBPORT}
DB_NAME={YOUR_DBNAME}
DB_PASSWORD={YOUR_DBPASSWORD}

REDIS_HOST={YOUR_REDIS_HOST}
REDIS_PORT={YOUR_REDIS_PORT}
REDIS_PASSWORD={YOUR_PASSWORD}

JWT_SECRET={YOUR_SECRETKEY}

ALLOWED_ORIGINS={YOUR_FRONTEND_URL}

REFRESH_TOKEN_MAX_AGE=86400 (refresh_token expires after 1 day)
```

#### 4. Setup Redis with Docker
```bash
docker pull redis
docker run -d --name redis -p {YOUR_REDIS_PORT}:{YOUR_REDIS_PORT} redis
```



<h3> Connect with me 🎊: <h3>
  <a href="https://www.linkedin.com/in/supakorn-yookack-39a730289/">
   <img align="left" alt="Supakorn Yookack | Linkedin" width="30px" src="https://www.vectorlogo.zone/logos/linkedin/linkedin-icon.svg" />
  </a>
  <a href="mailto:supakorn.yookack@gmail.com">
    <img align="left" alt="Supakorn Yookack | Gmail" width="32px" src="https://www.vectorlogo.zone/logos/gmail/gmail-icon.svg" />
  </a>
  <a href="https://medium.com/@yookack_s">
    <img align="left" alt="Supakorn Yookack | Medium" width="32px" src="https://www.vectorlogo.zone/logos/medium/medium-tile.svg" />
  </a>
   <a href="https://www.facebook.com/supakorn.yookaek/">
    <img align="left" alt="Supakorn Yookack | Facebook" width="32px" src="https://www.vectorlogo.zone/logos/facebook/facebook-tile.svg" />
  </a>
   <a href="https://github.com/darkfat123">
    <img align="left" alt="Supakorn Yookack | Github" width="32px" src="https://www.vectorlogo.zone/logos/github/github-tile.svg" />
  </a>
    <p align="right" > Created by <a href="https://github.com/darkfat123">darkfat</a></p> <p align="right" > <img src="https://komarev.com/ghpvc/?username=darkfat123&label=Profile%20views&color=0e75b6&style=flat" alt="darkfat123" /> </p>
