/*
 * Author: Galit Miller
 *
 * Copyright (c) 2020 RemoteCourthouse as an unpublished work.
 * All rights reserved.
 *
 * The information contained herein is confidential property of RemoteCourthouse.
 * The use, copying, transfer or disclosure of such
 * information is prohibited except by express written agreement with RemoteCourthouse.
 */

package common

import (
	"log"
	"os"
)

var (
	infolog    *log.Logger = log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	warninglog *log.Logger = log.New(os.Stdout, "WARNING: ", log.Ldate|log.Ltime)
	errlog     *log.Logger = log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
)

// Info logger
func Info(m string, ms ...interface{}) { infolog.Printf(m, ms...) }

// Warn logger
func Warn(m string, ms ...interface{}) { warninglog.Printf(m, ms...) }

// Error logger
func Error(m string, ms ...interface{}) { errlog.Printf(m, ms...) }
