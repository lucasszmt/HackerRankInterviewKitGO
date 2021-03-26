package warmupChallenges

//JumpingOnClouds is a problem where it has to find the minimum viable solution
//returns the minimum jumps required in the array
//[]int32{0, 0, 0, 1, 0, 0}: Example Input
func JumpingOnClouds(c []int32) int32 {
	var numberOfJumps int32 = 0
	for i := 0; i <= len(c)-2; {
		if i != len(c)-2 {
			if c[i+2] == 0 {
				i = i + 2
				numberOfJumps++
				continue
			}
		}
		i += 1
		numberOfJumps++
	}
	return numberOfJumps
}
