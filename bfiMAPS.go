package main

import "fmt"
import "io/ioutil"

type bfM struct{
	mem [30000]byte		// Define memory of 30000 bytes
	prgD []byte		// Slice of program data
	pPtr, dPtr int		// Pointers to program data and 'sand box' memory
}

func (m *bfM) incPtr(){
	m.dPtr++		// Increment sand box pointer
}

func (m *bfM) decPtr(){
	m.dPtr--		// Decrement sand box pointer
}

func (m *bfM) incD(){
	m.mem[m.dPtr]++		// Increment data that the sand box pointer is pointing to
}

func (m *bfM) decD(){
	m.mem[m.dPtr]--		// Decrement data that the sand box pointer is pointing to
}

func (m *bfM) printD(){
	fmt.Print(string(m.mem[m.dPtr]))	// Print the character the sand box pointer is pointing to
}

func (m *bfM) scanD(){
	inTemp := ""
	fmt.Scanf("%s", &inTemp)		// Read a character from stdin... 
	m.mem[m.dPtr] = byte(inTemp[0])		// ...and store it in memory where the sand box pointer is pointing to
}

func (m *bfM) begL(){
	count := 0				// Bracket counter
	if m.mem[m.dPtr] == 0{
		m.pPtr++			// Skip entry bracket
		for count >= 0{			// Loop until count == -1, meaning top-most bracket as been reached
			if  m.prgD[m.pPtr] == 91{		// [
				count++		// Increment count for every open bracket
			}else if m.prgD[m.pPtr] == 93{		// ]
				count--		// Decrement count for every open bracket
			}
			m.pPtr++		// Prepare next instruction for inspection
		}
		m.pPtr--
	}
}

func (m *bfM) endL(){
	count := 0				// Operates similar to endL function
	if m.mem[m.dPtr] != 0{
		m.pPtr--
		for count >= 0{
			if  m.prgD[m.pPtr] == 91{		// [
				count--
			}else if m.prgD[m.pPtr] == 93{		// ]
				count++
			}
			m.pPtr--
		}
		m.pPtr++
	}
}

func main(){
	myMachine := new(bfM)	// Create new BrainFuck machine

	// The map below maps strings to functions
	oper := map[string]func(){">": myMachine.incPtr, 
				"<": myMachine.decPtr, 
				"+": myMachine.incD, 
				"-": myMachine.decD,
				".": myMachine.printD,
				",": myMachine.scanD,
				"[": myMachine.begL,
				"]": myMachine.endL,
				}

	temp := ""		// Define temporary string for program directory
	fmt.Print("Program directory: ")
	fmt.Scanf("%s", &temp)
	myMachine.prgD, _ = ioutil.ReadFile(temp)

	for myMachine.pPtr = 0; myMachine.pPtr < len(myMachine.prgD); myMachine.pPtr++{
		doOp := oper[string(myMachine.prgD[myMachine.pPtr])]
		if (doOp != nil){	// If a valid instruction is read from program data...
			doOp()		// Do operation
		}
	}
}
