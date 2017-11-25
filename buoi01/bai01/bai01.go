// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 
// http://codeforces.com/problemset/problem/691/A
//

package bai01

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/hoahm/bigo/cmdutil"
)

// Execute run 
func Execute() {
	var n int
	fmt.Scanf("%d\n", &n)

	var a = make([]int, n)
	var s string
	s, _ = cmdutil.ReadLine()

	for i, str := range strings.Split(s, " ") {
		num, _ := strconv.Atoi(str)
		a[i] = num
	}

	count := 0
	for i := 0; i < n; i++ {
		if a[i] == 1 {
			count++
		}
	}

	if len(a) == 1 {
		if count == 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	} else {
		if count == n - 1 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}