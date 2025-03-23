package logger

import (
	"context"
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/AmirHossein82x/doctor-appointment/internal/config"
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
)

// CustomFormatter enhances log output with file name & line number
type CustomFormatter struct {
	logrus.TextFormatter
}

// NewCustomFormatter creates a new CustomFormatter with default settings
func NewCustomFormatter() *CustomFormatter {
	return &CustomFormatter{
		TextFormatter: logrus.TextFormatter{
			DisableColors:   false,
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
			PadLevelText:    true,
			ForceQuote:      true,
		},
	}
}

// Format modifies the log entry format
func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// Add file and line information to the log entry
	file, line := getCaller(7)
	entry.Data["file"] = fmt.Sprintf("%s:%d", file, line)

	// Call the original TextFormatter to format the log entry
	return f.TextFormatter.Format(entry)
}

// getCaller retrieves the file and line number of the log call site
func getCaller(skip int) (string, int) {
	// Adjust the skip value based on your call stack depth
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", 0
	}
	// Shorten the file path to make it more readable
	return shortenFilePath(file), line
}

// shortenFilePath shortens the file path to make it more readable
func shortenFilePath(file string) string {
	// Split the file path and keep only the last two parts
	parts := strings.Split(file, "/")
	if len(parts) > 2 {
		return strings.Join(parts[len(parts)-2:], "/")
	}
	return file
}

// ElasticsearchHook struct to send logs to Elasticsearch
type ElasticsearchHook struct {
	client *elastic.Client
	index  string
}

// NewElasticsearchHook initializes the hook with authentication
func NewElasticsearchHook(esURL, index, username, password string) (*ElasticsearchHook, error) {
	client, err := elastic.NewClient(
		elastic.SetURL(esURL),
		elastic.SetBasicAuth(username, password),
		elastic.SetSniff(false),
		elastic.SetHealthcheck(true),
	)
	if err != nil {
		return nil, fmt.Errorf("elasticsearch connection error: %w", err)
	}

	fmt.Println("Connected to Elasticsearch successfully!")
	return &ElasticsearchHook{client: client, index: index}, nil
}

// Fire sends logs to Elasticsearch asynchronously
func (hook *ElasticsearchHook) Fire(entry *logrus.Entry) error {
	// Capture file and line information before spawning the goroutine
	file, line := getCaller(8)

	// Prepare log data
	logData := map[string]interface{}{
		"message":   entry.Message,
		"level":     entry.Level.String(),
		"timestamp": entry.Time.Format(time.RFC3339),
		"file":      fmt.Sprintf("%s:%d", file, line),
	}

	// Send log to Elasticsearch in a goroutine without blocking
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		_, err := hook.client.Index().
			Index(hook.index).
			BodyJson(logData).
			Do(ctx)
		if err != nil {
			fmt.Printf("Error sending log to Elasticsearch: %v\n", err)
		}
	}()

	return nil
}

// Levels defines which log levels should be sent to Elasticsearch
func (hook *ElasticsearchHook) Levels() []logrus.Level {
	return logrus.AllLevels // Send all levels to Elasticsearch
}

// SetUpLogger configures and returns a new logger instance
func SetUpLogger() *logrus.Logger {
	config := config.LoadConfig()
	log := logrus.New()
	log.SetFormatter(NewCustomFormatter()) // Use the optimized custom formatter
	log.SetOutput(os.Stdout)

	esHook, err := NewElasticsearchHook(config.ELASTIC_ADDR, config.ELASTIC_INDEX, config.ELASTIC_USERNAME, config.ELASTIC_PASSWORD)
	if err != nil {
		fmt.Printf("Failed to initialize Elasticsearch hook: %v\n", err)
	} else {
		log.AddHook(esHook)
	}

	return log
}
