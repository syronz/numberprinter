package numtoword

import (
	"fmt"
	"strconv"
)

// TitleNum separate name and numbers title
type TitleNum struct {
	Num       uint
	Title     string
	condition bool
}

// ArConverter is used for converting number to kurdish words
func ArConverter(origNum float64) (finalResult string) {
	var result string

	if origNum == 0 {
		finalResult = "صفر"
		return
	}

	//converting original float number from  float to uint
	ipart := int64(origNum)
	//level := []string{"", " تريليون", "مليار", " مليون", " ألف"}
	level := make([]string, 4)
	level = append(level, "")
	level = append(level, "ألف ")
	level = append(level, "مليون ")
	level = append(level, "مليار ")
	level = append(level, "تريليون ")

	k := 0
	num := uint(ipart)

	for num > 0 {

		threeDigit := num % 1000

		num = num / 1000
		//	tmp := num + 1
		result = arHundred(threeDigit)

		if result != "" {
			switch k {
			case 1:
				fmt.Println("case 1 ..the number is: ", threeDigit)
				if threeDigit == 1 {
					result = ""
					level[k] = "ألف"
				} else if threeDigit == 2 {
					result = ""
					level[k] = "ألفين"
				} else if threeDigit > 2 {
					level[k] = "آلاف"
				}
			case 2:
				//num++
				fmt.Println("2 the number is ", num)

				if threeDigit == 1 {
					result = ""
					level[k] = "مليون"
				} else if threeDigit == 2 {
					result = ""
					level[k] = "مليونان"
				} else if threeDigit > 2 {
					level[k] = "ملايين"

				}
			case 3:
				//num++
				fmt.Println("case 3the number is ", num)

				if threeDigit == 1 {
					result = ""
					level[k] = "مليار"
				} else if threeDigit == 2 {
					result = ""
					level[k] = "ملياري"
				} else if threeDigit > 2 {
					level[k] = "مليارات"

				}
			}
			if k != 0 {
				finalResult = result + level[k] + " و " + finalResult

			} else {
				finalResult = result + " و " + finalResult

			}

			switch k {
			case 1:
				level[k] = "ألف"
			case 2:
				level[k] = "مليون"
			case 3:
				level[k] = "مليار"

			}
		}

		k++

	}

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

		switch cents {
		case 1, 2, 3, 4, 5, 6, 7, 8, 9:
			switch decpart[0] {
			case '1', '2', '3', '4', '5', '6', '7', '8', '9':
				cents *= 10
			}
		}

		result = arHundred(cents)

		finalResult = finalResult + " و " + result + " سنتًا "
	}
	return
}

func arHundred(num uint) (result string) {
	if num == 0 {
		return ""
	}

	var nums []TitleNum
	//checking the hundreads
	if num > 99 {
		tmpH := num / 100
		fmt.Println("tmpH is :::::", tmpH)
		hundredTitle := "مائة "
		var cond bool
		switch tmpH {
		case 1:
			cond = true
		case 2:
			cond = true
			hundredTitle = "مائتان "

		}
		tmpTitle := TitleNum{
			Title:     hundredTitle,
			Num:       tmpH,
			condition: cond,
		}
		num = num % 100
		nums = append(nums, tmpTitle)
	}

	var smlTitle TitleNum
	//checking from 20 to 99 for last digit
	if num > 20 {
		tmpRem := num % 10
		smlTitle = TitleNum{
			Title: "",
			Num:   tmpRem,
		}
		num = num / 10 * 10
	}

	//now we store it because in arabic last digit comes before first,
	//such as fifty two : إثنان و خمسون
	if smlTitle.Num != 0 {
		nums = append(nums, smlTitle)

	}

	//now we store the first digit
	tmpTitle := TitleNum{
		Title: "",
		Num:   num,
	}
	nums = append(nums, tmpTitle)

	// fmt.Println(nums)

	nMaps := make(map[uint]string, 9)

	nMaps[1] = "واحد"
	nMaps[2] = "اثنان"
	nMaps[3] = "ثلاثة"
	nMaps[4] = "أربعة"
	nMaps[5] = "خمسة"
	nMaps[6] = "ستة"
	nMaps[7] = "سبعة"
	nMaps[8] = "ثمانية"
	nMaps[9] = "تسعة"
	nMaps[10] = "عشرة"
	nMaps[11] = "إحدى عشر"
	nMaps[12] = "إثنا عشر"
	nMaps[13] = "ثلاثة عشر"
	nMaps[14] = "أربعة عشرة"
	nMaps[15] = "خمسة عشر"
	nMaps[16] = "السادس عشر"
	nMaps[17] = "سبعة عشر"
	nMaps[18] = "الثامنة عشر"
	nMaps[19] = "تسعة عشر"
	nMaps[20] = "عشرون"
	nMaps[30] = "ثلاثون"
	nMaps[40] = "أربعون"
	nMaps[50] = "خمسون"
	nMaps[60] = "ستون"
	nMaps[70] = "سبعون"
	nMaps[80] = "ثمانون"
	nMaps[90] = "تسعون"

	for _, v := range nums {
		if v.Num == 0 {
			continue
		}
		if v.condition {
			result += " و " + v.Title
			continue
		}
		result += " و " + nMaps[v.Num] + " " + v.Title
	}

	result = string([]rune(result)[3:])

	return

}
