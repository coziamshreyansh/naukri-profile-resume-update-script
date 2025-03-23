package main

import (
	"context"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/chromedp/chromedp"
)

const (
	email    = "shreyansh.tiwari8@gmail.com"
	password = "Test@123"
	resume   = "Users/LENOVO/Downloads/Shreyansh_2YOE_SDE1.pdf"
)

func updateNaukriResume(ctx context.Context, resumePath string) error {
	err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.naukri.com/nlogin/login"),
		chromedp.WaitVisible(`#usernameField, input[placeholder='Enter your active Email ID / Username']`, chromedp.ByQuery),
		chromedp.SendKeys(`#usernameField, input[placeholder='Enter your active Email ID / Username']`, email, chromedp.ByQuery),
		chromedp.SendKeys(`#passwordField, input[placeholder='Enter your password']`, password, chromedp.ByQuery),
		chromedp.Click(`button[type='submit']`, chromedp.ByQuery),
		chromedp.WaitVisible(`.user-name, .nI-gNb-drawer__bars`, chromedp.ByQuery),
		chromedp.Navigate("https://www.naukri.com/mnjuser/profile"),
		chromedp.WaitVisible(`input[type='button'].dummyUpload, button:contains('UPDATE RESUME')`, chromedp.ByQuery),
		chromedp.Click(`input[type='button'].dummyUpload, button:contains('UPDATE RESUME')`, chromedp.ByQuery),
		chromedp.SetUploadFiles(`input[type='file']`, []string{resumePath}, chromedp.ByQuery),
		chromedp.Sleep(20*time.Second), // Consider using a more robust wait here.
	)
	return err
}

func main() {
	// Create Chrome instance
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Set up timeout
	ctx, cancel = context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	// Get absolute path of resume file
	resumePath, err := filepath.Abs(resume)
	if err != nil {
		log.Fatalf("Failed to get resume path: %v", err)
	}

	err = updateNaukriResume(ctx, resumePath)
	if err != nil {
		log.Fatalf("Automation failed: %v", err)
	}

	fmt.Println("Resume update completed successfully!")
}
