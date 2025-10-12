# DevTracker CLI 

**DevTracker** is a minimalist command-line tool to help developers log their coding activity, track daily progress, and stay consistent — all from the terminal.

##  Features

- Check your streak 'devtrack today'
-  Log daily coding sessions with `devtrack add`
-  View all your entries for today with `devtrack today`
- Optional tagging for filtering and future stats
-  Simple JSON storage — no database needed
-  weekly xp gained. (well...you can change the xp rewarding implementation to what you like)

## Usage

### Add a new log entry:

-For now you can run 'go run main.go  <operation(e.g add)>  <YOUR LOG>  <optional tag>'

 Installation
1. Clone the repo
bash
Copy
Edit
git clone git@github.com:cargonew/devTracker.git
cd devTracker
2. Build the binary
bash
Copy
Edit
go build -o devtrack
3. (Optional) Move it to your PATH
bash
Copy
Edit
sudo mv devtrack /usr/local/bin
Now you can use devtrack from anywhere!..How cool is that!!

 Coming Soon
 -Git commit and push integration(very important!) 

 -Themes, badges, and gamification

 -Daily journaling mode

Log Format
All logs are saved in log.json:

json
Copy
Edit
[
  {
    "timestamp": "2025-07-04T21:50:00",
    "entry": "Solved array reversal problem",
    "tag": "leetcode"
  }
]
