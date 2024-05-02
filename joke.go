package main

import (
	"math/rand"
)

type Joke struct {
	ID    int
	Value string
}

var jokes = []Joke{
	{1, "Why did the golfer bring two pairs of pants? In case he got a hole in one!"},
	{2, "What do you call a fish wearing a crown? A king fish!"},
	{3, "Why did the tomato turn red? Because it saw the salad dressing!"},
	{4, "What do you call a bear with no teeth? A gummy bear!"},
	{5, "Why did the scarecrow win an award? Because he was outstanding in his field!"},
	{6, "What do you call a pile of cats? A meowtain!"},
	{7, "Why did the math book look sad? Because it had too many problems!"},
	{8, "What do you call a sleeping bull? A bulldozer!"},
	{9, "Why did the cookie go to the doctor? Because it was feeling crumbly!"},
	{10, "What do you call a belt made out of watches? A waist of time!"},
}

func GetRandomJoke() Joke {
	var randInt = rand.Intn(10) + 1
	for _, joke := range jokes {
		if joke.ID == randInt {
			return joke
		}
	}
	return Joke{ID: 404, Value: "No joke found"}
}
