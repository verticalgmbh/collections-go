package coll

import (
	"reflect"
)

// Any determines whether any item in the collection matches the predicate
//
// **Parameters**
//   - collection: collection to iterate
//   - predicate : predicate for any item to match for this function to return true
//
// **Returns**
//   - bool: true if any item of the collection matches the predicate, false otherwise
func Any(collection interface{}, predicate func(interface{}) bool) bool {
	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			if predicate(array.Index(i).Interface()) {
				return true
			}
		}
	default:
		return predicate(collection)
	}

	return false
}

// Interface converts data to an interface array
//
// **Parameters**
//   - collection: collection to convert
//
// **Returns**
//   - []interface{}: interface array containing all items of original collection
func Interface(collection interface{}) []interface{} {
	var result []interface{}

	datatype := reflect.TypeOf(collection)

	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			result = append(result, array.Index(i).Interface())
		}
	default:
		result = append(result, collection)
	}

	return result
}

// AddToWhere filters a collection for items which match a predicate
//
// **Parameters**
//    - collection : collection to check with predicate
//    - predicate  : predicate for items to match to be added to resultslice
//    - resultslice: slice where items matching the predicate are stored. This has to be a pointer to a slice, else this method will crash
func AddToWhere(collection interface{}, predicate func(interface{}) bool, resultslice interface{}) {
	resultptr := reflect.ValueOf(resultslice)
	result := resultptr.Elem()

	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			item := array.Index(i)
			if predicate(item.Interface()) {
				result.Set(reflect.Append(result, item))
			}
		}
	default:
		if predicate(collection) {
			result.Set(reflect.Append(result, reflect.ValueOf(collection)))
		}
	}
}

// Where filters a collection for items which match a predicate
//
// **Parameters**
//    - collection : collection to check with predicate
//    - predicate  : predicate for items to match to be added to resultslice
//
// **Returns**
//    - []interface{}: slice containing items matching predicate
func Where(collection interface{}, predicate func(interface{}) bool) []interface{} {
	var result []interface{} = make([]interface{}, 0)

	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			item := array.Index(i)
			if predicate(item.Interface()) {
				result = append(result, item.Interface())
			}
		}
	default:
		if predicate(collection) {
			result = append(result, collection)
		}
	}

	return result
}

// DoWhere filters a collection for items which match a predicate and executes an action for them
//
// **Parameters**
//    - collection: collection to check with predicate
//    - predicate : predicate for items to match to be added to resultslice
//    - action    : action for matching items to execute
func DoWhere(collection interface{}, predicate func(interface{}) bool, action func(interface{})) {
	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			item := array.Index(i)
			if predicate(item.Interface()) {
				action(item.Interface())
			}
		}
	default:
		if predicate(collection) {
			action(collection)
		}
	}
}

// FirstOrDefault returns the first item which matches the predicate or nil if no item matches the predicate
//
// **Parameters**
//   - collection: collection to iterate
//   - predicate : predicate for item to match to be returned
//
// **Returns**
//   - interface{}: first item which matches the predicate or nil otherwise
func FirstOrDefault(collection interface{}, predicate func(interface{}) bool) interface{} {
	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			item := array.Index(i).Interface()
			if predicate(item) {
				return item
			}
		}
	default:
		if predicate(collection) {
			return collection
		}
	}

	return nil
}

// All determines whether all items in the collection match a predicate
//
// **Parameters**
//   - collection: collection of which to check items
//   - predicate : predicate for items to match
//
// **Returns**
//   - bool: true if all items match the specified predicate, false otherwise
func All(collection interface{}, predicate func(interface{}) bool) bool {
	datatype := reflect.TypeOf(collection)
	switch datatype.Kind() {
	case reflect.Array, reflect.Slice, reflect.Chan, reflect.Map:
		array := reflect.ValueOf(collection)
		for i := 0; i < array.Len(); i++ {
			if !predicate(array.Index(i).Interface()) {
				return false
			}
		}
	default:
		return predicate(collection)
	}

	return true
}
