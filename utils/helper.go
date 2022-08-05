package utils

import "go.mongodb.org/mongo-driver/bson"

// @TODO:: not clear, doesnt work for all fields.
func ToDoc(v interface{}) (doc *bson.D, err error) {
	data, err := bson.Marshal(v)

	if err != nil {
		return
	}

	err = bson.Unmarshal(data, &doc)
	return
}
