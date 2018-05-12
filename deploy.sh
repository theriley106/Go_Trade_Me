GOOS=linux go build -o main *.go
sleep 2
zip deployment.zip main

