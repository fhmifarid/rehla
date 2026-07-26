#!/usr/bin/env sh
set -eu

diff \
  backend/migrations/000001_foundation.up.sql \
  backend/internal/database/migrations/sql/000001_foundation.up.sql
