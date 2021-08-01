docker-compose down
docker build --tag mockingo .
docker compose build --no-cache
docker compose up -d 
docker-compose logs -f 