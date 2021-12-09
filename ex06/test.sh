go build echo42.go
./echo42 12 34 555 | cat -e
./echo42 -s / a bc def | cat -e
./echo42 -s XXXXXXX 12 34 56789 | cat -e
./echo42 -n 12 34 555 | cat -e
./echo42 -help
./echo42 | cat -e
rm echo42
