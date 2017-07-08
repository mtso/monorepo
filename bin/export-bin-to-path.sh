#!/usr/bin/env bash

echo "Running 'export $PWD/bin/:\$PATH'"

export PATH="`pwd`/bin/:$PATH"
