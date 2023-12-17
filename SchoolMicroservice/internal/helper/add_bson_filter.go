package helper

import (
	"go.mongodb.org/mongo-driver/bson"
	"strings"
)

func AddFilter(bsonFilter []bson.E, key, op string, values []string) []bson.E {
	switch op {
	case "eq":
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$in", values}}})
	case "ne":
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$nin", values}}})
	case "starts":
		var regexPatterns []string
		for _, prefix := range values {
			regexPatterns = append(regexPatterns, "^"+prefix)
		}
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$regex", strings.Join(regexPatterns, "|")}}})
	case "ends":
		var regexPatterns []string
		for _, suffix := range values {
			regexPatterns = append(regexPatterns, suffix+"$")
		}
		bsonFilter = append(bsonFilter, bson.E{Key: key, Value: bson.D{{"$regex", strings.Join(regexPatterns, "|")}}})
	}

	return bsonFilter
}
