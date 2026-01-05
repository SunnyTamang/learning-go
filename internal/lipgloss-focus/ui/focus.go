// package ui
//
// type FocusedSection int
//
// const (
// 	//FocusHeader FocusedSection = iota
// 	FocusLeft FocusedSection = iota
// 	FocusBody
// 	FocusFooter
// 	FocusHeader
// 	FocusCount
// )

package ui

type FocusedSection int

const (
	FocusLeft FocusedSection = iota
	FocusBody
	FocusFooter
)

const FocusCount = 3
