# goBoard
A minimalist & cross-platform clipboard history manager written in Go

## Features
- System tray icon
- FZF-like commandline UI
- REST API
- React Web UI

## Usage
goBoard accepts the following arguments:
`goboard monitor` - starts the clipboard monitor service
`goboard fzf` - starts the fzf interface to search clipboard history and copy selection to clipboard
`goboard web` - start REST API and React Web UI on http://localhost:3000
`goboard ls` - list clipboard history
`goboard cp <id>` - copy history item directly by ID
`goboard fullstack` - run both the clipboard monitor service and the web frontend/API
