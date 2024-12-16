package main

func north(input *[]string, i int, j int) bool {
	if i < 3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i-1][j]) != "M" ||
		string((*input)[i-2][j]) != "A" ||
		string((*input)[i-3][j]) != "S" {
		return false
	}
	return true
}

func east(input *[]string, i int, j int) bool {
	if j >= len((*input)[i])-3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i][j+1]) != "M" ||
		string((*input)[i][j+2]) != "A" ||
		string((*input)[i][j+3]) != "S" {
		return false
	}
	return true
}

func south(input *[]string, i int, j int) bool {
	if i >= len((*input))-3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i+1][j]) != "M" ||
		string((*input)[i+2][j]) != "A" ||
		string((*input)[i+3][j]) != "S" {
		return false
	}
	return true
}

func west(input *[]string, i int, j int) bool {
	if j < 3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i][j-1]) != "M" ||
		string((*input)[i][j-2]) != "A" ||
		string((*input)[i][j-3]) != "S" {
		return false
	}
	return true
}

func northEast(input *[]string, i int, j int) bool {
	if i < 3 ||
		j >= len((*input)[i])-3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i-1][j+1]) != "M" ||
		string((*input)[i-2][j+2]) != "A" ||
		string((*input)[i-3][j+3]) != "S" {
		return false
	}
	return true
}

func southEast(input *[]string, i int, j int) bool {
	if i >= len((*input))-3 ||
		j >= len((*input)[i])-3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i+1][j+1]) != "M" ||
		string((*input)[i+2][j+2]) != "A" ||
		string((*input)[i+3][j+3]) != "S" {
		return false
	}
	return true
}

func southWest(input *[]string, i int, j int) bool {
	if i >= len((*input))-3 ||
		j < 3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i+1][j-1]) != "M" ||
		string((*input)[i+2][j-2]) != "A" ||
		string((*input)[i+3][j-3]) != "S" {
		return false
	}
	return true
}

func northWest(input *[]string, i int, j int) bool {
	if i < 3 ||
		j < 3 {
		return false
	}
	if string((*input)[i][j]) != "X" ||
		string((*input)[i-1][j-1]) != "M" ||
		string((*input)[i-2][j-2]) != "A" ||
		string((*input)[i-3][j-3]) != "S" {
		return false
	}
	return true
}
