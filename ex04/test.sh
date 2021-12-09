go build calculation.go
./calculation 12 4 | cat -e
./calculation  a 4 | cat -e
rm calculation
#提出前にエラー時の改行チェック
