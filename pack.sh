#!/bin/bash

# windows amd64
#echo 'Start pack windows amd64...'
#GOOS=windows GOARCH=amd64 go build ./
#tar -czvf "${RELEASE}/zhcp_api-windows-amd64.tar.gz" zhcp_api.exe config.toml log/.gitignore LICENSE README.md
#rm -rf zhcp_api.exe

#echo 'Start pack windows X386...'
#GOOS=windows GOARCH=386 go build ./
#tar -czvf "${RELEASE}/zhcp_api-windows-386.tar.gz" zhcp_api.exe config.toml log/.gitignore LICENSE README.md
#rm -rf zhcp_api.exe

#echo 'Start pack linux amd64'
#CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./
#tar -czvf "${RELEASE}/zhcp_api.tar.gz" zhcp_api config/config.ini cert/STAR.188sycp.com.crt cert/STAR.188sycp.com.key cert/STAR.37556.com.crt cert/STAR.37556.com.key log/.gitignore
#rm -rf zhcp_api

#echo 'Start pack linux 386'
#GOOS=linux GOARCH=386 go build ./
#tar -czvf "${RELEASE}/zhcp_api-linux-386.tar.gz" zhcp_api config.toml log/.gitignore LICENSE README.md
#rm -rf zhcp_api

#echo 'Start pack mac amd64'
#GOOS=darwin GOARCH=amd64 go build ./
#tar -czvf "${RELEASE}/zhcp_api-mac-amd64.tar.gz" zhcp_api config.toml log/.gitignore LICENSE README.md
#rm -rf zhcp_api

echo 'END'
