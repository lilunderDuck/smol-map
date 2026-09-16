run:
  go build -o ./out/main.exe ./src
  ./out/main.exe -tiny ./tiny/yarn-1.20.1-mappings.tiny
build:
  go build -o ./out/smolmap.exe ./src