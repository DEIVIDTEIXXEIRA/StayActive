package cmd

import (
	"fmt"
	"math/rand/v2"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/spf13/cobra"
)

var (
	interval string
	duration string
	verbose  bool
)

var rootCmd = &cobra.Command{
	Use:   "stay-active",
	Short: "Press random letters at specified intervals to keep your system active",
	Long: `Press random letters at specified intervals to keep your system active.

Examples:
  stay-active --interval 5m --duration 1h --verbose
  stay-active -i 30s -d 2h30m -v
  stay-active --interval 2 --duration 45`,
	Run: runAutoPress,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&interval, "interval", "i", "1m", "How often to press a key (e.g., 30s, 5m, 1h, 1h30m, or raw number in minutes)")
	rootCmd.Flags().StringVarP(&duration, "duration", "d", "1h30m", "How long to run the application (e.g., 30s, 5m, 1h, 1h30m, or raw number in minutes)")
	rootCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "Enable verbose output")
}

func runAutoPress(cmd *cobra.Command, args []string) {
	parsedInterval, err := parseDuration(interval)
	if err != nil {
		fmt.Printf("Error parsing interval: %v\n", err)
		return
	}

	parsedDuration, err := parseDuration(duration)
	if err != nil {
		fmt.Printf("Error parsing duration: %v\n", err)
		return
	}

	durationUntil := time.Now().Add(parsedDuration)

	printf("Starting with %s interval and %s duration, running until %s\n",
		parsedInterval, parsedDuration, durationUntil.Format(time.TimeOnly))

	ticker := time.NewTicker(parsedInterval)
	defer ticker.Stop()

	timer := time.NewTimer(parsedDuration)
	defer timer.Stop()

	println("Press Ctrl+C to stop early...")

	for {
		select {
		case <-ticker.C:
			randomLetter := string(rune(rand.IntN(26) + 'a'))
			robotgo.TypeStr(randomLetter)

		case <-timer.C:
			println("Duration completed. Stopping auto-press.")
			return
		}
	}
}

// parseDuration parses a duration string in format like "1h30m50s" or raw number (treated as minutes)
func parseDuration(durationStr string) (time.Duration, error) {
	if matched, _ := regexp.MatchString(`^\d+$`, durationStr); matched {
		durationInMinutes, err := strconv.Atoi(durationStr)
		if err != nil {
			return 0, err
		}
		return time.Duration(durationInMinutes) * time.Minute, nil
	}

	return time.ParseDuration(strings.ToLower(durationStr))
}

// printf calls fmt.Printf if verbose is enabled
func printf(format string, a ...any) {
	if !verbose {
		return
	}
	fmt.Printf(format, a...)
}

// println calls fmt.Println if verbose is enabled
func println(a ...any) {
	if !verbose {
		return
	}
	fmt.Println(a...)
}
