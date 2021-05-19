package main

import "testing"

func TestFind2DimensionArray(t *testing.T){

	data := [][] int {
		{1,2,8,9},
		{2,4,9,12},
		{4,7,10,13},
		{6,8,11,15},
	}

	target := 2
	res := Find2DimensionArray(target, data)


}
