package main

import (
	"log"
	"os"
	"fmt"
)
var InningTop, InningMiddle, InningBottom, InningEnd = 1, 2, 3, 4

func init() {
	err := os.Mkdir("gameinfo", 0750)
	if err != nil && !os.IsExist(err) {
		log.Printf("The game info directory already exists, proceeding!")
	}
}

func OsCheck() { // There has to be a better way to do this. I refuse to believe there isn't.
	if os.Getegid != -1 // -1 is always returned on Windows.
	log.Fatalln("Sorry, but your operating system is (temporarily) not supported. Please use Windows until we establish support for your system.")
}

func TballInningFlow() {
	Top1st := 1,
	Middle1st := 2,
	Bottom1st := 3,
	End1st := 4,
	Top2nd := 5,
	Middle2nd := 6,
	Bottom2nd := 7,
	End2nd := 8,
	Top3rd := 9,
	Middle3rd := 10,
	Bottom3rd := 11,
	End3rd := 12,
	Top4th := 13,
	Middle4th := 14,
	Bottom4th := 15,
	End4th := 16,
}

func CoachPitchInningFlow() {
	Top1st := 1,
	Middle1st := 2,
	Bottom1st := 3,
	End1st := 4,
	Top2nd := 5,
	Middle2nd := 6,
	Bottom2nd := 7,
	End2nd := 8,
	Top3rd := 9,
	Middle3rd := 10,
	Bottom3rd := 11,
	End3rd := 12,
	Top4th := 13,
	Middle4th := 14,
	Bottom4th := 15,
	End4th := 16,
}

func MinorInningFlow() {
	Top1st := 1,
	Middle1st := 2,
	Bottom1st := 3,
	End1st := 4,
	Top2nd := 5,
	Middle2nd := 6,
	Bottom2nd := 7,
	End2nd := 8,
	Top3rd := 9,
	Middle3rd := 10,
	Bottom3rd := 11,
	End3rd := 12,
	Top4th := 13,
	Middle4th := 14,
	Bottom4th := 15,
	End4th := 16,
	Top5th := 17,
	Middle5th := 18,
	Bottom5th := 19,
	End5th := 20,
	Top6th := 21,
	Middle6th := 22,
	Bottom6th := 23,
	End6th := 24,
}

func MajorInningFlow() {
	Top1st := 1,
	Middle1st := 2,
	Bottom1st := 3,
	End1st := 4,
	Top2nd := 5,
	Middle2nd := 6,
	Bottom2nd := 7,
	End2nd := 8,
	Top3rd := 9,
	Middle3rd := 10,
	Bottom3rd := 11,
	End3rd := 12,
	Top4th := 13,
	Middle4th := 14,
	Bottom4th := 15,
	End4th := 16,
	Top5th := 17,
	Middle5th := 18,
	Bottom5th := 19,
	End5th := 20,
	Top6th := 21,
	Middle6th := 22,
	Bottom6th := 23,
	End6th := 24,
}

func Innings(TBallInningFlow, CoachPitchInningFlow, MinorInningFlow, MajorInningFlow) {

}

type DAuth string
