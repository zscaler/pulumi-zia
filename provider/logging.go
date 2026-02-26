package provider

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"

	sdkLogger "github.com/zscaler/zscaler-sdk-go/v3/logger"
)

// providerLogger implements the SDK's logger.Logger interface, writing to
// stderr (always) and optionally to a file specified by ZSCALER_SDK_LOG_FILE.
// Unlike the SDK's built-in defaultLogger which writes to os.Stdout,
// this avoids corrupting Pulumi's gRPC channel on stdout.
type providerLogger struct {
	underlying *log.Logger
	verbose    bool
}

func (l *providerLogger) Printf(format string, v ...interface{}) {
	trimmed := strings.TrimSpace(format)
	if (strings.HasPrefix(trimmed, "[DEBUG]") || strings.HasPrefix(trimmed, "[TRACE]")) && !l.verbose {
		return
	}
	l.underlying.Printf(format, v...)
}

var (
	logWriter    io.Writer
	logSetupMu   sync.Mutex
	logSetupDone bool
)

// SetupProviderLogging configures the Go standard logger and returns an
// SDK-compatible logger. Safe to call multiple times; only the first call
// performs the setup.
//
// When ZSCALER_SDK_LOG=true:
//   - All log output goes to stderr (visible with `pulumi up --logtostderr -v=9`)
//   - If ZSCALER_SDK_LOG_FILE is also set, output is tee'd to that file
//
// When ZSCALER_SDK_LOG is not true, returns a no-op logger.
func SetupProviderLogging() sdkLogger.Logger {
	logSetupMu.Lock()
	defer logSetupMu.Unlock()

	enabled, _ := strconv.ParseBool(os.Getenv("ZSCALER_SDK_LOG"))
	if !enabled {
		return sdkLogger.NewNopLogger()
	}

	verbose, _ := strconv.ParseBool(os.Getenv("ZSCALER_SDK_VERBOSE"))

	if !logSetupDone {
		writers := []io.Writer{os.Stderr}

		if logPath := os.Getenv("ZSCALER_SDK_LOG_FILE"); logPath != "" {
			if !filepath.IsAbs(logPath) {
				if cwd, err := os.Getwd(); err == nil {
					logPath = filepath.Join(cwd, logPath)
				}
			}
			f, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
			if err == nil {
				writers = append(writers, f)
			} else {
				log.Printf("[WARN] could not open log file %q: %v", logPath, err)
			}
		}

		logWriter = io.MultiWriter(writers...)
		log.SetOutput(logWriter)
		logSetupDone = true
	}

	return &providerLogger{
		underlying: log.New(logWriter, "[oneapi-logger] ", log.LstdFlags|log.Lshortfile),
		verbose:    verbose,
	}
}
