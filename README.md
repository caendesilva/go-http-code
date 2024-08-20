# GO HTTP Code

[![GitHub Release](https://img.shields.io/github/v/release/caendesilva/go-http-code)](https://github.com/caendesilva/go-http-code/releases)

`http-code` is a simple command-line tool written in Go that provides information about HTTP status codes. Just pass in an HTTP status code, and the tool will return a brief description of the code.

## Features

- Supports common HTTP status codes
- Simple and blazing fast, in a small binary
- Cross-platform: Works on Windows, macOS, and Linux

## Installation

You can download the latest release of `http-code` from the [GitHub Releases](https://github.com/caendesilva/go-http-code/releases) page.

### Quick Install

If you're on a Unix-based system (macOS or Linux), you can run the following command to download and install the tool:

```bash
curl -L https://github.com/caendesilva/go-http-code/releases/latest/download/http-code-$(uname -s | awk '{print tolower($0)}')-amd64 -o http-code
chmod +x http-code
sudo mv http-code /usr/local/bin/

http-code 500 # Get information about HTTP status code 500
```

If this does not work for your system, keep reading for more detailed installation instructions.


### Steps to Install

1. Go to the [Releases](https://github.com/caendesilva/go-http-code/releases) page on GitHub.
2. Download the appropriate binary for your operating system:
   - `http-code-windows-amd64.exe` for Windows
   - `http-code-darwin-amd64` for macOS
   - `http-code-linux-amd64` for Linux
3. Rename the downloaded file to `http-code` (or keep it as `http-code.exe` on Windows).
4. Place the binary in a directory included in your system's `PATH` to use it from anywhere.

### Example Commands for Installing

#### Windows

1. Download `http-code-windows-amd64.exe`.
2. Rename it to `http-code.exe` (optional).
3. Move it to a directory included in your `PATH` (e.g., `C:\Windows\System32`).

#### macOS/Linux

1. Download `http-code-darwin-amd64` (macOS) or `http-code-linux-amd64` (Linux).
2. Rename it to `http-code`:
   ```bash
   mv http-code-darwin-amd64 http-code  # For macOS
   mv http-code-linux-amd64 http-code   # For Linux
   ```
3. Make it executable:
   ```bash
   chmod +x http-code
   ```
4. Move it to a directory in your `PATH`, such as `/usr/local/bin`:
   ```bash
   sudo mv http-code /usr/local/bin/
   ```

## Usage

Once installed, you can use `http-code` from the terminal or command prompt. Simply pass an HTTP status code as an argument to get information about the code.

### Example

```sh
http-code 500
```

**Output:**
```
500 Internal Server Error - The server has encountered a situation it doesn't know how to handle.
```

### More Examples

```sh
http-code 200
http-code 404
http-code 503
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests on [GitHub](https://github.com/caendesilva/go-http-code).

## License

This project is licensed under the MIT License. See the [LICENSE.md](LICENSE.md) file for more details.
