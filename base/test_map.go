package main

import "fmt"

/*语言集合map*/

func TestMap() {
	// countryCapitalMap := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo", "India": "New delhi"}
	var countryCapitalMap map[string]string
	countryCapitalMap = make(map[string]string)

	capitalNameArray := []string{"巴黎", "罗马", "东京", "新德里"}
	countryNameArray := []string{"France", "Italy", "Japan", "India"}
	for index, country := range countryNameArray {
		countryCapitalMap[country] = capitalNameArray[index]
	}
	for country := range countryCapitalMap {
		fmt.Println(country, "首都是", countryCapitalMap[country])
	}
	/*查看元素在集合中是否存在 */
	capital, ok := countryCapitalMap [ "American" ] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if (ok) {
		fmt.Println("American 的首都是", capital)
	} else {
		fmt.Println("American 的首都不存在")
	}


}

func main() {
	TestMap()
}
