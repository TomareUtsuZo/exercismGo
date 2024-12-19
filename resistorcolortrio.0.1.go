/*
First pass at this code. My goal is readability, and
proffesionalism of code. If there is something that is
not readable, or would be written for greater proffesionalism
please let me know.
*/
package resistorcolortrio
import (
    "fmt"
)
type Color int
const (
        Black Color = iota
        Brown
        Red
        Orange
        Yellow
        Green
        Blue
        Violet
        Grey
        White
)
func Label(colors []string) (formatedOutput string) {
    firstTwoBandValue, prefixColor := colorToInt(colors)
    
    var resistance int
    var prefix string
    if firstTwoBandValue > 0 {
    	resistance, prefix = formatResistance(firstTwoBandValue, prefixColor)
    }
    
    formatedOutput = fmt.Sprintf("%v %vohms", resistance, prefix)
    
    return formatedOutput
}
func formatResistance(leadingDigits int, color Color) (resistance int, prefix string) {
    prefixes := []string{"", "kilo", "mega", "giga"}
    resistance = leadingDigits
    for i:=0; i<int(color);i++ {
        resistance = resistance * 10        
    }
    prefixIndex := 0
    for resistance % 1000 == 0 && prefixIndex < len(prefixes) {
        resistance = resistance / 1000
        prefixIndex++
    }
    return resistance, prefixes[prefixIndex]
}
