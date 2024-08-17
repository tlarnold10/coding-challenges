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
    - To run the tool, navigate to the uniqTool directory and run `uniqTool`. You can then specify the mode as an optional argument (-c to add a line count for each corresponding line, -d to just display only repeated lines, -u to display only unique lines, or nothing in order to run in standard mode, which will not display any duplicated lines), input file (required argument either the file name or `-` to read from the standard input), and output file (optional argument).

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

5. Build your own CronTab Tool (https://codingchallenges.fyi/challenges/challenge-cron)
    - Written in TypeScript, utilizing the Svelte framework and Bun package manager.
    - To run the tool, navigate to the `crontabTool/crontabTool` directory and run `bun --bun run dev`
    - Navigate to `http://localhost:5173` in your browser to use the crontab tool.
    - On the page there are some notes about using crontab and a pattern input box to enter your pattern. A new box will appear with the human readable version of the crotab pattern, or an error if it is not a valid crontab pattern.

6. Build your own Password Manager (https://codingchallenges.fyi/challenges/challenge-password-manager)
    - Written in Rust, using htmx for frontend interactivity, amux web framework, and askama for template rendering.
    - To run the tool, navigate to the `./passwordManager/password_manager` directory. Run `cargo run` to run the application.
        - You can also run `cargo build` to create the binary in the `target/debug` directory. From there you can run `./password_manager` to run the application.
    - Before the server started, the application checks if there is a Master password or not. 
        - If no Master password has been created, it will ask you to create one. The application with then stop and you must restart it to get it working again. 
        - If a Master password exists, you will be asked to enter it. You have three chances to do so.
    - Once the server is running, you can open up a web browser and navigate to `http://localhost:8675/` to see the web app.
