package lc

// Time: O(n^2)

func numJewelsInStones(J string, S string) int {
	var count int
	for i := 0; i < len(J); i++ {
		for j := 0; j < len(S); j++ {
			if J[i] == S[j] {
				count++
			}
		}
	}
	return count
}
