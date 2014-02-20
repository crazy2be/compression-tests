Testing chunk compression for buildblast.

main.go is a "server" that outputs compressed and base64'd data on stdout

main.js in a "client" that reads compressed and base64'd data on stdin, and
reverses the operation.

Running ./test.sh will make sure the client's output is the same as the server's
input, and when it's done, it'll display the size difference of the two files
(i.e. compressed and "raw").
For the currently provided files, it's about 2 orders of magnitude (49248 vs
676).

The difference would be even better (~33%) if we didn't base64 the compressed
output, but I didn't do that here for ease of debugging (base64'd data is easier
to work with in most cases than binary data).
