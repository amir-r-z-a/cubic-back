
MIGRATION_PATH="."
DATABASE_URL="postgresql://postgres:pgpassword@localhost:5432/cubikdb?sslmode=disable"

# Force the version to clean up dirty state
migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" force 20250103095001

migrate -path="$MIGRATION_PATH" -database "$DATABASE_URL" -verbose up
