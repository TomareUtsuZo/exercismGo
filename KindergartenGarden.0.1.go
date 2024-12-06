package kindergarten

import (
    "errors"
)

// Define the Garden type here.
type Garden struct {
    Diagram string
    Children []string 
}


func NewGarden(diagram string, children []string) (*Garden, error) {
    var garden Garden
    garden.Diagram = diagram
    garden.Children = append([]string{}, children...)
    if inputIsBad(garden.Diagram, garden.Children) {
        return nil, errors.New("inputBad")
    }
    return &garden, nil
}

func (g *Garden) Plants(child string) (childsGarden []string, err bool) {
    err = false
    mapOfPlants := map[rune]string {
        'R': "radishes",
        'C': "clover",
        'G': "grass",
        'V': "violets",
    }
    workingChildren := g.alphabetized()
    indexOfChild := theChildIsAtIndex(workingChildren, child)
    // Not in list
    if indexOfChild == -1 {
        return workingChildren, err
    }
    childsGarden = make([]string, 4)
    
    newLineOffset := 1 // deal w/ \n
    beginingOfSecondRow := len(workingChildren) * 2 + newLineOffset
    firstRowStart := (indexOfChild * 2) + newLineOffset
    secondRowStart := beginingOfSecondRow + (indexOfChild * 2) + newLineOffset
    
    frontRow := g.Diagram[firstRowStart:firstRowStart+2]
    backRow := g.Diagram[secondRowStart:secondRowStart+2]
    childsGarden[0] = mapOfPlants[rune(frontRow[0])]
    childsGarden[1] = mapOfPlants[rune(frontRow[1])]
    childsGarden[2] = mapOfPlants[rune(backRow[0])]
    childsGarden[3] = mapOfPlants[rune(backRow[1])]
	err = true
    return childsGarden, err 
}

func (g *Garden)alphabetized() (childrenOut []string) {
    childrenOut = g.Children
    
    indexBeingTestedFor := 0
	currentCandidateName := childrenOut[indexBeingTestedFor]
    
    indexCurrentCandidate := 0
    namePushedRight := ""
    swapNeedsDoingFlag := false
    
    i := indexBeingTestedFor + 1
    for i < len(childrenOut) {
        if childrenOut[i][0] <= currentCandidateName[0] {
            indexShorterThanBothNames := func(nameA, nameB string, index int) bool {
                indexLessThanA := index < len(nameA)
                indexLessThanB := index < len(nameB)
                return indexLessThanA && indexLessThanB
            }
            isCandidateEarlier := func(childrenOut, currentCandidateName string, j int) bool {
                return childrenOut[0] < currentCandidateName[0] || childrenOut[j] < currentCandidateName[j]
            }
            for j:=1;indexShorterThanBothNames(childrenOut[i], currentCandidateName, j) ;j++ {
                if isCandidateEarlier(childrenOut[i], currentCandidateName, j) {
                	currentCandidateName = childrenOut[i]
                    indexCurrentCandidate = i
                    swapNeedsDoingFlag = true
                }
            }
        }
        if swapNeedsDoingFlag {
            namePushedRight = childrenOut[indexBeingTestedFor]
            childrenOut[indexBeingTestedFor] = currentCandidateName
            childrenOut[indexCurrentCandidate] = namePushedRight
            //reset after
            swapNeedsDoingFlag = false
            indexBeingTestedFor++
            currentCandidateName = childrenOut[indexBeingTestedFor]
    		i = indexBeingTestedFor
        }
        i++
    }
    return childrenOut
}

func inputIsBad(diagram string, children []string) bool {
    noChildren := children == nil
    diagramNotCorrectLen := len(diagram) != len(children) * 4 + 2
    namesRepeat := false
    isntLower := false
    for i, name := range children {
        for j:=i+1;j<len(children);j++ {
            if children[j] == name {
                namesRepeat = true
                break
            }
        }
    }
    for _, char := range diagram {
        if char != 'R' && char != 'C' && char != 'G' && char != 'V' && char != '\n'{
            isntLower = true
        }
    }
    return noChildren || diagramNotCorrectLen || namesRepeat || isntLower
}

func theChildIsAtIndex(workingChildren []string, child string) (index int) {
    index = -1
    for i, gChild := range workingChildren {
        if child == gChild {
            index = i
        }
    }
    return index
}
