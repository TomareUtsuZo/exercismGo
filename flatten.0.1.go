/*
This is my first version. Looking for feedback on the
readability, and professionalism of the code.
*/
package flatten
func Flatten(nested interface{}) (result []interface{}) {
    result = make([]interface{},0)
    switch value := nested.(type) {
        case []interface{}:
        for _, item := range value {
            result = append(result, Flatten(item)...) 
        }
        default:
        if value != nil {
        	result = append(result, value)
        }
    }
    return result
}
