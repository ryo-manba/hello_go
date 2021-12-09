go build isEmailAddress.go
./isEmailAddress | cat -e
./isEmailAddress marvin@student.42tokyo.jp abc@def.123 | cat -e
rm isEmailAddress
