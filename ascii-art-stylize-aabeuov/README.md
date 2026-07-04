# ASCII Art Web

A web-based ASCII art generator written in Go that renders text as ASCII art using banner style files.

## Description

ASCII Art Web is a full-stack web application that provides a graphical user interface for generating ASCII art. Users can input text, select from multiple banner styles (standard, shadow, thinkertoy), and view the generated ASCII art directly in their browser. The application features a clean, modern UI with real-time feedback and proper error handling.

## Authors

- ASCII Art Web Contributors
Abeuov Alisher
Arman Tanatarov

## Usage

### Running the Server

1. Make sure Go is installed on your system
2. Navigate to the project directory
3. Run the server:

```bash
go run .
```

The server will start on `http://localhost:8080`

4. Open your browser and navigate to `http://localhost:8080`
5. Enter your text, select a banner style, and click "Generate ASCII Art"

### HTTP Endpoints

- **GET /**: Serves the main HTML page with the GUI
- **POST /ascii-art**: Accepts form data with `text` and `banner` parameters, returns ASCII art

### Form Parameters

| Parameter | Description | Required |
|-----------|-------------|----------|
| `text` | The text to render as ASCII art | Yes |
| `banner` | Banner style: `standard`, `shadow`, or `thinkertoy` | No (defaults to `standard`) |

## Implementation Details

### Algorithm

1. **Banner File Reading**: The application reads banner files (standard.txt, shadow.txt, thinkertoy.txt) from the `banners/` directory. Each banner file contains ASCII representations of printable characters (32–126).

2. **Character Mapping**: Each character in the banner file is mapped to an 8-line representation. The files are parsed line-by-line, creating a map where each ASCII character is associated with its 8-line visual representation.

3. **ASCII Rendering**: When text is submitted:
   - The input text is validated
   - Each character in the input is looked up in the character map for the selected banner
   - For each input character, the algorithm iterates through the 8 rows and concatenates the visual representations horizontally
   - Multiple input lines are processed line-by-line, with 8 rows of output generated per input line
   - Unknown characters are replaced with spaces

4. **HTTP Status Codes**:
   - `200 OK`: Request successful, ASCII art generated
   - `400 Bad Request`: Invalid input (empty text, invalid banner)
   - `404 Not Found`: Banner file not found
   - `500 Internal Server Error`: Unexpected server errors

5. **Error Handling**: The application gracefully handles various error conditions including missing banner files, invalid form data, and unexpected server errors, returning appropriate HTTP status codes.

### Project Structure

```
.
├── main.go              # Web server entry point
├── handlers.go          # HTTP route handlers (GET /, POST /ascii-art)
├── ascii.go             # Core banner reading and ASCII rendering logic
├── ascii_test.go        # Unit tests
├── go.mod               # Go module definition
├── README.md            # Documentation
├── banners/             # Banner style files
│   ├── standard.txt
│   ├── shadow.txt
│   └── thinkertoy.txt
└── templates/           # HTML templates
    └── index.html       # Main GUI page
```

### Features

- **Multiple Banner Styles**: Choose from standard, shadow, or thinkertoy ASCII art styles
- **Real-time Feedback**: Instant error messages and success notifications
- **Responsive Design**: Works on desktop and mobile devices
- **Clean GUI**: Modern, intuitive user interface with gradient styling
- **Multi-line Support**: Generate ASCII art for text with multiple lines
- **Proper Error Handling**: Returns appropriate HTTP status codes and user-friendly error messages
- **Standards Compliance**: Uses only Go standard library packages
- **Good Practices**: Follows Go conventions and best practices

## Banner Files

Banner files are plain text files located in the `banners/` directorys, named `<standard>.txt, <thinkertoy>.txt, <shadow>.txt`.

Each file encodes all printable ASCII characters (32–126) in order. Every character is represented by exactly **8 lines** of ASCII art, with consecutive characters separated by a single blank line.

```
banners/
└── standard.txt  # Default banner style
    shadow.txt
    thinkertoy.txt
```

## Core Functions

### `RenderASCII(input, bannerName string) (string, error)`
Renders the input string as ASCII art using the specified banner style. Splits input on newlines and builds 8-row art for each line.

### `ReadBannerFile(bannerName string) (map[rune][]string, error)`
Reads and parses a banner file, returning a map of each character's 8-line ASCII art representation.

### `GetCharacterWidth(characters map[rune][]string) int`
Returns the fixed column width of characters in a given banner.

### `processEscapeSequences(s string) string`
Converts escape sequences in CLI input (`\n`, `\t`, `\r`, `\\`) into their actual characters before rendering.

## Error Handling

- If fewer than 2 arguments are provided, usage instructions are printed and the program exits with code `1`.
- If the banner file cannot be found or read, an error is written to `stderr` and the program exits with code `1`.
- Characters not present in the banner file are rendered as blank columns to preserve alignment.