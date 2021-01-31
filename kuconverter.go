package numtoword

import (
	"fmt"
	"strconv"
)

// KuConverter is used for converting number to kurdish words
func KuConverter(origNum float64) (finalResult string) {
	var result string

	level := []string{"", "هەزار", "ملیۆن", "ملیار", " هەزار ملیار"}
	if origNum == 0 {
		finalResult = "سفر"
		return
	}

	//converting original float number from  float to uint
	ipart := int64(origNum)

	k := 0
	num := uint(ipart)

	for num > 0 {
		threeDigit := num % 1000
		num = num / 1000
		result = kuHundred(threeDigit)
		if result != "" {
			finalResult = result + level[k] + " و " + finalResult
		}
		k++

	}
	//    string([]rune(s)[2:9])

	//len := len(finalResult)
	//finalResult = string([]rune(finalResult)[0 : len(finalResult)-3])
	finalResult = finalResult[0 : len(finalResult)-3]

	//checking if number has cents.. if it has cents will will take the cents and append that to the final result
	if !(origNum == float64(int64(origNum))) {
		decpart := fmt.Sprintf("%.2g", origNum-float64(ipart))[2:]

		//now we convert the decpart back to uint
		u64, err := strconv.ParseUint(decpart, 10, 32)
		if err != nil {
			fmt.Println(err)
		}
		cents := uint(u64)

		result = kuHundred(cents)

		finalResult = finalResult + " و " + result + " سەنت "
	}
	return
}

func kuHundred(num uint) (result string) {
	if num == 0 {
		return
	}

	var nums []TitleNum
	if num > 99 {
		tmpH := num / 100
		tmpTitle := TitleNum{
			Title: " سەد",
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

	nMaps[1] = "یەک"
	nMaps[2] = "دوو"
	nMaps[3] = "سێ"
	nMaps[4] = "چوار"
	nMaps[5] = "پێنج"
	nMaps[6] = "شەش"
	nMaps[7] = "حەوت"
	nMaps[8] = "هەشت"
	nMaps[9] = "نۆ"
	nMaps[10] = "دە"
	nMaps[11] = "یانزە"
	nMaps[12] = "دوانزە"
	nMaps[13] = "سیانزە"
	nMaps[14] = "چواردە"
	nMaps[15] = "پانزدە"
	nMaps[16] = "شانزە"
	nMaps[17] = "حەڤدە"
	nMaps[18] = "هەژدە"
	nMaps[19] = "نۆزدە"
	nMaps[20] = "بیست"
	nMaps[30] = "سی"
	nMaps[40] = "چل"
	nMaps[50] = "پەنجا"
	nMaps[60] = "شەست"
	nMaps[70] = "حەفتا"
	nMaps[80] = "هەشتا"
	nMaps[90] = "نەوەد"

	for _, v := range nums {
		if v.Num == 0 {
			continue
		}
		result += " و " + nMaps[v.Num] + v.Title
	}

	result = string([]rune(result)[3:])
	//result =result[5:0]
	return

}
