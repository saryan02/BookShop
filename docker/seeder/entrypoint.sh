#!/bin/sh

echo "==== Starting DB Seeding ===="

echo "Seeds:"
find /seeds -type f

for f in /seeds/common/*.sql; do
    if [ -f "$f" ]; then
        echo "Applying common seed: $f"
        psql -h postgres -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$f"
    fi
done

if [ -d "/seeds/$SEED_ENV" ]; then
    for f in /seeds/$SEED_ENV/*.sql; do
        if [ -f "$f" ]; then
            echo "Applying $SEED_ENV seed: $f"
            psql -h postgres -U "$POSTGRES_USER" -d "$POSTGRES_DB" -f "$f"
        fi
    done
fi

echo "==== Seeding Completed Successfully! ===="