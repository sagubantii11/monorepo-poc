# Monorepo Setup with NestJS and Golang

## 🛠️ Initial Setup

1.  **Create Monorepo Directory and Initialize Git:**

    ```bash
    mkdir monorepo-poc && cd monorepo-poc
    git init
    touch .gitignore
    ```

    * `mkdir monorepo-poc`: Creates a new directory named `monorepo-poc`.
    * `cd monorepo-poc`: Navigates into the newly created directory.
    * `git init`: Initializes a new Git repository in the current directory.
    * `touch .gitignore`: Creates an empty `.gitignore` file for specifying files Git should ignore.

2.  **Folder Structure:**

    ```
    monorepo-poc/
    ├── README.md
    ├── build-compose.sh
    ├── run-compose.sh
    ├── docker-compose.yml
    ├── backend/
    │   ├── nest-basic-poc/
    │   │   ├── README.md
    │   │   ├── src/
    │   │   │   ├── main.ts
    │   │   │   ├── app.module.ts
    │   │   │   ├── common/
    │   │   │   │   ├── dtos/
    │   │   │   │   │   └── user.dto.ts
    │   │   │   │   ├── interfaces/
    │   │   │   │   │   └── user.interface.ts
    │   │   │   │   ├── filters/
    │   │   │   │   │   └── http-exception.filter.ts
    │   │   │   │   ├── interceptors/
    │   │   │   │   │   └── logging.interceptor.ts
    │   │   │   │   ├── guards/
    │   │   │   │   │   └── auth.guard.ts
    │   │   │   │   ├── decorators/
    │   │   │   │   │   └── user.decorator.ts
    │   │   │   │   ├── utils/
    │   │   │   │   │   └── logger.ts
    │   │   │   ├── modules/
    │   │   │   │   ├── auth/
    │   │   │   │   │   ├── auth.controller.ts
    │   │   │   │   │   ├── auth.module.ts
    │   │   │   │   │   ├── auth.service.ts
    │   │   │   │   │   ├── jwt.strategy.ts
    │   │   │   │   │   ├── auth.middleware.ts
    │   │   │   │   │   ├── tests/
    │   │   │   │   │   │   ├── auth.controller.spec.ts
    │   │   │   │   │   │   └── auth.service.spec.ts
    │   │   │   │   ├── users/
    │   │   │   │   │   ├── users.controller.ts
    │   │   │   │   │   ├── users.module.ts
    │   │   │   │   │   ├── users.service.ts
    │   │   │   │   │   ├── users.repository.ts
    │   │   │   │   │   ├── tests/
    │   │   │   │   │   │   ├── users.controller.spec.ts
    │   │   │   │   │   │   └── users.service.spec.ts
    │   │   │   │   ├── ping/
    │   │   │   │   │   ├── ping.controller.ts
    │   │   │   │   │   ├── ping.module.ts
    │   │   │   │   │   ├── ping.service.ts
    │   │   │   │   │   ├── tests/
    │   │   │   │   │   │   ├── ping.controller.spec.ts
    │   │   │   │   │   │   └── ping.service.spec.ts
    │   │   │   │   ├── go-client/
    │   │   │   │   │   ├── go-client.module.ts
    │   │   │   │   │   ├── go-client.service.ts
    │   │   │   │   │   ├── tests/
    │   │   │   │   │   │   └── go-client.service.spec.ts
    │   │   │   │   └── database/
    │   │   │   │       └── database.module.ts
    │   │   │   ├── config/
    │   │   │   │   └── configuration.ts
    │   │   │   ├── app.controller.ts
    │   │   │   ├── app.service.ts
    │   │   ├── test/
    │   │   │   └── app.e2e-spec.ts
    │   │   ├── .env
    │   │   ├── Dockerfile
    │   │   ├── .dockerignore
    │   │   ├── package-lock.json
    │   │   ├── package.json
    │   │   └── tsconfig.json
    │   └── go-basic-poc/
    │       ├── cmd/
    │       │   ├── main.go
    │       │   └── main_test.go
    │       ├── go.mod
    │       ├── go.sum
    │       ├── internal/
    │       │   ├── api/
    │       │   │   └── handlers/
    │       │   │       ├── ping.go
    │       │   │       ├── ping_test.go
    │       │   │       └── user.go
    │       │   │       └── user_test.go
    │       │   ├── app/
    │       │   │   ├── app.go
    │       │   │   └── app_test.go
    │       │   ├── auth/
    │       │   │   ├── auth.go
    │       │   │   ├── auth_test.go
    │       │   │   ├── middleware.go
    │       │   │   └── middleware_test.go
    │       │   ├── config/
    │       │   │   ├── config.go
    │       │   │   └── config_test.go
    │       │   ├── data/
    │       │   │   ├── user_repository.go
    │       │   │   └── user_repository_test.go
    │       │   ├── pkg/
    │       │   │   └── user_service/
    │       │   │       ├── user_service.go
    │       │   │       └── user_service_test.go
    │       │   ├── utils/
    │       │   │   ├── db.go
    │       │   │   ├── db_test.go
    │       │   │   ├── errors.go
    │       │   │   ├── errors_test.go
    │       │   │   ├── logger.go
    │       │   │   ├── logger_test.go
    │       │   │   ├── observability.go
    │       │   │   └── observability_test.go
    │       │   ├── routes/
    │       │   │   ├── routes.go
    │       │   │   └── routes_test.go
    │       ├── .env
    │       └── Dockerfile
    ```

    * This shows the expected directory structure for the monorepo.

---

## 🚀 Create NestJS Project

1.  **Install Nest CLI Globally:**

    ```bash
    npm i -g @nestjs/cli
    ```

    * Installs the NestJS command-line interface globally, allowing you to use the `nest` command.

2.  **Create NestJS Project:**

    ```bash
    nest new apps/nest-basic-poc --package-manager=npm
    ```

    * Creates a new NestJS project named `nest-basic-poc` inside the `apps` directory.
    * `--package-manager=npm` specifies that npm should be used as the package manager.

3.  **Install NestJS Dependencies:**

    ```bash
    cd apps/nest-basic-poc
    npm i @nestjs/jwt passport-jwt @nestjs/passport passport axios dotenv winston nest-winston
    ```

    * Navigates into the NestJS project directory.
    * Installs the required npm packages for JWT authentication, HTTP requests, logging, and environment variables.

4.  **NestJS File: `.env`:**

    ```env
    PORT=3000
    CLIENT_ID=sharedClientID
    CLIENT_SECRET=sharedClientSecret
    GO_SERVICE_URL=http://localhost:4000/ping
    ```

    * Create a `.env` file in the `apps/nest-basic-poc` directory and add the environment variables.

5.  **NestJS File: `src/controllers/app.controller.ts`:**

    ```ts
    import { Controller, Get, LoggerService, Headers, UnauthorizedException } from '@nestjs/common';
    import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';
    import { Inject } from '@nestjs/common';
    import { HttpService } from '../services/http.service';

    @Controller()
    export class AppController {
      constructor(
        @Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: LoggerService,
        private readonly httpService: HttpService
      ) {}

      @Get('ping')
      ping() {
        this.logger.log('NestJS /ping called');
        return { message: 'NestJS is alive!' };
      }

      @Get('call-go')
      async callGo(@Headers('client-id') clientId: string, @Headers('client-secret') clientSecret: string) {
        if (clientId !== process.env.CLIENT_ID || clientSecret !== process.env.CLIENT_SECRET) {
          this.logger.warn('Unauthorized access attempt to call-go');
          throw new UnauthorizedException();
        }
        this.logger.log('NestJS calling Go service');
        return this.httpService.callGoService();
      }
    }
    ```

    * Implements the `/ping` and `/call-go` endpoints, using the `HttpService` to make HTTP requests to the Go service.

6.  **NestJS File: `src/services/http.service.ts`:**

    ```ts
    import { Injectable, LoggerService, Inject, HttpException } from '@nestjs/common';
    import axios from 'axios';
    import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';

    @Injectable()
    export class HttpService {
      constructor(@Inject(WINSTON_MODULE_NEST_PROVIDER) private readonly logger: LoggerService) {}

      async callGoService() {
        const url = process.env.GO_SERVICE_URL;
        try {
          const res = await axios.get(url, {
            headers: {
              'client-id': process.env.CLIENT_ID,
              'client-secret': process.env.CLIENT_SECRET
            }
          });
          this.logger.log(`Response from Go service: ${JSON.stringify(res.data)}`);
          return { fromGo: res.data };
        } catch (err) {
          this.logger.error('Error calling Go service', err);
          throw new HttpException('Failed to reach Go service', 500);
        }
      }
    }
    ```

    * Defines the `HttpService` to handle HTTP requests to the Go service, including error handling and logging.

7.  **NestJS File: `src/app.module.ts`:**

    ```ts
    import { Module } from '@nestjs/common';
    import { AppController } from './controllers/app.controller';
    import { HttpService } from './services/http.service';
    import { WinstonModule } from 'nest-winston';
    import * as winston from 'winston';

    @Module({
      imports: [
        WinstonModule.forRoot({
          transports: [
            new winston.transports.Console({
              format: winston.format.combine(
                winston.format.timestamp(),
                winston.format.simple()
              )
            })
          ]
        })
      ],
      controllers: [AppController],
      providers: [HttpService],
    })
    export class AppModule {}
    ```

    * Configures the NestJS module, including Winston logger setup and service providers.

8.  **NestJS File: `src/main.ts`:**

    ```ts
    import { NestFactory } from '@nestjs/core';
    import { AppModule } from './app.module';
    import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';
    import * as dotenv from 'dotenv';

    dotenv.config();

    async function bootstrap() {
      const app = await NestFactory.create(AppModule);
      const logger = app.get(WINSTON_MODULE_NEST_PROVIDER);
      app.useLogger(logger);
      await app.listen(process.env.PORT || 3000);
      logger.log(`NestJS service running on port ${process.env.PORT}`);
    }
    bootstrap();
    ```

    * Sets up the NestJS application, including loading environment variables and configuring the logger.

---

## 🚀 Create Golang Project

1.  **Create Golang Project Directory and Initialize Module:**

    ```bash
    mkdir -p apps/go-basic-poc && cd apps/go-basic-poc
    go mod init go-basic-poc
    go get https://github.com/joho/godotenv
    ```

    * Creates the `apps/go-basic-poc` directory and navigates into it.
    * Initializes a new Go module named `go-basic-poc`.
    * Downloads the `godotenv` package for loading environment variables.

2.  **Golang File: `.env`:**

    ```env
    PORT=4000
    CLIENT_ID=sharedClientID
    CLIENT_SECRET=sharedClientSecret
    NEST_SERVICE_URL=http://localhost:3000/ping
    ```

    * Creates a `.env` file in the `apps/go-basic-poc` directory and adds the environment variables.

3.  **Golang File: `handlers/ping.go`:**

    ```go
    package handlers

    import (
        "encoding/json"
        "net/http"
        "log"
    )

    type Response struct {
        Message string `json:"message"`
    }

    func PingHandler(w http.ResponseWriter, r *http.Request) {
        log.Println("Received /ping request on Go service")
        json.NewEncoder(w).Encode(Response{Message: "Go is alive!"})
    }
    ```

    * Implements the `/ping` endpoint for the Go service.

4.  **Golang File: `handlers/call_nest.go`:**

    ```go
    package handlers

    import (
        "io"
        "log"
        "net/http"
        "os"
    )

    func CallNestHandler(w http.ResponseWriter, r *http.Request) {
        url := os.Getenv("NEST_SERVICE_URL")
        log.Printf("Calling NestJS service at: %s", url)

        req, err := http.NewRequest("GET", url, nil)
        if err != nil {
            log.Printf("Failed to create request: %v", err)
            http.Error(w, "Request creation failed", http.StatusInternalServerError)
            return
        }

        req.Header.Set("client-id", os.Getenv("CLIENT_ID"))
        req.Header.Set("client-secret", os.Getenv("CLIENT_SECRET"))

        res, err := http.DefaultClient.Do(req)
        if err != nil {
            log.Printf("Error calling NestJS service: %v", err)
            http.Error(w, "Failed to reach NestJS service", http.StatusInternalServerError)
            return
        }
        defer res.Body.Close()

        body, err := io.ReadAll(res.Body)
        if err != nil {
            log.Printf("Error reading response from NestJS: %v", err)
            http.Error(w, "Error reading response", http.StatusInternalServerError)
            return
        }

        log.Printf("Response from NestJS: %s", string(body))
        w.Header().Set("Content-Type", "application/json")
        w.Write(body)
    }
    ```

    * Implements the `/call-nest` endpoint, making an HTTP request to the NestJS service.

5.  **Golang File: `main.go`:**

    ```go
    package main

    import (
        "log"
        "net/http"
        "os"
        "go-basic-poc/handlers"
        "[github.com/joho/godotenv](https://github.com/joho/godotenv)"
    )

    func main() {
        if err := godotenv.Load(".env"); err != nil {
            log.Fatalf("Error loading .env file: %v", err)
        }

        port := os.Getenv("PORT")
        log.Printf("Starting Go service on port %s", port)

        http.HandleFunc("/ping", handlers.PingHandler)
        http.HandleFunc("/call-nest", handlers.CallNestHandler)

        log.Fatal(http.ListenAndServe(":"+port, nil))
    }
    ```

    * Sets up the Go HTTP server and registers the handlers.

---

## 🐳 Docker Files

1.  **`apps/nest-basic-poc/Dockerfile`:**

    ```dockerfile
    FROM node:18-alpine
    WORKDIR /app
    COPY package*.json ./
    RUN npm install
    COPY . .
    RUN npm run build
    CMD ["node", "dist/main"]
    ```

    * Defines the Dockerfile for building the NestJS image.

2.  **`apps/go-basic-poc/Dockerfile`:**

    ```dockerfile
    FROM golang:1.20
    WORKDIR /app
    COPY go.mod go.sum ./
    RUN go mod download
    COPY . .
    RUN go build -o go-service .
    CMD ["./go-service"]
    ```

    * Defines the Dockerfile for building the Go image.

---

## 🧪 Docker Compose

1.  **`docker-compose.yml`:**

    ```yaml
    version: '3.8'
    services:
      nest-service:
        build: ./apps/nest-basic-poc
        ports:
          - "3000:3000"
        env_file:
          - ./apps/nest-basic-poc/.env

      go-service:
        build: ./apps/go-basic-poc
        ports:
          - "4000:4000"
        env_file:
          - ./apps/go-basic-poc/.env
    ```

    * Defines the Docker Compose configuration for running the NestJS and Go services.

---

## 📘 Workspace README.md

```markdown
# Monorepo: NestJS + Golang

## 📦 Prerequisites

* Node.js
* Go
* Docker

## 📂 Project Structure

* `apps/nest-basic-poc`: NestJS API
* `apps/go-basic-poc`: Go API
```

## 🔧 Setup

```bash
git clone <repo>
cd monorepo-poc
npm install -g @nestjs/cli
cd apps/nest-basic-poc && npm install
```
## 🚀 Run Projects (Docker Compose with Exclude)

### Run All Services

```bash
./run-compose.sh
```

### Run Specific Services

```bash
./run-compose.sh nest-service go-service # Add or remove service names
```

### Exclude Specific Services

```bash
./run-compose.sh --exclude nest-service
./run-compose.sh --exclude nest-service --exclude go-service
./run-compose.sh --exclude nest-service go-service # runs go service, excludes nest.
```

## 🚀 Build and Run Projects (Docker Compose)

### Build All Services

```bash
./build-compose.sh
```

### Build Specific Services

```bash
./build-compose.sh nest-service go-service # Add or remove service names
```

### Exclude Specific Services

```bash
./build-compose.sh --exclude nest-service
./build-compose.sh --exclude nest-service --exclude go-service
./build-compose.sh --exclude nest-service go-service # runs go service, excludes nest.
```

# New NestJS project
```bash
nest new apps/new-nest-service
```

# New Go project
```bash
mkdir apps/new-go-service && cd apps/new-go-service && go mod init new-go-service
```