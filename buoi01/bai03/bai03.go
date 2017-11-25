// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 
// http://codeforces.com/problemset/problem/673/A
//

package bai03

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/hoahm/bigo/cmdutil"
)

func check(a []int, n int) int {
	if a[0] > 15 {
		return 15
	}

	time := 0
	for i := 0; i < n; i++ {
		if time + 15 < a[i] {
			return time + 15
		}
		
		time = a[i]
	}

	if time + 15 > 90 {
		time = 90
	} else {
		time += 15
	}

	return time
}

// Execute run 
func Execute() {
	var n int
	fmt.Scanf("%d\n", &n)

	var a = make([]int, n)
	s, _ := cmdutil.ReadLine()
	for i, str := range strings.Split(s, " ") {
		num, _ := strconv.Atoi(str)
		a[i] = num
	}

	fmt.Println(check(a, n))
}
