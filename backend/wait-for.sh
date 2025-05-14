#!/bin/sh

host="$1"
shift
cmd="$@"

until nc -z "$host" 5432; do
  echo "Esperando a la base de datos en $host:5432..."
  sleep 1
done

echo "Base de datos disponible — iniciando backend"
exec $cmd
