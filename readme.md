# Recipe System
A simple CRUD API for managing recipes

## Tech Stack
- **Language**: Go 1.25
- **Framework**: Fiber (HTTP server)
- **Database**: SQLite
- **Database Migration**: Goose
- **Testing**: Postman, Integration Testing in app.tracks.run
- **Containerization**: Docker, Docker Compose
- **Server**: Fly.io

---

## How to Run (Docker)
### Build & Run with Docker Compose
```bash
docker-compose up --build
```

This will:
- Start **Recipe System** on port `3000`

## 📬 API Endpoints
| Method | Endpoint       | Description      |
| ------ | -------------- | ---------------- |
| POST   | `/recipes`     | Create recipe    |
| GET    | `/recipes`     | Get all recipes  |
| GET    | `/recipes/:id` | Get recipe by ID |
| PATCH  | `/recipes/:id` | Update recipe    |
| DELETE | `/recipes/:id` | Delete recipe    |

## Database Migrations
Migrations are stored in `migrations/` folder.
To run migration or create a new one, please run this
```bash
goose create migrations/<migration_name_or_purpose> sql
```
Goose runs migrations automatically on startup.

## Testing
you can import postman file inside folder postman to test all the api.

## Deploy on Fly.io
1. Install Fly CLI
```bash
   curl -L https://fly.io/install.sh | sh
```

2. Login
```bash
   fly auth login
```

3. Launch project
```bash
   fly launch
```
- Select a region close to you 
- Confirm Dockerfile detection 
- This generates a fly.toml config

4. Deploy
```bash
   fly deploy
```
Fly.io will build the Docker image, deploy, and give you a public URL like:
`https://recipe-api.fly.dev`

## Author

Apriyanto Arkiang — Backend engineer with 8+ years of experience.

---