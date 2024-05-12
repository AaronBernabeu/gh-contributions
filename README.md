# GitHub Contributions

This project retrieves contribution statistics from GitHub from the last 365 days for a given user and displays them in JSON format.

## Installation

1. Clone the project from the following URL:

```bash
git clone https://github.com/Aaronidas/gh-contributions.git
```

2. Navigate to the project directory:

```bash
cd gh-contributions
```

3. Install dependencies using `go mod tidy`:

```bash
go mod tidy
```

4. Create a `.env` file based on the provided `.env.dist`. You can also set the required environment variables directly without creating a `.env` file. The required variables are:

- `GH_TOKEN`: Your GitHub token.
- `GH_USERNAME`: The username whose contributions you want to retrieve.

## Usage

Run the program to display contribution statistics in JSON format:

```bash
go run cmd/gh-cli/main.go contributions
```

The output will be in the following format:

```json
{
  "name": "Aaron Bernabeu",
  "today": 2,
  "week": 14,
  "month": 63,
  "year": 1324
}
```