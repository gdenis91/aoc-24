package aoc

import (
	"fmt"
	"os"
)

func GetSessionCookie() string {
	return os.Getenv(EnvAOCSession)
}

func PrintSessionHelp() {
	fmt.Printf(
		"Set %s environment variable\n",
		EnvAOCSession,
	)
	fmt.Println("Instructions:")
	fmt.Println("1. Visit the Advent of Code website: https://adventofcode.com/2024")
	fmt.Println("2. Log in to your account.")
	fmt.Println("3. Open the developer tools (F12).")
	fmt.Println("4. Go to the Application tab (or Storage tab in some browsers).")
	fmt.Println("5. Select adventofcode in the Cookies section in the left sidebar.")
	fmt.Println("6. Copy the value of the session cookie.")
	fmt.Println("7. Set the AOC_SESSION environment variable to the copied value.")
	fmt.Println()
	fmt.Println("Note: You can open the developer tools directly by clicking this link: ")
	fmt.Println("[Developer Tools](javascript:document.dispatchEvent(new KeyboardEvent('keydown',{'key':'F12'})))")
}
