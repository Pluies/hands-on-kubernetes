#!/bin/bash

set -euo pipefail

for DIR in $(find . -type d -d 1)
do
  echo "Queuing directory: $DIR"
  (cd $DIR/src && make install) &
done
wait
