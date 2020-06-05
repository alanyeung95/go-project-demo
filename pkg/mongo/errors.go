package mongo

import (
	"go.mongodb.org/mongo-driver/mongo"
)

func isDuplicateKeyErr(err error) (bool, error) {
	const duplicateKeyErrCode = 11000

	if wes, ok := err.(mongo.WriteErrors); ok {
		for _, we := range wes {
			if we.Code == duplicateKeyErrCode {
				return true, we
			}
		}
	}
	return false, nil
}
