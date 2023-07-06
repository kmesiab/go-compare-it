#!bin/sh

r=$(curl localhost:8080/healthcheck -sw "%{response_code}" -o /dev/null)

if [ "$r" = "200" ]; then
   echo "The API has successfully responded to a healthcheck"
   exit 0
else
   echo "The API did not respond to a healthcheck"
   exit 1
fi
