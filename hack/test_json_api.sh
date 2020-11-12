#!/usr/bin/env bash

set -exo pipefail

function cleanup()
{
  echo "stopping web-cli"
  kill "$PID"
}

trap cleanup EXIT
nohup go run ./ serve &
PID="$!"

# wait until web-cli seems ready
while ! nc -z localhost 8080; do
  sleep 0.1 # wait for 1/10 of the second before check again
done
echo "running web-cli with PID: "$PID""


echo "testing if json api is returning 5 news elements..."
JSONNEWS=$(curl -s http://localhost:8080/api/news | jq '.news | length')

if [ "$JSONNEWS" != 5 ]
then
  echo "failed to get news from 'api/news', got response: ""$JSONNEWS"
  exit 1
fi

echo "success! :-)"