@echo off

REM Step 1: Ensure Go is installed on the machine
REM Assuming Go is already installed and added to the system's PATH

REM Step 2: Download project dependencies
go mod download

REM Step 3: Start the server
go run cmd/main.go

REM Pause to keep the console window open
pause