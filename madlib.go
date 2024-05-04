package main

import (
	"fmt"
	"math/rand"
)

var names = []string{"John", "Joy", "Rebecca", "Rob", "Smith", "Roger", "Hayley"}
var occupations = []string{"doctor", "teacher", "engineer", "lawyer", "scientist", "chef", "artist"}
var computingDevices = []string{"laptop", "smartphone", "tablet", "desktop", "smartwatch", "smartglasses", "smartshoes"}
var bodyParts = []string{"head", "shoulder", "knee", "toe", "eye", "ear", "nose"}
var moodAdjectives = []string{"happy", "sad", "angry", "excited", "bored", "tired", "hungry"}
var deviceActions = []string{"play", "pause", "stop", "rewind", "fast forward", "record", "delete"}

func GetRandomMadlib() string {
	var randInt = rand.Intn(7)
	return fmt.Sprintf("%s is a a %d year old %s who has been struggling with a lot of job-related stress. He/she decided to try a new application to relieve stress, which runs on a/an %s to help improve his/her mood.\n\nThe application senses his/her mood through a device he/she wears on his/her %s.\n\nWhen the device senses that he/she is %s, it respond by %s.", names[randInt], randInt+18, occupations[randInt], computingDevices[randInt], bodyParts[randInt], moodAdjectives[randInt], deviceActions[randInt])
}
