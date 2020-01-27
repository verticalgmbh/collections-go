package coll

// AddToWhereString filters a string collection for items which match a predicate
//
// **Parameters**
//    - collection: collection to check with predicate
//    - predicate : predicate for items to match to be added to resultslice
//    - result    : result to add items which match the predicate
func AddToWhereString(collection []string, predicate func(string) bool, result *[]string) {
	for _, item := range collection {
		if predicate(item) {
			*result = append(*result, item)
		}
	}
}

// WhereString filters a string collection for items which match a predicate
//
// **Parameters**
//    - collection : collection to check with predicate
//    - predicate  : predicate for items to match to be added to resultslice
//
// **Returns**
//    - []string: all items which match the predicate
func WhereString(collection []string, predicate func(string) bool) []string {
	result := make([]string, 0)

	for _, item := range collection {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// DoWhereString filters a string collection for items which match a predicate and executes an action for them
//
// **Parameters**
//    - collection: collection to check with predicate
//    - predicate : predicate for items to match
//    - action    : action for matching items to execute
func DoWhereString(collection []string, predicate func(string) bool, action func(string)) {
	for _, item := range collection {
		if predicate(item) {
			action(item)
		}
	}
}

// AnyString determines whether any item in the collection matches the predicate
//
// **Parameters**
//   - collection: collection to check against predicate
//   - predicate : predicate for any item to match
//
// **Returns**
//   - bool: true if any string matches the predicate, false otherwise
func AnyString(collection []string, predicate func(string) bool) bool {
	for _, item := range collection {
		if predicate(item) {
			return true
		}
	}

	return false
}

// AllString determines whether all items in the collection match a predicate
//
// **Parameters**
//   - collection: collection to check against predicate
//   - predicate : predicate for all items to match to return true
//
// **Returns**
//   - bool: true if all items of the collection match the predicate, false otherwise
func AllString(collection []string, predicate func(string) bool) bool {
	for _, item := range collection {
		if !predicate(item) {
			return false
		}
	}

	return true
}
