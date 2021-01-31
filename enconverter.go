package numtoword

import (
	"fmt"
	"strconv"
)

// EnConverter is used for converting number to kurdish words
func EnConverter(origNum float64) (finalResult string) {
	var result string
	if origNum == 0 {
		finalResult = "zero"
		return
	}

	//converting original float number from  float to uint
	ipart := int64(origNum)

	level := []string{"", " thousand", " million", " billion", " trilion"}
	// if num <= -1 {

	// }
	k := 0
	num := uint(ipart)
	for num > 0 {
		threeDigit := num % 1000
		num = num / 1000
		result := enHundred(threeDigit)

		if result != "" {
			finalResult = result + level[k] + " and " + finalResult
		}
		k++

	}

	finalResult = finalResult[0 : len(finalResult)-5]

	//checking if number has cents.. if it has cents will will take the cents and append that to the final result
	if !(origNum == float64(int64(origNum))) {
		decpart := fmt.Sprintf("%.2g", origNum-float64(ipart))[2:]

		//now we convert the decpart back to uint
		u64, err := strconv.ParseUint(decpart, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		cents := uint(u64)

		result = enHundred(cents)

		finalResult = finalResult + " and " + result + " cents"
	}
	return
}

func enHundred(num uint) (result string) {
	if num == 0 {
		return
	}

	var nums []TitleNum
	if num > 99 {
		tmpH := num / 100
		tmpTitle := TitleNum{
			Title: " hundred",
			Num:   tmpH,
		}
		num = num % 100
		nums = append(nums, tmpTitle)
	}

	var smlTitle TitleNum
	if num > 20 {
		tmpRem := num % 10
		smlTitle = TitleNum{
			Title: "",
			Num:   tmpRem,
		}
		num = num / 10 * 10
	}

	tmpTitle := TitleNum{
		Title: "",
		Num:   num,
	}
	nums = append(nums, tmpTitle)
	if smlTitle.Num != 0 {
		nums = append(nums, smlTitle)
	}

	// fmt.Println(nums)

	nMaps := make(map[uint]string, 9)

	nMaps[1] = "one"
	nMaps[2] = "two"
	nMaps[3] = "three"
	nMaps[4] = "four"
	nMaps[5] = "five"
	nMaps[6] = "six"
	nMaps[7] = "seven"
	nMaps[8] = "eight"
	nMaps[9] = "nine"
	nMaps[10] = "ten"
	nMaps[11] = "eleven"
	nMaps[12] = "twenteen"
	nMaps[13] = "thirteen"
	nMaps[14] = "fourteen"
	nMaps[15] = "fifteen"
	nMaps[16] = "sixteen"
	nMaps[17] = "seventeen"
	nMaps[18] = "eighteen"
	nMaps[19] = "nineteen"
	nMaps[20] = "twenty"
	nMaps[30] = "thirty"
	nMaps[40] = "forty"
	nMaps[50] = "fifty"
	nMaps[60] = "sixty"
	nMaps[70] = "seventy"
	nMaps[80] = "eighty"
	nMaps[90] = "ninety"

	for _, v := range nums {
		if v.Num == 0 {
			continue
		}
		result += " and " + nMaps[v.Num] + v.Title
	}

	result = result[5:]

	return

}
