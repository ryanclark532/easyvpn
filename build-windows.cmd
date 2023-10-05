@echo off
cd src\management


cd ..
call go build .\vpnauth\vpnauth.go
call go build
move src.exe ..\dist\easyvpn.exe
copy server.conf ..\dist\server.conf
cd ..