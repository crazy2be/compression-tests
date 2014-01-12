#!/bin/bash

go build main.go
./main < test.raw > test.base64
js main.js < test.base64 > test.out
diff test.raw test.out
stat test.raw test.base64