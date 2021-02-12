Fiber + mongodb + redis Boilerplate
for real world application example

1. chmod 777 ./docker-compose.yml

2. export -p | grep UID
   ini untuk export UID di Linux

3. chmod 777 ./mongo-init.sh

4. mkdir data

5. chmod 777 ./docker-up.sh

6. ./docker-up.sh

7. go get github.com/joho/godotenv/cmd/godotenv

8. ensure $GOPATH/bin in your $PATH
   `close all bash & editor and reopen test call "godotenv"`

9. ./start.sh

10. Update Documents MUST ALWAYS MAKE NEW STRUCT

11. Create Doc always create new Struct From model for database consistency
