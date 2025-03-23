# naukri-profile-resume-update-script (Go)

This Go script automates the process of logging into Naukri.com and updating the resume using `chromedp`.

## Prerequisites

1. Install [Go](https://go.dev/dl/) (Version 1.21+ recommended).
2. Install Google Chrome.
3. Install required Go dependencies.

## Installation

Clone the repository and navigate to the directory:

```sh
git clone https://github.com/coziamshreyansh/naukri-profile-resume-update-script.git
cd naukri-profile-resume-update-script
```

Initialize a Go module:

```sh
go mod init naukri-profile-resume-update-script
```

Install dependencies:

```sh
go get github.com/chromedp/chromedp
```

## Usage

### 1. Update Credentials

Modify the `main.go` file to include your Naukri login credentials:

```go
const (
    email    = "your-email@example.com"
    password = "your-secure-password"
    resume   = "path/to/your-resume.pdf"
)
```

### 2. Run the Script

```sh
go run main.go
```

### 3. Debugging with Headful Mode

If the script fails, enable headful mode to visually debug:

```go
opts := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.Flag("headless", false))
ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
ctx, cancel = chromedp.NewContext(ctx)
```

### 4. Build Executable (Optional)

To create an executable file:

```sh
go build -o naukri-profile-resume-update-script main.go
./naukri-profile-resume-update-script
```

## Troubleshooting

- **Context Deadline Exceeded**: Increase the timeout in `context.WithTimeout`.
- **Captcha Issues**: Log in manually once before running the script.
- **Selectors Not Found**: Check Naukri's website for updated element IDs.
- **Slow Execution**: Increase `Sleep` durations in the script.
