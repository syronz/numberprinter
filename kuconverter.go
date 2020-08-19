package numtoword

// TitleNum seperate name and numbers title
type TitleNum struct {
	Num   uint
	Title string
}

// KuConverter is used for converting number to kurdish words
func KuConverter(num uint) (finalResult string) {
	level := []string{"", " هەزار", " ملیۆن", " ملیار", " هەزار ملیار"}
	if num == 0 {
		finalResult = "سفر"
		return
	}

	k := 0
	for num > 0 {
		threeDigit := num % 1000
		num = num / 1000
		result := kuHundred(threeDigit)
		if result != "" {
			finalResult = result + level[k] + " و " + finalResult
		}
		k++

	}

	finalResult = finalResult[0 : len(finalResult)-5]

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

	result = result[5:]

	return

}
