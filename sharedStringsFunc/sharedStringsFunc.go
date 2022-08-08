package sharedStringsFunc

import ("strings"
"fmt"
"errors"
"strconv"
)
// ||(!DoesContain(lMsg,"withdraw"))|| (!DoesContain(lMsg,"transfer"))


func SplitString(msg string) (string , string, error){

	var lMsg string= strings.ToLower(msg)
	fmt.Println(lMsg)
	if (DoesContain(lMsg,"deposit"))||(DoesContain(lMsg,"withdraw"))|| (DoesContain(lMsg,"transfer")){
		
		split:=strings.Split(lMsg, " £")
		if len(split) == 1{
			
			return "", "", errors.New("wrong input")
		}
	return split[0], split[1], nil
		
	}else{
		return "", "", errors.New("wrong input")
	}
	
}

func SplitAtSpace(msg string) (string, string){
	split:=strings.Split(msg, " ")
	return split[0], split[1]
}

func DoesContain(msg string, word string) bool{
	return strings.Contains(msg, word)
}

func StringToFloat(msg string)(float64, error){
	amount, err := strconv.ParseFloat(msg, 64)
	if err != nil {
		return 0.00, errors.New("can not convert your number")
	}
	return amount, nil 
}