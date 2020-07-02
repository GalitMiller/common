/*
 * Author: Galit Miller
 */

package utils

import (
	"log"
	"os"
)

var (
	infolog    *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warninglog *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errlog     *log.Logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
)

// Infoo logger
func Info(m string, ms ...interface{}) { infolog.Printf(m, ms...) }

// Warn logger
func Warn(m string, ms ...interface{}) { warninglog.Printf(m, ms...) }

// Error logger
func Error(m string, ms ...interface{}) { errlog.Printf(m, ms...) }
