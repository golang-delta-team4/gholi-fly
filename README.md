# Gholi Fly Project

Welcome to **Gholi Fly**, a multi-service Go application orchestrated with Docker Compose and NGINX. This project contains multiple primary microservices:

1. **Bank Service** – located in `micro-services/bank`
2. **Hotel Service** – located in `micro-services/hotel`

The project provides separate configurations for **development** (hot reload) and **production** environments.

---

## How to Run the Project

### 1. Setup Configuration Files

Each service requires two configuration files:  
- **Development**: `config-dev.json`  
- **Production**: `config.json`  

1. Navigate to each service folder:  
   ```bash
   cd micro-services/bank
   cp sample-config.json config-dev.json
   cp sample-config.json config.json

   cd ../hotel
   cp sample-config.json config-dev.json
   cp sample-config.json config.json
   ```

2. Copy the `air-template.toml` to each service folder as `.air.toml`:
    ``` bash
    cp air-template.toml micro-services/bank/.air.toml
    cp air-template.toml micro-services/hotel/.air.toml
    ```

---

### 2. Run in Development Mode
To run the project in development mode with hot reload:
1. Start all services:
```bash
docker-compose -f docker-compose-dev.yaml up --build
```
2. Access the services:
- **Bank Service**: `http://localhost:8080`  
- **Hotel Service**: `http://localhost:8081`
- **Terminal Route Service**: `http://localhost:8082`
- **User**: `http://localhost:8083` 
- **Transportation**: `http://localhost:8084`

- **NGINX (Proxy)**: `http://localhost:8888`  

---

### 3. Run in Production Mode
To run the project in production mode:
1. Ensure the `config.json` files are set up for each service.
2. Start the services:
```bash
docker-compose up --build
```
3. Access the NGINX reverse proxy:
```
http://localhost/api/v1/
```

---

## Notes
- **Configurations**: Always create `config-dev.json` and `config.json` from the provided `sample-config.json` in each service directory.

- **Hot Reload**: Air is used in development mode. Make sure `.air.toml` exists in each service directory.

- **NGINX**: The reverse proxy configuration is in `nginx-config-dev.conf` (development) and `nginx-config.conf` (production).

- **Databases**: PostgreSQL and Redis are started automatically via Docker Compose.


