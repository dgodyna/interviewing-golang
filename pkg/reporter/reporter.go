// Package reporter provides performance tracking and statistics collection utilities
// for the network event generator. This package enables benchmarking and performance
// analysis by recording execution metrics and persisting them for analysis.
package reporter

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// ExecutionStatistic represents a single performance measurement record
// containing timing information and event generation metrics.
//
// This structure is used to track and analyze performance improvements
// during optimization efforts, allowing candidates to measure the impact
// of their changes quantitatively.
type ExecutionStatistic struct {
	// ExecutionStart records the timestamp when event generation began
	ExecutionStart time.Time `json:"execution_start"`

	// Duration contains the total time taken for event generation
	Duration time.Duration `json:"duration"`

	// NumbOfEvents specifies how many events were generated in this run
	NumbOfEvents int `json:"number_of_events"`
}

// Save serializes an ExecutionStatistic record to JSON and appends it to the specified file.
//
// This function enables continuous performance tracking by accumulating multiple
// benchmark runs in a single file. Each record is written as a separate JSON line,
// making the file format compatible with line-delimited JSON processors.
//
// The function handles file creation if it doesn't exist and appends to existing files,
// allowing for long-term performance trend analysis across optimization iterations.
//
// Parameters:
//   - stat: The ExecutionStatistic record to save
//   - filename: Path to the output file where the record will be appended
//
// File Format:
//
//	Each line contains a complete JSON object representing one execution record.
//	Example:
//	  {"execution_start":"2024-01-01T10:00:00Z","duration":5000000000,"number_of_events":1000000}
//	  {"execution_start":"2024-01-01T10:05:00Z","duration":450000000,"number_of_events":1000000}
func Save(stat ExecutionStatistic, filename string) error {
	// Serialize the statistic record to JSON format
	jsonData, err := json.Marshal(stat)
	if err != nil {
		return fmt.Errorf("failed to marshal ExecutionStatistic: %w", err)
	}

	// Append newline to create line-delimited JSON format
	jsonData = append(jsonData, '\n')

	// Open file in append mode, create if it doesn't exist
	// Use 0644 permissions (owner read/write, group/others read-only)
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	// Write the JSON record to the file
	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("failed to write to file %s: %w", filename, err)
	}

	return nil
}

// GetAllStatistics reads all ExecutionStatistic records from the specified file.
//
// This function parses a line-delimited JSON file containing performance records
// and returns them as a slice of ExecutionStatistic structs. It's designed to
// work with files created by the Save function, enabling performance analysis
// and trend tracking across multiple benchmark runs.
func GetAllStatistics(filename string) ([]ExecutionStatistic, error) {
	// Attempt to open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		// If file doesn't exist, return empty slice (not an error condition)
		if os.IsNotExist(err) {
			return []ExecutionStatistic{}, nil
		}
		// Other file access errors should be reported
		return nil, fmt.Errorf("failed to open file %s: %w", filename, err)
	}
	defer file.Close()

	var statistics []ExecutionStatistic
	scanner := bufio.NewScanner(file)

	// Process each line in the file
	for scanner.Scan() {
		line := scanner.Text()

		// Skip empty lines
		if len(line) == 0 {
			continue
		}

		// Attempt to parse the JSON line
		var stat ExecutionStatistic
		if err := json.Unmarshal([]byte(line), &stat); err != nil {
			// Skip malformed lines but continue processing
			// This allows for mixed content or partial corruption recovery
			continue
		}

		statistics = append(statistics, stat)
	}

	// Check for scanner errors (I/O issues during reading)
	if err := scanner.Err(); err != nil {
		return statistics, fmt.Errorf("error reading file %s: %w", filename, err)
	}

	return statistics, nil
}

// SaveAndReport saves the current execution statistic and displays performance analysis.
//
// This function combines performance tracking with immediate feedback by:
// 1. Saving the current run statistics to the specified file
// 2. Analyzing historical performance for the same event count
// 3. Displaying current performance and improvement trends
func SaveAndReport(stat ExecutionStatistic, filename string) error {
	// First, read all existing statistics for analysis (before adding the current run)
	allStats, err := GetAllStatistics(filename)
	if err != nil {
		return fmt.Errorf("failed to read statistics for analysis: %w", err)
	}

	// Save the current statistic
	if err := Save(stat, filename); err != nil {
		return fmt.Errorf("failed to save statistic: %w", err)
	}

	fmt.Println("================================================================")

	// Display current run information
	fmt.Printf("Current Run: %s events in %v\n",
		formatNumber(stat.NumbOfEvents), stat.Duration)

	// Find statistics for the same number of events
	var sameEventStats []ExecutionStatistic
	for _, s := range allStats {
		if s.NumbOfEvents == stat.NumbOfEvents {
			sameEventStats = append(sameEventStats, s)
		}
	}

	// If we have historical data for this event count, show comparisons
	if len(sameEventStats) > 1 {
		// Find first and last runs (excluding current)
		first := sameEventStats[0]
		var last ExecutionStatistic

		// Find the most recent previous run (second to last)
		if len(sameEventStats) >= 2 {
			last = sameEventStats[len(sameEventStats)-2]
		}

		// Calculate and display improvement vs first run
		if first.Duration > 0 {
			firstImprovement := calculateImprovement(first.Duration, stat.Duration)
			fmt.Printf("Comparin With First Run:       %v (%s improvement)\n", first.Duration, firstImprovement)
		}

		// Calculate and display improvement vs last run (if different from first)
		if len(sameEventStats) >= 2 && last.Duration > 0 && last.ExecutionStart != first.ExecutionStart {
			lastImprovement := calculateImprovement(last.Duration, stat.Duration)
			fmt.Printf("Comparin With Last  Run:       %v (%s improvement)\n", last.Duration, lastImprovement)
		}
	} else {
		fmt.Println("First run for this event count - no comparison data available")
	}

	fmt.Println("================================================================")
	return nil
}

// calculateImprovement computes the percentage improvement between two durations with color formatting.
//
// The function calculates how much faster the new duration is compared to the old duration.
// Positive percentages indicate improvement (faster execution) and are displayed in green.
// Negative percentages indicate regression (slower execution) and are displayed in red.
// Zero improvement is displayed without color.
//
// Formula: ((oldDuration - newDuration) / oldDuration) * 100
//
// ANSI Color Codes:
//   - Green: \033[32m (for improvements)
//   - Red: \033[31m (for regressions)
//   - Reset: \033[0m (return to default color)
//
// Parameters:
//   - oldDuration: The baseline duration for comparison
//   - newDuration: The current duration being evaluated
//
// Returns:
//   - string: Formatted percentage improvement with color (e.g., "\033[32m+1025%\033[0m", "\033[31m-15%\033[0m")
func calculateImprovement(oldDuration, newDuration time.Duration) string {
	if oldDuration == 0 {
		return "N/A"
	}

	// Calculate percentage improvement
	improvement := float64(oldDuration-newDuration) / float64(oldDuration) * 100

	// Format with the appropriate sign and color
	if improvement > 0 {
		// Green for positive improvements (faster execution)
		return fmt.Sprintf("\033[32m+%.0f%%\033[0m", improvement)
	} else if improvement < 0 {
		// Red for negative improvements (slower execution)
		return fmt.Sprintf("\033[31m%.0f%%\033[0m", improvement)
	}
	return "0%"
}

// formatNumber formats large numbers with comma separators for better readability.
//
// This function improves the display of event counts by adding a thousand separators,
// making large numbers easier to read and understand at a glance.
//
// Examples:
//
//	1000000 -> "1,000,000"
//	10000 -> "10,000"
//	500 -> "500"
//
// Parameters:
//   - n: The integer to format
//
// Returns:
//   - string: Formatted number with comma separators
func formatNumber(n int) string {
	if n < 1000 {
		return fmt.Sprintf("%d", n)
	}

	// Convert to string and add commas
	str := fmt.Sprintf("%d", n)
	result := ""

	for i, digit := range str {
		if i > 0 && (len(str)-i)%3 == 0 {
			result += ","
		}
		result += string(digit)
	}

	return result
}
