package main

func sort2(l []int) {
	for a := 0; a < len(l)-1; a++ {
		for b := a + 1; b < len(l); b++ {
			if l[a] > l[b] {
				l[a], l[b] = l[b], l[a]
			}
		}
	}
}

func sort1(l []int) {
	changed := true

	for changed {
		changed = false

		for i := 0; i < len(l)-1; i++ {
			if l[i] > l[i+1] {
				changed = true
				l[i], l[i+1] = l[i+1], l[i]
			}
		}
	}
}
