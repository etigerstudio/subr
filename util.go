// Copyright (c) 2020 E-Tiger Studio. All rights reserved.

package subr

import "errors"

func CastBucketToBytes(buckets map[string]interface{}, key string) (data []byte, err error)  {
	bytes, ok := buckets[key].([]byte)
	if !ok {
		return nil, errors.New("Bucket of key: " + key + " is not []byte")
	}
	return bytes, nil
}