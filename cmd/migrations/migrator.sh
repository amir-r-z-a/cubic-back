
MIGRATION_PATH="./cmd/migrations"
DATABASE_URL="postgresql://postgres:pgpassword@localhost:5432/cubikdb?sslmode=disable"

migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" -verbose down

migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" -verbose up
