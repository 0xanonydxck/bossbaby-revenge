package module

/*
Questions:
	1. Boss Baby(sigma boi), known for his cool and clever ways, deals with teasing from the neighborhood kids who shoot water guns at his house.
	2. In response, Boss Baby seeks revenge by "shooting at least one shot back".
	3. But only if the kids have already shoot at him first.
	4. Check if Boss Baby has sought revenge for:
		- every shot aimed at him at least once.
		- has not initiated any shooting himself.

Input:
	- A string (S, 1 <= len(S) <= 1_000_000) containing characters "S" - (Shot) and "R" - (Revenge).

Output:
	1. "Good Boy"
		- All shot have been revenged at least once.
		- Boss Baby has not initiated any shots.
	2. "Bad Boy"
		- If this conditions are not satisfied.
*/

// --------------------------------------------

const (
	// Fixed Input
	// -----------------

	SHOT    rune = 'S'
	REVENGE rune = 'R'

	// Fixed Output
	// -----------------

	GOOD_BOY string = "Good boy"
	BAD_BOY  string = "Bad boy"
)

/** The BossBabyRevenge function is a function that takes a string as input and returns a string as output. At the core of the function is a stack data structure that is used to keep track of the shots from neighborhood kids and the stack will pop when Boss Baby revenge them.
 * @name: BossBabyRevenge
 * @params: s (string)
 * @returns: r (string)
 * @time-complexity: O(n)
 * @memory-complexity: O(n)
 */
func BossBabyRevenge(s string) (r string) {
	scenarios := []rune(s)
	if !ValidateInput(scenarios) {
		panic("🚨 invalid input!")
	}

	if scenarios[0] == REVENGE {
		r = BAD_BOY
		return
	}

	r = GOOD_BOY
	stack := NewStack[rune]()
	for _, scenario := range scenarios {
		if scenario == SHOT {
			stack.Push(scenario)
		} else {
			stack.Pop()
		}
	}

	if stack.Peek() == SHOT {
		r = BAD_BOY
	}

	return
}

// ------------------------------------------------

func ValidateInput(chars []rune) bool {
	size := len(chars)
	if size <= 1 && size >= 1_000_000 {
		return false
	}

	for _, c := range chars {
		if c != SHOT && c != REVENGE {
			return false
		}
	}
	return true
}
