@echo off
cd src\management


cd ..
call go build .\vpnauth\vpnauth.go
call go build
move vpnauth.exe ..\dist\vpnauth.exe
move src.exe ..\dist\easyvpn.exe
copy server.conf ..\dist\server.conf
cd ..