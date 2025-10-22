package main

import "sort"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	first, last := 1, n
	for first < last {
		mid := first + (last-first)/2
		if isBadVersion(mid) {
			last = mid
		} else {
			first = mid + 1
		}
	}
	return first
}

func main() {

}
