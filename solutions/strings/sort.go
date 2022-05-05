//https://leetcode-cn.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/submissions/

//这道题目比较坑,起始的时候默认了人名和时间都是有序的,实际的示例中并不是,需要再做额外的处理
package strings

import (
	"sort"
	"strconv"
)

func RunAlertNames() []string {
	return alertNames([]string{"a", "a", "a", "a", "a", "b", "b", "b", "b", "b", "b"},
		[]string{"23:20", "11:09", "23:30", "23:02", "15:28", "22:57", "23:40", "03:43", "21:55", "20:38", "00:19"})
}

// func alertNames(keyName []string, keyTime []string) []string {
// 	//假设数组按照人名分段,且每个人名的时间升序排列

// 	//第二次, 给出的示例中,名字可能不是按照升序排列的,最后需要对结果做下排序

// 	//第三次,给出的示例中,时间也不是有序的.......
// 	//做下预处理,把输入的两个数组数据转换为map格式,key为人名,value为时间的数组形式,value可以进行预排序
// 	var result = []string{}
// 	if len(keyName) == 0 {
// 		return result
// 	}
// 	var lastName = keyName[0]
// 	var startTime = keyTime[0]
// 	var times = 1
// 	for i := 1; i < len(keyName); i++ {
// 		if keyName[i] != lastName {
// 			lastName = keyName[i]
// 			startTime = keyTime[i]
// 			times = 1
// 		}
// 		if isSameSpan(startTime, keyTime[i]) {
// 			times++
// 		} else {
// 			startTime = keyTime[i]
// 			times = 1
// 		}
// 		if times >= 3 {
// 			if len(result) == 0 || result[len(result)-1] != keyName[i] {
// 				result = append(result, keyName[i])
// 			}
// 		}
// 	}
// 	sort.Strings(result)
// 	return result
// }
// func isSameSpan(start, end string) bool {
// 	startSlice := start[0:2]
// 	endSlice := end[0:2]
// 	if startSlice == endSlice {
// 		return true
// 	} else {
// 		s, _ := strconv.Atoi(startSlice)
// 		e, _ := strconv.Atoi(endSlice)
// 		return e-s == 1
// 	}
// }

//startTime按照人名的第一个时间,获取endTime

//最终的解法
//这道题其实不难,主要是一开始默认的给定示例比较理想,对题目的意思理解不太到位,包括给定两个数组的有序设定,对一个小时时间间隔的理解.导致在考虑解法的时候有些偏差
//修改过程中,实现思路和最初思路不太一样,导致修改的时候会存在小的细节点改动不到位
//所以,在开始考虑算法之前,需要对题目的意思理解清楚,多多考虑异常case,才能保证在一开始的时候考虑比较完善
// import (
// 	"sort"
// 	"strconv"
// )

// func alertNames(keyName []string, keyTime []string) []string {
// 	//假设数组按照人名分段,且每个人名的时间升序排列

// 	//第二次, 给出的示例中,名字可能不是按照升序排列的,最后需要对结果做下排序

// 	//第三次,给出的示例中,时间也不是有序的.......
// 	//做下预处理,把不同人名分到不同的数组中,然后对数组进行排序后遍历

// 	// keyName=[]string{"a","a","a","a","a","b","b","b","b","b","b"}
// 	// keyTime=[]string{"23:20","11:09","23:30","23:02","15:28","22:57","23:40","03:43","21:55","20:38","00:19"}

// 	//    keyName=[]string{"a","a","a","a","b","b","b","b","b","b","c","c","c","c"}
// 	// keyTime=[]string{"01:35","08:43","20:49","00:01","17:44","02:50","18:48","22:27","14:12","18:00","12:38","20:40","03:59","22:24"}

// 	//    keyName=[]string{"leslie","leslie","leslie","clare","clare","clare","clare"}
// 	// keyTime=[]string{"13:00","13:20","14:00","18:00","18:51","19:30","19:49"}

// 	//第四次修改: 时间不是按照整段划分的,而是两两相减; 另一个需要做变动的是,对startTime的修改
// 	var result = []string{}
// 	if len(keyName) == 0 {
// 		return result
// 	}
// 	var nameMap = make(map[string][]string) //key为人名,value 为时间序列
// 	for i := 0; i < len(keyName); i++ {
// 		_, ok := nameMap[keyName[i]]
// 		if !ok {
// 			nameMap[keyName[i]] = []string{keyTime[i]}
// 		} else {
// 			nameMap[keyName[i]] = append(nameMap[keyName[i]], keyTime[i])
// 		}
// 	}
// 	for key, value := range nameMap {
// 		sort.Strings(value)
// 		startTime := value[0]
// 		times := 1
// 		isAdded := false
// 		// fmt.Println("times=",value)
// 		for i := 1; i < len(value); i++ {
// 			if isSameSpan(startTime, value[i]) {
// 				// fmt.Println("in timespan:",startTime,value[i]               )
// 				times++
// 			} else {
// 				//重新定义startTime的时候,需要向前回溯到一个在一个时间间隔内的时间作为起始值
// 				startTime = value[i]
// 				times = 1
// 				for j := i - 1; j >= 0; j-- {
// 					if !isSameSpan(value[j], value[i]) {
// 						break
// 					}
// 					startTime = value[j]
// 					times++
// 					// fmt.Println("re starttime=",value[j])
// 				}
// 			}
// 			if times >= 3 && !isAdded {
// 				result = append(result, key)
// 				isAdded = true
// 			}
// 		}

// 	}
// 	// var lastName=keyName[0]
// 	// var startTime=keyTime[0]
// 	// var times=1
// 	// for i:=1;i<len(keyName);i++{
// 	//     if keyName[i]!=lastName{
// 	//         lastName=keyName[i]
// 	//         startTime=keyTime[i]
// 	//         times=1
// 	//     }
// 	//     if isSameSpan(startTime,keyTime[i]){
// 	//         times++
// 	//     }else{
// 	//         startTime=keyTime[i]
// 	//         times=1
// 	//     }
// 	//     if times>=3{
// 	//         if len(result)==0||result[len(result)-1]!=keyName[i]{
// 	//             result=append(result,keyName[i])
// 	//         }
// 	//     }
// 	// }
// 	sort.Strings(result)
// 	return result
// }
// func isSameSpan(start, end string) bool {
// 	// startSlice:=start[0:2]
// 	// endSlice:=end[0:2]
// 	// if startSlice==endSlice{
// 	//     return true
// 	// }else{
// 	//     s,_:=strconv.Atoi(startSlice)
// 	//     e,_:=strconv.Atoi(endSlice)
// 	//     return e-s==1&&(end[3:]=="00")
// 	// }
// 	//懒得费劲了,直接转成时间好了
// 	startHour, _ := strconv.Atoi(start[0:2])
// 	endHour, _ := strconv.Atoi(end[0:2])
// 	startMin, _ := strconv.Atoi(start[3:])
// 	endMin, _ := strconv.Atoi(end[3:])
// 	return (endHour-startHour)*60+endMin-startMin <= 60
// }

//startTime按照人名的第一个时间,获取endTime

//看到题解里有一个比较巧妙的判断是否存在三次的逻辑,就是以长度为3进行移动,判断前后两个时间间隔是否在一小时以内,最终的实现会非常简洁
//整理好的实现如下

func alertNames(keyName []string, keyTime []string) []string {
	var result = []string{}
	if len(keyName) == 0 {
		return result
	}
	var nameMap = make(map[string][]string) //key为人名,value 为时间序列
	for i := 0; i < len(keyName); i++ {
		_, ok := nameMap[keyName[i]]
		if !ok {
			nameMap[keyName[i]] = []string{keyTime[i]}
		} else {
			nameMap[keyName[i]] = append(nameMap[keyName[i]], keyTime[i])
		}
	}
	for key, value := range nameMap {
		if len(value) < 3 {
			continue
		}
		sort.Strings(value)
		isAdded := false
		for i := 2; i < len(value); i++ {
			if isSameSpan(value[i-2], value[i]) && !isAdded {
				result = append(result, key)
				isAdded = true
			}
		}
	}
	sort.Strings(result)
	return result
}
func isSameSpan(start, end string) bool {
	//懒得费劲了,直接转成时间好了
	startHour, _ := strconv.Atoi(start[0:2])
	endHour, _ := strconv.Atoi(end[0:2])
	startMin, _ := strconv.Atoi(start[3:])
	endMin, _ := strconv.Atoi(end[3:])
	return (endHour-startHour)*60+endMin-startMin <= 60
}
