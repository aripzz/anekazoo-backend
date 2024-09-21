## Getting Started

These instructions will help you get the project up and running on your local machine.

### Prerequisites

- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- [Dbmate](https://github.com/amacneil/dbmate)

### Usage

Explain how to use the project and provide examples.
1. **Clone Repository**

   Clone repositori:

   ```shell
   git clone https://github.com/aripzz/anekazoo-backend
   cd anekazoo-backend
    ```

2. **Create .env File**

   Copy the example environment file:

   ```shell
    cp .env.example .env
    ```

3. **Start the Docker Compose**

   ```shell
    docker-compose up --build
    ```
   Run Containers in Detached Mode (containers will run in the background):
   ```shell
    docker-compose up --build -d
    ```

4. **Start the Migration**
   ```shell
    dbmate -u "mysql://cPZKdNerJD:XMhRNy6m78@127.0.0.1:3306/anekazoodb" up
    ```

5. **Services**
- API (Golang)          = http://localhost:3000
- Documention Swagger   = http://localhost:3000/swagger 
- MySQL Database        = localhost:3306
- PhpMyAdmin            = http://localhost:8081 (Server: mysql, user: root, password: root)
