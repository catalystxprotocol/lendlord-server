package logs

import (
	"context"
	"errors"
	"fmt"
	"time"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type gormLogger struct {
	logger                *log.Logger
	SlowThreshold         time.Duration
	SourceField           string
	Colorful              bool
	SkipErrRecordNotFound bool
}

func NewGormLogger(logger *log.Logger) *gormLogger {
	return &gormLogger{
		logger:                logger,
		SlowThreshold:         1 * time.Millisecond,
		Colorful:              false,
		SkipErrRecordNotFound: true,
	}
}

func (l *gormLogger) LogMode(logger.LogLevel) logger.Interface {
	return l
}

func (l *gormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Infof(s, args)
}

func (l *gormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Warnf(s, args)
}

func (l *gormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	l.logger.WithContext(ctx).Errorf(s, args)
}

func (l *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	fields := log.Fields{
		"sql_file":           utils.FileWithLineNum(),
		"sql_err":            err,
		"sql_execution_time": fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6),
	}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if rows != -1 {
		fields["sql_rows"] = rows
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		fields[log.ErrorKey] = err
		l.logger.WithContext(ctx).WithFields(fields).Errorf("%s [%s]", sql, elapsed)
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
		fields["sql_slow_log"] = slowLog
		if rows != -1 {
			fields["sql_rows"] = rows
		}
		l.logger.WithFields(fields).Printf("SQL: %s", sql)
		return
	}

	l.logger.WithContext(ctx).WithFields(fields).Debugf("%s [%s]", sql, elapsed)
}
