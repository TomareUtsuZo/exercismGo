package wordy
import (
    //"fmt"
    "errors"
    "strconv"
)
func Answer(question string) (int, bool) {
    isBadlyAskedQuestion := func(question string) bool {
        isBadlyAskedQuestion := false
        if question[:8] != "What is " {
            isBadlyAskedQuestion = true
        }
        if question[len(question)-1] != '?' {
            isBadlyAskedQuestion = true
        }
        return isBadlyAskedQuestion
    }
    leadingNumberAndStringWithoutN := func(question string) (int, string, error) {
        sumOfDigits := 0
        cutPoint := 0
        makeNeg := false
        // if for some reason, there is nothing to read, error out.
        if len(question) == 0 {
            return 0, "", errors.New("leadingNumberAndStringWithoutN: no # to draw from")
        }
		for i, char := range question {
            isANumber := char >= '0' && char <= '9'
            if isANumber {
                sumOfDigits = (sumOfDigits * 10) + int(char-'0')
            } else if char == '-'{ // but if is negNumber
            	    makeNeg = true
            } else { // number found (probably) assume cutPoint for slicing
                cutPoint = i + 1
                break
            }
            cutPoint++
        } // finish for negative numbers
        if makeNeg {
            sumOfDigits *= -1
        } // if zero, no numbers found, error out
        if sumOfDigits == 0 {
            return 0, "", errors.New("noNumberFound")
        }
        return sumOfDigits, question[cutPoint:], nil
    }
    operationAndStringWithOutOp := func(question string) (string, string, error) {
        cutPoint := 0
		for i, char := range question {
            isANumber := char >= '0' && char <= '9'
            isNumberOrNegSign := isANumber || char == '-'
            if isNumberOrNegSign {
                cutPoint = i
                break
            }
        }
        if cutPoint == 0 {
            return "", "", errors.New("operationNotFollowedByNumber")
        }
        operation := question[:cutPoint-1]
        remainingQuestion := question[cutPoint:]
        return operation, remainingQuestion, nil
    }
    findNumbersAndOperationFirst := func(question string) (int, string, int, string, error) {
        leadingNumber, operationalPartOfQ, err := leadingNumberAndStringWithoutN(question)
        operation, operationalPartOfQ, err := operationAndStringWithOutOp(operationalPartOfQ)
        trailingNumber, operationalPartOfQ, err := leadingNumberAndStringWithoutN(operationalPartOfQ)
        return leadingNumber, operation, trailingNumber, operationalPartOfQ, err
    }
    continueFindingNumbersAndOperations := func(operationalPartOfQ string) (string, string, int, error){
            operation, operationalPartOfQ, err := operationAndStringWithOutOp(operationalPartOfQ)
            trailingNumber, operationalPartOfQ, err := leadingNumberAndStringWithoutN(operationalPartOfQ)
        return operation, operationalPartOfQ, trailingNumber, err
    }
    doMath := func(operation string, leadingNumber, trailingNumber int) (int, error) {
        switch operation {
            case "plus":
            	return leadingNumber + trailingNumber, nil
            case "minus":
            	return leadingNumber - trailingNumber, nil
            case "multiplied by":
            	return leadingNumber * trailingNumber, nil
            case "divided by":
            	return leadingNumber / trailingNumber, nil
        }
        return 0, errors.New("badOperation")
    }
    //test input, return fail if is not correctly asked and valid
    if isBadlyAskedQuestion(question) {
        return 0, false
    }
    output := 0
    leadingNumber := 0
    trailingNumber := 0
    operation := ""
    // remove "What is "
    operationalPartOfQ := question[8:len(question)-1]
    // test that q isn't "What is N?"
    output, err := strconv.Atoi(operationalPartOfQ)
    // if q wasn't that, then do operations to find what the problem is
    // otherwise go to end and return N
    if err != nil {
        // start trying to find full set of to do maths with
        leadingNumber, operation, trailingNumber, operationalPartOfQ, err = findNumbersAndOperationFirst(operationalPartOfQ)
        output, err = doMath(operation, leadingNumber, trailingNumber)
        // if that resulted in an error, exit w/ error report
            if err != nil {
                return -1, false
            }
        // if there is more problems remaining, 
        // do operations until error met, or operations complete.
        operationsRemain := operationalPartOfQ != ""
        for operationsRemain {
            operation, operationalPartOfQ, trailingNumber, err = continueFindingNumbersAndOperations(operationalPartOfQ)
        	output, err = doMath(operation, output, trailingNumber)
            if err != nil {
                return -1, false
            }
        // perpare for end of looptest
        operationsRemain = operationalPartOfQ != ""
        }
    }
    return output, true
}
