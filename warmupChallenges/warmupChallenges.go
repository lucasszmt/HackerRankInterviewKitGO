package warmupChallenges

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
