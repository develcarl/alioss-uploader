package operate

import (
	"strconv"
	"time"
)

var (
	bucketName string = "typora-pic-plat"
)

func PutObject(path string) (string, error) {
	bucket, err := GetExistBucket(bucketName)
	var uuid string = GenUUID()
	var now time.Time = time.Now()
	var remotePath = "typora/" + strconv.Itoa(now.Local().Year()) + strconv.Itoa(int(now.Local().Month())) + strconv.Itoa(now.Local().Day()) + "/" + uuid
	if err != nil {
		HandleError(err)
	}
	return remotePath, bucket.PutObjectFromFile(remotePath, path)
}
