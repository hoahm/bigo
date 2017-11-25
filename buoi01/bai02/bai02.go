// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 
// http://codeforces.com/problemset/problem/677/A
//

package bai02

import (
	"fmt"
	"strings"
	"strconv"
	"github.com/hoahm/bigo/cmdutil"
)

// Execute run 
func Execute() {
	var n, h int
	fmt.Scanf("%d %d\n", &n, &h)

	var a = make([]int, n)
	s, _ := cmdutil.ReadLine()
	for i, str := range strings.Split(s, " ") {
		num, _ := strconv.Atoi(str)
		a[i] = num
	}

	sum := 0
	for i := 0; i < n; i++ {
		if a[i] > h  {
			sum += 2
		} else {
			sum++
		}
	}
	fmt.Println(sum)
}
