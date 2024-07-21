## Cara Menjalankan
Langkah:
1. docker compose --build
2. docker exec -it <db-container-id> psql -U postgres -d delos -f /docker-entrypoint-initdb.d/init.sql

## API Documentation
https://app.swaggerhub.com/apis/MASELON2030/go-delos/1.0.0
