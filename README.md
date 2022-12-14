# hts-assoc-cli-tool

A simple and hastily put together CLI tool that queries a mirror node for accounts with an exact balance for a specific token ID. Results are saved to a CSV file in a format compatible with Calaxy's HTS distrubition tool.

## Usage

Download the executable for your machine (macOS, Windows, Linux) or build the tool yourself from source.

Pass the token ID using the ```-token``` flag (required).

Enter the amount to airdrop using the ```-amount``` flag (optional -- defaults to 0).

Enter the exact balance of the token an account must have using the ```-balance``` flag (optional -- defaults to 0)

Enter the output CSV filename using the ```-file``` flag (optional -- defaults to results.csv).

### macOS
```
$ ./hts-assoc-cli-tool-macos -token=0.0.1 -amount=10 -balance=1 -file=results.csv
```

### Windows
```
$ ./hts-assoc-cli-tool-windows.exe -token=0.0.1 -amount=10 -balance=1 -file=results.csv
```
### Linux
```
$ ./hts-assoc-cli-tool-linux -token=0.0.1 -amount=10 -balance=1 -file=results.csv
```
