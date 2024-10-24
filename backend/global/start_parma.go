package global

import "os"

var (
	EvRpcPort        string
	EvRpcKey         string
	TmpFileStorePath string
	Debug            bool
)

func GetTmpFileStorePath() string {
	if _, err := os.Stat(TmpFileStorePath); os.IsNotExist(err) {
		os.MkdirAll(TmpFileStorePath, os.ModePerm)
	}
	return TmpFileStorePath
}
