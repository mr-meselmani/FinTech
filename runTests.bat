@echo off

REM Step 1: Ensure Go is installed on the machine
REM Assuming Go is already installed and added to the system's PATH

REM Step 2: Run tests
echo Running tests...
go test ./tests/... -v

REM Pause to keep the console window open
pause