package utils

func Count(s []string, position int) (string, string){

	oneCount :=0
	zeroCount :=0

	for _, v := range s{

		if string(v[position]) == "0"{
			zeroCount++
		}else {
			oneCount ++
		}
	}

	if oneCount>zeroCount{
		return "1", "0"
	} else {
		return "0", "1"
	}
}

