windows:
        GOOS=windows go build -o bin/lol.exe -ldflags="-s -w" ./src