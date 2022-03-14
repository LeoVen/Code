package error

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// Default error type for every API response
type Error struct {
	errorList []errorData
}

type errorData struct {
	err      error
	callSite string
}

func NewErrorResponse(err ...error) Error {
	res := make([]errorData, 0)

	for _, e := range err {
		res = append(res, errorData{
			err: e,
			// 2 so it skips the getMetadata() and this function
			callSite: getMetadata(2),
		})
	}

	return Error{errorList: res}
}

func (self *Error) Error() string {
	sb := strings.Builder{}

	for i, data := range self.errorList {
		sb.WriteString(fmt.Sprintf("{%03d} %s:\n\t%s\n", i, data.callSite, data.err.Error()))
	}

	return sb.String()
}

// Returns the last added error, if any
func (self *Error) Last() error {
	l := len(self.errorList)

	if l == 0 {
		return nil
	}

	return self.errorList[l-1].err
}

func (self *Error) AddError(err error) {
	self.errorList = append(self.errorList, errorData{
		err: err,
		// 2 so it skips the getMetadata() and this function
		callSite: getMetadata(2),
	})
}

func getMetadata(skip int) string {
	pc, filename, line, _ := runtime.Caller(skip)

	return fmt.Sprintf("[%s] at %s [%s:%d]",
		time.Now().Format("2006-01-02 15:04:05 -0700 MST 2006"),
		filename,
		runtime.FuncForPC(pc).Name(),
		line)
}
