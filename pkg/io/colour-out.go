package io

import (
	"fmt"
	"github.com/fatih/color"
	"log"
)

type PrintStyle int

const (
	Red PrintStyle = iota
	Green
	Blue
)

func GetStyle(code PrintStyle) *color.Color {
	switch code {
	case Red:
		return color.New(color.BgHiRed, color.Bold)
	case Green:
		return color.New(color.BgHiGreen, color.Bold)
	case Blue:
		return color.New(color.BgHiBlue, color.Bold)
	default:
		return color.New(color.BgBlack, color.Bold)
	}
}

func Panic(err ...interface{}) {
	c := GetStyle(Red)
	c.Print("PANIC")
	print(" ")
	log.Panic(err)
}
func Fatal(err ...interface{}) {
	c := GetStyle(Red)
	c.Print("FATAL")
	print(" ")

	log.Fatal(err)

}

func Println(message ...interface{}) {
	fmt.Println(message...)
}

func Debug(message ...interface{}) {
	// fmt.Println(message...)
}

func Printf(message string, args ...interface{}) {
	fmt.Printf(message, args)
}

func Print(message ...interface{}) {
	fmt.Print(message...)
}

func PrintColor(message string, st PrintStyle) {

	c := GetStyle(st)
	c.Print(" " + message + " ")
	fmt.Print("\t")
}
