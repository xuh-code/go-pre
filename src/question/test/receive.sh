#!/bin/bash

while true; do
  data=$(nc -u -w0 -l 12345)
  echo "$data"
done