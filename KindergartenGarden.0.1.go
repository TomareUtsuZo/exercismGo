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
func (g *Garden) Plants(child string) ([]string, bool) {
    mapOfPlants := map[rune]string {
        'R': "radishes",
        'C': "clover",
        'G': "grass",
        'V': "violets",
    }
    workingChildren := alphabetize(g.Children)
    indexOfChild := -1
    childsGarden := make([]string, 4)
    for i, gChild := range workingChildren {
        if child == gChild {
            indexOfChild = i
        }
    }
    // Not in list
    if indexOfChild == -1 {
        return workingChildren, false
    }
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
    return childsGarden, true 
}
func alphabetize(children []string) []string {
    indexBeingTestedFor := 0
	mostFirstestName := children[indexBeingTestedFor]
    mostFirstestIndex := 0
    namePushedRight := ""
    swapNeedsDoing := false
    i := indexBeingTestedFor + 1
    for i < len(children) {
        if children[i][0] <= mostFirstestName[0] {
            indexShorterThanBothNames := func(nameA, nameB string, index int) bool {
                indexLessThanA := index < len(nameA)
                indexLessThanB := index < len(nameB)
                return indexLessThanA && indexLessThanB
            }
            for j:=1;indexShorterThanBothNames(children[i], mostFirstestName, j) ;j++ {
                if children[i][j] < mostFirstestName[j] || children[i][0] < mostFirstestName[0] {
                	mostFirstestName = children[i]
                    mostFirstestIndex = i
                    swapNeedsDoing = true
                }
            }
        }
        if swapNeedsDoing {
            namePushedRight = children[indexBeingTestedFor]
            children[indexBeingTestedFor] = mostFirstestName
            children[mostFirstestIndex] = namePushedRight
            //reset after
            swapNeedsDoing = false
            indexBeingTestedFor++
            mostFirstestName = children[indexBeingTestedFor]
    		i = indexBeingTestedFor
        }
        i++
    }
    return children
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
