STEPS TO UPDATE MESSAGE LIMITS

1. Locate where you installed d3tourrr's NomiKin Discord files.
    eg. nomikin-discord
2. Go to 'src' folder
3. Open the following files in Visual Studio or Notepad:
    a. companion.go
    b. companionutils.go
    c. nomirooms.go
    d. rooms.go
4. In each of these files, find the URL near the top of the page: github.com/d3tourrr/NomiKinGo and change it to github.com/cydonia1978/NomiKinGo
5. In your 'src' folder, go to the location line at the top of the screen (usually says This PC > Drive > nomikin-discord > src). Click in this line and type cmd
6. In the Command Prompt paste: 
    go get github.com/cydonia1978/NomiKinGo@v0.3.51 
    then
    go mod tidy
7. In the main nomikin-discord folder, run the start-linux-companion.sh file (make sure Docker Desktop is running); Your Nomi should now be using the modified version of d3tourr's integration, with a character receiving limit increase to 800!
