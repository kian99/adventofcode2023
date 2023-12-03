package main

import "regexp"

const top = true
const bottom = false

var reNumberMatch = regexp.MustCompile(`\d+`)
var reGearMatch = regexp.MustCompile(`\*`)

var reSpecialChar = regexp.MustCompile(`[^\d.]`)
var reDigitMatch = regexp.MustCompile(`\d+`)
