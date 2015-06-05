package main
 
import ("fmt"
	 "io/ioutil"
)	
 
func backTrack(n, count int, loc []byte) int{
	if count < 0{
		return 0
	}
	if loc[n] == 91{		// [
		count--
	}else if loc[n] == 93{		// ]
		count++
	}
	return backTrack(n - 1, count, loc) + 1
	
}
 
func forwardTrack(n, count int, loc []byte) int{
	if count < 0{
		return 0
	}
	if loc[n] == 91{		// [
		count++
	}else if loc[n] == 93{		// ]
		count--
	}
	return forwardTrack(n + 1, count, loc) + 1
	
}
 
func main(){
	fmt.Print("Program directory: ")
	var fileDir string
	fmt.Scanf("%s", &fileDir)
 
	prog, _ := ioutil.ReadFile(fileDir)
 
	mem := make([]byte, 30000)
	thing := 0
	var temp string
	for i := 0; i < len(prog); i++{
		switch string(prog[i]){
			case ">": 
				thing++
			case "<":
				thing--
			case "+":
				mem[thing]++
			case "-":
				mem[thing]--
			case ".":
				fmt.Print(string(mem[thing]))
			case ",":
				fmt.Scanf("%s", &temp)
				mem[thing] = byte(temp[0])
			case "[":
				if mem[thing] == 0{
					i += forwardTrack(i + 1, 0, prog)
				}
			case "]":
				if mem[thing] != 0{
					i -= backTrack(i - 1, 0, prog)
				}
			}
			
	}
}

