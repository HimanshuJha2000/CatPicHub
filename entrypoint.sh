#!/usr/bin/dumb-init /bin/sh
set -euo pipefail

initialize()
{
    echo "Initializing CatPicHub microservice"
}

run_migrations()
{
    echo "Run migrations on live db"
    su-exec appuser /bin/migrate  up
}

collect_process()
{
    # Get pid for app process
    APP_PID=$!

    # Wait for app to finish.
    # To resolve graceful shutdown, app needs to run in foreground and nginx in background
    wait "$APP_PID"
}

start_catPicHub()
{
    echo "Starting CatPicHub"
    su-exec appuser /bin/players  &

    collect_process
}

initialize
run_migrations
start_catPicHub