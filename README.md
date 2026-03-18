# devTracker

A minimalist CLI tool to log your coding activity, track daily progress, and stay consistent — all from the terminal.

## Features

- Log what you worked on with `devtrack add`
- View today's entries with `devtrack today`
- Track streaks and earn XP for consistent output
- Optional tags for filtering entries
- Simple JSON storage — no database needed

## Install

```bash
git clone https://github.com/roketrakunn/devTracker.git
cd devTracker
go build -o devtrack
sudo mv devtrack /usr/local/bin
```

## Usage

```bash
devtrack add "implemented variable storage in the JIT compiler" --tag compilers
devtrack today
```

## Log format

Entries are saved to `log.json`:

```json
[
  {
    "timestamp": "2025-07-04T21:50:00",
    "entry": "implemented variable storage in the JIT compiler",
    "tag": "compilers"
  }
]
```

## Roadmap

- [ ] Git commit integration
- [ ] Weekly stats summary
- [ ] Daily journaling mode
