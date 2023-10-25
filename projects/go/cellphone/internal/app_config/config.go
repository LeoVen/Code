package app_config

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const VAR_APITYPE = "CELL_APITYPE"
const VAR_DBTYPE = "CELL_DBTYPE"
const VAR_REPOTYPE = "CELL_REPOTYPE"

type Main struct {
	DbType   int
	RepoType int
	ApiType  int
	Flags    map[string]string
}

var (
	EnvMissing = errors.New("Missing required variables:")
)

// Gets all configs from CLI and then tries to get the required args from ENV if they are missing
func GetConfig(required map[string]bool) (result Main, err error) {
	required[VAR_APITYPE] = true
	required[VAR_DBTYPE] = true
	required[VAR_REPOTYPE] = true

	result.Flags = make(map[string]string)

	for _, arg := range os.Args[1:] {
		kvs := strings.Split(arg, "=")
		if len(kvs) == 2 {
			result.Flags[kvs[0]] = kvs[1]
		}
	}

	var missing []string

	for name, v := range required {
		if !v {
			continue
		}

		val, ok := result.Flags[name]

		if !ok || val == "" {
			// If not in args, get from env
			val = os.Getenv(name)

			if val == "" {
				missing = append(missing, name)
			}

			result.Flags[name] = val
		}
	}

	if len(missing) != 0 {
		var buff bytes.Buffer
		buff.WriteString(EnvMissing.Error())
		buff.WriteByte('\n')
		buff.WriteByte('\t')
		for _, miss := range missing {
			buff.WriteString(" ")
			buff.WriteString(miss)
		}
		buff.WriteByte('\n')
		buff.WriteString("Note: These variables can be passed by CLI or ENV")
		err = errors.New(buff.String())
		return
	}

	apiType, err := strconv.Atoi(result.Flags[VAR_APITYPE])
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to parse variable %s: %s", VAR_APITYPE, err.Error()))
	}

	dbType, err := strconv.Atoi(result.Flags[VAR_DBTYPE])
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to parse variable %s: %s", VAR_DBTYPE, err.Error()))
	}

	repoType, err := strconv.Atoi(result.Flags[VAR_REPOTYPE])
	if err != nil {
		return result, errors.New(fmt.Sprintf("Failed to parse variable %s: %s", VAR_REPOTYPE, err.Error()))
	}

	result.ApiType = apiType
	result.DbType = dbType
	result.RepoType = repoType

	return result, nil
}
