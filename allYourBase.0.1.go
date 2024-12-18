/*
I am looking feed back on the cleaness and professionalism of 
my code.
*/

package allyourbase

import (
    "errors"
    "math"
    "slices"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) (output []int, err error) {
    err = inputTesting(inputBase, inputDigits, outputBase)
    if err == nil {

        baseTenValue := convertToBaseTen(inputBase, inputDigits)

        output = convertToOutputBase(outputBase, baseTenValue)
    }
    
    return output, err
}

func convertToBaseTen(inputBase int, inputDigits []int) (baseTenValue int) {
    place := 0 
    startIndex := len(inputDigits) - 1
    for i:=startIndex; i>=0; i-- {
        baseTenValue += inputDigits[i] * int(math.Pow(float64(inputBase), float64(place)))
        place++ 
    }
    return baseTenValue
}

func convertToOutputBase(outputBase, baseTenValue int) (output []int) {
    remainder := baseTenValue % outputBase
    quotient := baseTenValue / outputBase
    output = append(output, remainder)
    for quotient > 0 {
        remainder = quotient % outputBase
        quotient = quotient / outputBase
    	output = append(output, remainder)
    }
    slices.Reverse(output)
    
    return output
}

func inputTesting(inputBase int, inputDigits []int, outputBase int) (err error) {
    if inputBase < 2 {
        return errors.New("input base must be >= 2")
    } 
    if outputBase < 2 {
        return errors.New("output base must be >= 2")
    } 
    for _, digit := range inputDigits {
        if digit < 0 || digit >= inputBase {
            return errors.New("all digits must satisfy 0 <= d < input base")
        }
	}
    return err
}
