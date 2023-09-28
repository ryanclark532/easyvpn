@echo off
cd src\management
call npm install
call npm run build
cd ..
call go build
move src.exe ..\dist\easyvpn.exe
copy server.conf ..\dist\server.conf
cd ..