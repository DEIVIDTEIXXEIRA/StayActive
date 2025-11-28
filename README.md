# Stay Active

A CLI application that helps you maintain an "active" status on Slack and other apps by automatically pressing random keys at specified intervals. This tool is designed to prevent micromanagement issues by keeping your system active when you need to step away briefly.

## 📋 Requirements

- [Go (Golang)](https://go.dev/doc/install)

## 🚀 Installation

```bash
go install github.com/DEIVIDTEIXXEIRA/StayActive@latest
```
## 💻 Usage

### Basic Usage

```bash
# Default: Press a key every 1 minute for 1 hour 30 minutes
stay-active

# Custom interval and duration
stay-active --interval 5m --duration 1h --verbose

# Short flags
stay-active -i 30s -d 2h30m -v

# Raw numbers (treated as minutes)
stay-active --interval 2 --duration 45
```

### Command Line Options

| Flag | Short | Default | Description |
|------|-------|---------|-------------|
| `--help` | `-h` | - | Help for stay-active |
| `--interval` | `-i` | `1m` | How often to press a key (e.g., `30s`, `5m`, `1h`, `1h30m`, or raw number in minutes) |
| `--duration` | `-d` | `1h30m` | How long to run the application (e.g., `30s`, `5m`, `1h`, `1h30m`, or raw number in minutes) |
| `--verbose` | `-v` | `false` | Enable verbose output to see what's happening |