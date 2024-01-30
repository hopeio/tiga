package log

import (
	"fmt"
	"github.com/hopeio/tiga/utils/log/output"
	neti "github.com/hopeio/tiga/utils/net"
	"github.com/hopeio/tiga/utils/slices"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
	"strconv"
	"time"
)

func init() {
	output.RegisterSink()
	SetDefaultLogger(&Config{Development: true, Caller: true, Level: zapcore.DebugLevel})
}

var (
	Default     *Logger
	skipLoggers = make([]*Logger, 10)
)

func SetDefaultLogger(lf *Config) {
	Default = lf.NewLogger()
}

func GetSkipLogger(skip int) *Logger {
	if skip > 10 {
		panic("最大不超过10")
	}
	if skipLoggers[skip] == nil {
		skipLoggers[skip] = Default.AddSkip(skip)
	}
	return skipLoggers[skip]
}

type Logger struct {
	*zap.Logger
}

type ZapConfig zap.Config

type Config struct {
	Development     bool          `json:"development,omitempty"`
	Caller          bool          `json:"caller,omitempty"`
	Level           zapcore.Level `json:"level,omitempty"`
	EncodeLevelType string        `json:"encodeLevelType,omitempty"`
	OutputPaths     OutPutPaths   `json:"outputPaths"`
	ModuleName      string        `json:"moduleName,omitempty"` //系统名称namespace.service
	zapcore.EncoderConfig
}

func (lc *Config) Init() {

	if !lc.Development && lc.ModuleName != "" && lc.NameKey == "" {
		lc.NameKey = "module"
		lc.EncodeName = zapcore.FullNameEncoder
	}

	if lc.TimeKey == "" {
		lc.TimeKey = "time"
		lc.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006/01/02 15:04:05.000"))
		}
	}

	if lc.LevelKey == "" {
		lc.LevelKey = "level"
		if lc.EncodeLevelType == "" {
			if lc.Development {
				lc.EncodeLevelType = "capitalColor"
			}
		}
		var el zapcore.LevelEncoder
		el.UnmarshalText([]byte(lc.EncodeLevelType))
		lc.EncodeLevel = el
	}

	if lc.CallerKey == "" {
		lc.CallerKey = "caller"
		if lc.Development {
			lc.EncodeCaller = zapcore.FullCallerEncoder
		} else {
			lc.EncodeCaller = func(caller zapcore.EntryCaller, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(caller.TrimmedPath())
			}
		}
	}
	if !lc.Development && lc.FunctionKey == "" {
		lc.FunctionKey = "func"
	}
	if lc.MessageKey == "" {
		lc.MessageKey = "msg"
	}
	if lc.StacktraceKey == "" {
		lc.StacktraceKey = "stack"
	}
	if lc.LineEnding == "" {
		lc.LineEnding = zapcore.DefaultLineEnding
	}

	if lc.ConsoleSeparator == "" {
		lc.ConsoleSeparator = "\t"
	}

	lc.EncodeDuration = func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(strconv.FormatInt(d.Nanoseconds()/1e6, 10) + "ms")
	}
}

type OutPutPaths struct {
	Console []string `json:"console,omitempty"`
	Json    []string `json:"json,omitempty"`
}

// 初始化日志对象
func (lc *Config) NewLogger(cores ...zapcore.Core) *Logger {
	logger := lc.initLogger(cores...)
	// 不是测试环境要加主机名和ip
	if !lc.Development {
		hostname, _ := os.Hostname()
		logger = logger.With(
			zap.String("hostname", hostname),
			zap.String("ip", neti.GetIP()),
		)
	}

	return &Logger{logger}
}

// Named adds a sub-scope to the logger's name. See Logger.Named for details.
func (l *Logger) Named(name string) *Logger {
	return &Logger{l.Logger.Named(name)}
}

// WithOptions warp the zap WithOptions, applies the supplied Options, and
// returns the resulting Logger. It's safe to use concurrently.
func (l *Logger) WithOptions(opts ...zap.Option) *Logger {
	return &Logger{l.Logger.WithOptions(opts...)}
}

// With warp the zap With. Fields added
// to the child don't affect the parent, and vice versa.
func (l *Logger) With(fields ...zap.Field) *Logger {
	return &Logger{l.Logger.With(fields...)}
}

// Sugar warp the zap Sugar.
func (l *Logger) Sugar() *zap.SugaredLogger {
	l.WithOptions(zap.AddCallerSkip(-1))
	return l.Logger.Sugar()
}

// AddCore warp the zap AddCore.
func (l *Logger) AddCore(newCore zapcore.Core) *Logger {
	return l.WithOptions(zap.WrapCore(func(core zapcore.Core) zapcore.Core {
		return zapcore.NewTee(core, newCore)
	}))
}

// AddSkip warp the zap AddCallerSkip.
func (l *Logger) AddSkip(skip int) *Logger {
	return &Logger{l.Logger.WithOptions(zap.AddCallerSkip(skip))}
}

// 构建日志对象基本信息
func (lc *Config) initLogger(cores ...zapcore.Core) *zap.Logger {
	lc.Init()

	var consoleEncoder, jsonEncoder zapcore.Encoder

	if len(lc.OutputPaths.Console) > 0 {
		consoleEncoder = zapcore.NewConsoleEncoder(lc.EncoderConfig)
		// 如果输出同时有stdout和stderr,那么warn级别及以下的用stdout,error级别及以上的用stderr
		stdout, stderr := false, false
		consolePaths := make([]string, 0, len(lc.OutputPaths.Console))
		slices.ForEachIndex(lc.OutputPaths.Console, func(i int) {
			if lc.OutputPaths.Console[i] == "stdout" {
				stdout = true
			} else if lc.OutputPaths.Console[i] == "stderr" {
				stderr = true
			} else {
				consolePaths = append(consolePaths, lc.OutputPaths.Console[i])
			}
		})
		if stdout && stderr {
			cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), StdOutLevel(lc.Level)),
				zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), StdErrLevel(lc.Level)))
		} else {
			if stdout {
				consolePaths = append(consolePaths, "stdout")
			}
			if stderr {
				consolePaths = append(consolePaths, "stderr")
			}
		}
		sink, _, err := zap.Open(consolePaths...)
		if err != nil {
			log.Fatal(err)
		}
		cores = append(cores, zapcore.NewCore(consoleEncoder, sink, lc.Level))
	}

	if len(lc.OutputPaths.Json) > 0 {
		lc.EncodeLevel = zapcore.LowercaseLevelEncoder
		jsonEncoder = zapcore.NewJSONEncoder(lc.EncoderConfig)
		sink, _, err := zap.Open(lc.OutputPaths.Json...)
		if err != nil {
			log.Fatal(err)
		}
		cores = append(cores, zapcore.NewCore(jsonEncoder, sink, lc.Level))
	}
	//如果没有设置输出，默认控制台
	if len(cores) == 0 {
		consoleEncoder = zapcore.NewConsoleEncoder(lc.EncoderConfig)
		cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), StdOutLevel(lc.Level)),
			zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stderr), StdErrLevel(lc.Level)))
	}

	core := zapcore.NewTee(cores...)

	logger := zap.New(core, lc.hook()...)
	if lc.ModuleName != "" {
		logger = logger.Named(lc.ModuleName)
	}
	return logger
}

func (lc *Config) hook() []zap.Option {
	var hooks []zap.Option

	if lc.Development {
		hooks = append(hooks, zap.Development(), zap.AddStacktrace(zapcore.DPanicLevel))
	}
	hooks = append(hooks, zap.AddCallerSkip(1))
	if lc.Caller {
		hooks = append(hooks, zap.AddCaller())
	}
	return hooks
}

func Sync() error {
	return Default.Sync()
}

func Print(args ...interface{}) {
	if ce := Default.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Debug(args ...interface{}) {
	if ce := Default.Check(zap.DebugLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Info(args ...interface{}) {
	if ce := Default.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Warn(args ...interface{}) {
	if ce := Default.Check(zap.WarnLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Error(args ...interface{}) {
	if ce := Default.Check(zap.ErrorLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Panic(args ...interface{}) {
	if ce := Default.Check(zap.PanicLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Fatal(args ...interface{}) {
	if ce := Default.Check(zap.FatalLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func Printf(template string, args ...interface{}) {
	if ce := Default.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Debugf(template string, args ...interface{}) {
	if ce := Default.Check(zap.DebugLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Infof(template string, args ...interface{}) {
	if ce := Default.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Warnf(template string, args ...interface{}) {
	if ce := Default.Check(zap.WarnLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Errorf(template string, args ...interface{}) {
	if ce := Default.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Panicf(template string, args ...interface{}) {
	if ce := Default.Check(zap.PanicLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func Fatalf(template string, args ...interface{}) {
	if ce := Default.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Printf(template string, args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// 兼容gormv1
func (l *Logger) Print(args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Debug uses fmt.Sprint to construct and log a message.
func (l *Logger) Debug(args ...interface{}) {
	if ce := l.Check(zap.DebugLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Info uses fmt.Sprint to construct and log a message.
func (l *Logger) Info(args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Warn uses fmt.Sprint to construct and log a message.
func (l *Logger) Warn(args ...interface{}) {
	if ce := l.Check(zap.WarnLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Error uses fmt.Sprint to construct and log a message.
func (l *Logger) Error(args ...interface{}) {
	if ce := l.Check(zap.ErrorLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanic(args ...interface{}) {
	if ce := l.Check(zap.DPanicLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (l *Logger) Panic(args ...interface{}) {
	if ce := l.Check(zap.PanicLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (l *Logger) Fatal(args ...interface{}) {
	if ce := l.Check(zap.FatalLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// Debug logs a message at DebugLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Debugw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.DebugLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Info logs a message at InfoLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Infow(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.InfoLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Warn logs a message at WarnLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Warnw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.WarnLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Error logs a message at ErrorLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
func (l *Logger) Errorw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.ErrorLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// DPanic logs a message at DPanicLevel. The message includes any fields
// passed at the log site, as well as any fields accumulated on the logger.
//
// If the logger is in development mode, it then panics (DPanic means
// "development panic"). This is useful for catching errors that are
// recoverable, but shouldn't ever happen.
func (l *Logger) DPanicw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.DPanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Panic logs a message at PanicLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then panics, even if logging at PanicLevel is disabled.
func (l *Logger) Panicw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.PanicLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Fatal logs a message at FatalLevel. The message includes any fields passed
// at the log site, as well as any fields accumulated on the logger.
//
// The logger then calls os.Exit(1), even if logging at FatalLevel is
// disabled.
func (l *Logger) Fatalw(msg string, fields ...zap.Field) {
	if ce := l.Check(zap.FatalLevel, msg); ce != nil {
		ce.Write(fields...)
	}
}

// Debugf uses fmt.Sprintf to log a templated message.
func (l *Logger) Debugf(template string, args ...interface{}) {
	if ce := l.Check(zap.DebugLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Infof uses fmt.Sprintf to log a templated message.
func (l *Logger) Infof(template string, args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Warnf uses fmt.Sprintf to log a templated message.
func (l *Logger) Warnf(template string, args ...interface{}) {
	if ce := l.Check(zap.WarnLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Errorf uses fmt.Sprintf to log a templated message.
func (l *Logger) Errorf(template string, args ...interface{}) {
	if ce := l.Check(zap.ErrorLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (l *Logger) DPanicf(template string, args ...interface{}) {
	if ce := l.Check(zap.DPanicLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (l *Logger) Panicf(template string, args ...interface{}) {
	if ce := l.Check(zap.PanicLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (l *Logger) Fatalf(template string, args ...interface{}) {
	if ce := l.Check(zap.FatalLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

// 兼容grpclog
func (l *Logger) Infoln(args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Warning(args ...interface{}) {
	if ce := l.Check(zap.WarnLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Warningln(args ...interface{}) {
	if ce := l.Check(zap.WarnLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Warningf(template string, args ...interface{}) {
	if ce := l.Check(zap.WarnLevel, fmt.Sprintf(template, args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Errorln(args ...interface{}) {
	if ce := l.Check(zap.ErrorLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Fatalln(args ...interface{}) {
	if ce := l.Check(zap.FatalLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

func (l *Logger) Println(args ...interface{}) {
	if ce := l.Check(zap.InfoLevel, fmt.Sprint(args...)); ce != nil {
		ce.Write()
	}
}

// grpclog
func (l *Logger) V(level int) bool {
	level -= 2
	return l.Logger.Core().Enabled(zapcore.Level(level))
}

// sugar
const (
	_oddNumberErrMsg    = "Ignored key without a value."
	_nonStringKeyErrMsg = "Ignored key-value pairs with non-string keys."
)

func (l *Logger) log(lvl zapcore.Level, template string, fmtArgs []interface{}, context []interface{}) {
	// If logging at this level is completely disabled, skip the overhead of
	// string formatting.
	if lvl < zap.DPanicLevel && !l.Core().Enabled(lvl) {
		return
	}

	// Format with Sprint, Sprintf, or neither.
	msg := template
	if msg == "" && len(fmtArgs) > 0 {
		msg = fmt.Sprint(fmtArgs...)
	} else if msg != "" && len(fmtArgs) > 0 {
		msg = fmt.Sprintf(template, fmtArgs...)
	}

	if ce := l.Check(lvl, msg); ce != nil {
		ce.Write(l.sweetenFields(context)...)
	}
}

func (l *Logger) sweetenFields(args []interface{}) []zap.Field {
	if len(args) == 0 {
		return nil
	}

	// Allocate enough space for the worst case; if users pass only structured
	// fields, we shouldn't penalize them with extra allocations.
	fields := make([]zap.Field, 0, len(args))
	var invalid invalidPairs

	for i := 0; i < len(args); {
		// This is a strongly-typed field. Consume it and move on.
		if f, ok := args[i].(zap.Field); ok {
			fields = append(fields, f)
			i++
			continue
		}

		// Make sure this element isn't a dangling key.
		if i == len(args)-1 {
			l.DPanic(_oddNumberErrMsg, zap.Any("ignored", args[i]))
			break
		}

		// Consume this value and the next, treating them as a key-value pair. If the
		// key isn't a string, add this pair to the slice of invalid pairs.
		key, val := args[i], args[i+1]
		if keyStr, ok := key.(string); !ok {
			// Subsequent errors are likely, so allocate once up front.
			if cap(invalid) == 0 {
				invalid = make(invalidPairs, 0, len(args)/2)
			}
			invalid = append(invalid, invalidPair{i, key, val})
		} else {
			fields = append(fields, zap.Any(keyStr, val))
		}
		i += 2
	}

	// If we encountered any invalid key-value pairs, log an error.
	if len(invalid) > 0 {
		l.DPanic(_nonStringKeyErrMsg, zap.Array("invalid", invalid))
	}
	return fields
}

type invalidPair struct {
	position   int
	key, value interface{}
}

func (p invalidPair) MarshalLogObject(enc zapcore.ObjectEncoder) error {
	enc.AddInt64("position", int64(p.position))
	zap.Any("key", p.key).AddTo(enc)
	zap.Any("value", p.value).AddTo(enc)
	return nil
}

type invalidPairs []invalidPair

func (ps invalidPairs) MarshalLogArray(enc zapcore.ArrayEncoder) error {
	var err error
	for i := range ps {
		err = multierr.Append(err, enc.AppendObject(ps[i]))
	}
	return err
}

type StdOutLevel zapcore.Level

func (l StdOutLevel) Enabled(lvl zapcore.Level) bool {
	return lvl >= zapcore.Level(l) && lvl >= zapcore.WarnLevel
}

type StdErrLevel zapcore.Level

func (l StdErrLevel) Enabled(lvl zapcore.Level) bool {
	return lvl >= zapcore.Level(l) && lvl < zapcore.WarnLevel
}
