package search_int

import (
	"sort"
	"strconv"
)

type IntCount struct {
	num int
	cnt int
}

//排序 Start
type Lists []IntCount

func (l Lists) Len() int { return len(l) }

func (l Lists) Swap(i, j int) { l[i], l[j] = l[j], l[i] }

type OrderByNum struct{ Lists }

func (s OrderByNum) Less(i, j int) bool { return s.Lists[i].num < s.Lists[j].num }
//排序 End

//GetIntCount 获取字符串中int和出现次数以及排序输出
//optString := "ee1039ccc254ad222321qaads22321341wwwq1039a"
//输出 [{254 1} {1039 2} {222321 1} {22321341 1}]
func GetIntCount(optString string) (intCount []IntCount) {
	strCnt := len(optString)
	//对总的字符串进行遍历
	for i := 0;i<strCnt;i++ {
		runeString := rune(optString[i])

		var intStart = i
		var intString string
		//判断ascii值是否是数字字符
		if IsInt(runeString){
			//从出现的第一个数字字符遍历
			for ;intStart <strCnt;intStart++ {
				i = intStart
				runeIntString := rune(optString[intStart])
				//对每一个字符进行判断是否是数字字符，找到第一个不是数字字符的位置并进行操作
				if IsInt(runeIntString) {
					//如果是数字字符，则与前面的数字字符串进行拼接
					intString = intString + string(runeIntString)
				}else {
					//如果不是数字字符，则进行 append 数组，并且break到上层循环，并从该下标位置进行继续查找，此时下标位置在intStart
					numInt, _ := strconv.Atoi(intString)
					index := InArrayString(numInt,intCount)
					if index >= 0{
						//如果此数字字符串已经出现过，则进行加1操作
						intCount[index].cnt++
					}else {
						//如果没有出现过，则进行append操作
						intCount = append(intCount, IntCount{
							num : numInt,
							cnt: 1,
						})
					}
					break
				}
			}
		}
	}

	//进行排序操作
	sort.Sort(OrderByNum{intCount})
	return intCount
}

//InArrayString 判断是否存在
func InArrayString(idle int,array []IntCount) int {
	for i ,v := range array {
		if idle == v.num {
			return i
		}
	}
	return -1
}

//IsInt 判断是否是int类型
func IsInt(s rune) bool {
	return s >= 48 && s <= 57
}

