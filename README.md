# Introduction
This repo is used to store my coding challenges from https://codingchallenges.fyi. I want to improve my coding skills and this is just one of the ways I am doing so. 

# Challenges
1. Write Your Own wc Tool: https://codingchallenges.fyi/challenges/challenge-wc/
    - Written in Go.
    - To run the tool, navigate the the wcTool directory and run `wcTool`. You can then specify specific arguments in order to count lines, characters, words, or bytes. Run `man wc` to get a full breakdown of the existing wc tool in Linux.
2. Create your own web server: https://codingchallenges.fyi/challenges/challenge-webserver/
    - Written in Go. 
    - Go already has concurrency built into their http api, so I didn't have to do anything special with that.
    - In order to run the http web server, navigate to the webServer directory and run `webServer`. Then open another terminal and run the `curl http://localhost:6969` command. Otherwise, you can open a browser and enter `http://localhost:6969` in the address bar. Currently, the only valid routes are `/` and `/about`, all other routes result in a 404 error.
3. Build your own uniq tool: https://codingchallenges.fyi/challenges/challenge-uniq/
    - Written in Go.
    - To run the tool, naviage to the uniqTool directory and run `uniqTool`. You can then specify the mode as an optional argument (-c to add a line count for each corresponding line, -d to just display only repeated lines, -u to display only unique lines, or nothing in order to run in standard mode, which will not display any duplicated lines), input file (required argument either the file name or `-` to read from the standard input), and output file (optional argument).
