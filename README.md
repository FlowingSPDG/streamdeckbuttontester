# StreamDeck Button Tester

Simple Stream Deck plugin that logs key events using `github.com/FlowingSPDG/streamdeck`.

## Build locally

- Linux/macOS:
  - `cd Source/code/cmd && go build -o ../../button_tester .`
- Windows:
  - `cd Source/code/cmd && GOOS=windows GOARCH=amd64 go build -o ../../button_tester.exe .`

## Package

Requires `DistributionTool`/`DistributionTool.exe` in repository root.

```bash
task package        # mac+win
task package:windows
```


