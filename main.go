// MIT License
// 
// Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>
// 
// Class: BigO Coding Blue 06
//
// Buoi	01
//	Materials:
//		- https://www.dropbox.com/sh/vtn2l310ivn4j7g/AADrtlIEdTwR4Xr9I081q931a?dl=0
//	BTTL:
//		- http://codeforces.com/problemset/problem/691/A
//		- http://codeforces.com/problemset/problem/677/A
//		- http://codeforces.com/problemset/problem/673/A
//

package main

import (
	"fmt"
	buoi01bai01 "github.com/hoahm/bigo/buoi01/bai01"
	buoi01bai02 "github.com/hoahm/bigo/buoi01/bai02"
	buoi01bai03 "github.com/hoahm/bigo/buoi01/bai03"
)

func welcome() {
	fmt.Println(`
##############################################################
# MIT License                                                #
#                                                            #
# Copyright (c) 2017 Nobi Hoang <nobi.hoa@gmail.com>         #
#                                                            #
# Class: BigO Coding Blue 06                                 #
##############################################################	

		`)
}

func printProjectMenu() {
	fmt.Println(`
==============================
= PROJECT LIST:		
==============================		
Warm up:
	1. Buoi 01 - Bai 01 (Fashion in Berland)
	2. Buoi 02 - Bai 02 (Vanya and Fence)
	3. Buoi 03 - Bai 03 (Bear and Game)
		`)
}

func executeProject(projectID int) {
	fmt.Println("----------------------------")
	switch projectID {
	case 1:
		fmt.Printf("\n\nBuoi 01 - Bai 01\n")
		buoi01bai01.Execute()
		break
	case 2:
		fmt.Printf("\n\nBuoi 01 - Bai 02\n")
		buoi01bai02.Execute()
		break
	case 3:
		fmt.Printf("\n\nBuoi 01 - Bai 03\n")
		buoi01bai03.Execute()
		break
	default:
		fmt.Println("You have chosen the wrong project ID. Please try again!")
		break
	}
}

func main() {
	var cmd string
	var projectID int

	welcome()

	for cmd != "exit()" {
		printProjectMenu()

		fmt.Printf("Choose project ID: ")
		fmt.Scanf("%d", &projectID)

		executeProject(projectID)

		fmt.Println("Press any key to continue, or type `exit()` to quit.")
		fmt.Scanf("%s", &cmd)
	}
}