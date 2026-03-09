#!/bin/sh

IP=1.1.1.1

echo "Checking network connectivity to $IP"

if ping -c 1 $IP &> /dev/null; then
    echo "Host seems to be reachable. Good."
else
    echo "Host could not be reached. Bad."
fi
