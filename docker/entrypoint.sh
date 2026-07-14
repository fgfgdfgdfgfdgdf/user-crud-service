#!/bin/sh
set -e

atlas migrate apply --env local

exec "$@"
