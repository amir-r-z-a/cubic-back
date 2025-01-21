MIGRATION_PATH="."
DATABASE_URL="postgresql://postgres:pgpassword@localhost:5432/cubikdb?sslmode=disable"

echo "Cleaning up dirty state..."
migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" force 20250103164809

echo "Running down migrations..."
migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" down -all

echo "Running up migrations..."
migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" -verbose up

