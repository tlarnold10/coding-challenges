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
4. Build your own Discord bot: https://codingchallenges.fyi/challenges/challenge-discord
    - Written in Go.
    - To run the tool, navigate to the disdaccbot directory and run `disdaccbot` and provide the discord bot token after the `-t` flag.
    - Once the bot is running, you should be able to see a log stating the bot is now running. The bot is currently connected to my own personal discord server. But I could go to my discord server and run the following commands:
        1. `Hello`: which would result in the bot responding with `Hello`
        2. `!quote`: will grab a random quote and respond with that quote and the author.
        3. `!challenge`: will grab a random challenge from the Coding Challenge json file.
        4. `!list`: will grab and list all the Coding Challenges from the json file.
        5. `!add`: will take a url of a Coding Challenge, confirm it is a valid url and then respond with the title from the html page.
    - Next steps would be to add the functionality to run on something other than my PC as well as actually add the coding challenge to my local json.
