/*
refactored to improve readability, and clarity
*/
package wordy

import (
    //"fmt"
    "errors"
    "strconv"
)

type Operation int

const (
    Plus Operation = iota
    Minus
    Multiply
    Divide
    Fail
)

func Answer(question string) (output int, success bool) {
	success = false
    var err error
    output = 0
    //test input, return fail if is not correctly asked and valid
    if isBadlyAskedQuestion(question) {
        return output, success
    }

    leadingNumber := 0
    trailingNumber := 0
    var opCode Operation
    // remove "What is "
    operationalPartOfQ := question[8:len(question)-1]

    // test that q isn't "What is N?"
    AtoiErr := errors.New("")
    output, AtoiErr = strconv.Atoi(operationalPartOfQ)

    // if q wasn't that, then do operations to find what the problem is
    // otherwise go to end and return N
    if AtoiErr != nil {
        // start trying to find full set of to do maths with
        leadingNumber, opCode, trailingNumber, operationalPartOfQ, err = findNumbersAndOperationFirst(operationalPartOfQ)
        if err != nil {
            return output, success
        }
        output, err = doMath(opCode, leadingNumber, trailingNumber)
        // if that resulted in an error, exit w/ error report
        if err != nil {
            return output, success
        }
        // if there is more problems remaining, 
        // do operations until error met, or operations complete.
        operationsRemain := operationalPartOfQ != ""
        for operationsRemain {
            opCode, operationalPartOfQ, trailingNumber, err = continueFindingNumbersAndOperations(operationalPartOfQ)
            output, err = doMath(opCode, output, trailingNumber)
            if err != nil {
                return output, success
            }
            // perpare for end of looptest
            operationsRemain = operationalPartOfQ != ""
        }
    }
    success = true
    return output, success
}

func isBadlyAskedQuestion(question string) bool {
    if question[:8] != "What is " || question[len(question)-1] != '?' {
        return true
    }
    return false
}    

func leadingNumberAndStringWithoutN(question string) (sumOfDigits int, remainsOfQuestion string, err error) {
    sumOfDigits = 0
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
    err = nil
    remainsOfQuestion = question[cutPoint:]
    return sumOfDigits, remainsOfQuestion, err
}

func operationAndStringWithOutOp(question string) (opCode Operation, remaininsOfQuestion string, err error) {
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
        return Fail, "", errors.New("operationNotFollowedByNumber")
    }
    err = nil
    opCode = returnOpCode(question[:cutPoint-1])
    remaininsOfQuestion = question[cutPoint:]
    return opCode, remaininsOfQuestion, err
}

func findNumbersAndOperationFirst(question string) (leadingNumber int, opCode Operation, trailingNumber int, remainsOfQuestion string, err error) {
    leadingNumber, remainsOfQuestion, err = leadingNumberAndStringWithoutN(question)
    if err != nil {
        return 0, Plus, 0, "", err
    }
    opCode, remainsOfQuestion, err = operationAndStringWithOutOp(remainsOfQuestion)
    if err != nil {
        return 0, Plus, 0, "", err
    }
    trailingNumber, remainsOfQuestion, err = leadingNumberAndStringWithoutN(remainsOfQuestion)
    if err != nil {
        return 0, Plus, 0, "", err
    }
    return leadingNumber, opCode, trailingNumber, remainsOfQuestion, err
}

func continueFindingNumbersAndOperations(question string) (opCode Operation, remainsOfQuestion string, trailingNumber int, err error) {
    opCode, remainsOfQuestion, err = operationAndStringWithOutOp(question)
    trailingNumber, remainsOfQuestion, err = leadingNumberAndStringWithoutN(remainsOfQuestion)
    return opCode, remainsOfQuestion, trailingNumber, err
}

func doMath(opCode Operation, leadingNumber, trailingNumber int) (int, error) {
    switch opCode {
        case Plus:
        return leadingNumber + trailingNumber, nil
        case Minus:
        return leadingNumber - trailingNumber, nil
        case Multiply:
        return leadingNumber * trailingNumber, nil
        case Divide:
        return leadingNumber / trailingNumber, nil
    }
    return 0, errors.New("badOperation")
}

func returnOpCode(opString string) (opCode Operation) {
    switch opString {
        case "plus":
        opCode = Plus
        case "minus":
        opCode = Minus
        case "multiplied by":
        opCode = Multiply
        case "divided by":
        opCode = Divide
        default:
        opCode = Fail
    }
    return opCode
}
