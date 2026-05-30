/*
Design a search query autocomplete system for a search engine.

The users will input a sentence ( which may have multiple words and ends with special character '#').

For each character they type except '#', you need to return the top 3 previously entered and most frequently queried sentences that have prefix the same as the part of sentence already typed.

Here are the specific rules:

The frequency for a sentence is defined as the number of times a user typed the exactly same sentence before.
The returned top 3 sentences should be sorted by frequency (The first is the most frequent).  If several sentences have the same frequency, you need to use ASCII-code order (smaller one appears first).
If less than 3 valid sentences exist, then just return as many as you can.
When the input is a special character, it means the sentence ends, and in this case, you need to return an empty list.
 

Your job is to implement the methods of the AutoCompleteSystem:

AutoCompleteSystem(String[] sentences, int[] times): This is the constructor. The input is previously used data. Sentences is a string array consists of previously typed sentences. Times is the corresponding times a sentence has been typed. Your system should record these historical sentences.
Now, the user wants to input a new sentence. The following function will provide the next character the user types:

String[] input(char c): The input c is the next character typed by the user. The character will only be lower-case letters ('a' to 'z'), blank space (' ') or a special character ('#'). Also, the previously typed sentence should be recorded in your system. The output an array will be the top 3 historical sentences that have prefix the same as the part of sentence already typed.
 

Example:

Operation:
AutoCompleteSystem(["i love you", "island",
"ironman", "i love geeksforgeeks"], [5,3,2,2])

The system have already tracked down the 
following sentences and their corresponding 
times: 
"i love you" : 5 times 
"island" : 3 times 
"ironman" : 2 times 
"i love geeksforgeeks" : 2 times 

Now, the user begins another search: 

Operation: input('i') 
Output: 
["i love you", "island","i love 
                       geeksforgeeks"] 

Explanation: 
There are four sentences that have prefix 
"i". Among them, "ironman" and "i love 
geeksforgeeks" have same frequency. Since 
' ' has ASCII code 32 and 'r' has ASCII code
 114, "i love geeksforgeeks" should be in 
front of "ironman". Also we only need to 
output top 3 most frequent sentences, so 
"ironman" will be ignored. 

Operation: input(' ') 
Output: ["i love you","i love geeksforgeeks"] 
Explanation: 
There are only two sentences that have prefix 
"i ". 

Operation: input('a') 
Output: [] 
Explanation: 
There are no sentences that have prefix "i a"

Operation: input('#') 
Output: [] 
Explanation: 
The user finished the input, the sentence 
"i a" should be saved as a historical 
sentence in system. And the next input 
will be counted as a new search.
Your Task:

You don't need to take inputs or give outputs . You just have to complete the input() method and the constructor.

Expected Time Complexity: O(n*max|L|), per input query where n represents the number of historical sentences in the system and L is the maximum length of the words.

Constraints:

The input sentence will always start with a letter and end with '#', and at most one blank space will exist between two words.
The number of complete sentences that to be searched won't exceed 102.
The length of each sentence including those in the historical data and query data won't exceed 102.

*/


/*
Step 1: Insert "car" with frequency 2
Start from sys.root.

Insert characters one by one ('c', 'a', 'r').

At each step, update the freq map.

Trie after inserting "car"

Root
 ├── 'c' (freq={"car": 2})
 │   ├── 'a' (freq={"car": 2})
 │   │   ├── 'r' (freq={"car": 2}) ← End of word "car"
👉 Explanation:

"car"'s frequency 2 is stored at every node of the path.

So if the user types "c", "ca", or "car", we can quickly find "car".

Step 2: Insert "cart" with frequency 3
We start from sys.root.

'c' and 'a' already exist, so we reuse those nodes.

Add 't' as a new node.

Trie after inserting "cart"

Root
 ├── 'c' (freq={"car": 2, "cart": 3})
 │   ├── 'a' (freq={"car": 2, "cart": 3})
 │   │   ├── 'r' (freq={"car": 2, "cart": 3})
 │   │   │   ├── 't' (freq={"cart": 3}) ← End of word "cart"
👉 Explanation:

"car" was already present, so we just updated its frequency map.

"cart" is stored in the path "c -> a -> r -> t".
*/

/*
Example Walkthrough
Let's insert the sentence "i love you" (5 times).

Step 1: Insert "i love you"
We move character by character in the Trie:

Root
 ├── 'i' (freq={"i love you": 5})
 │   ├── ' ' (freq={"i love you": 5})
 │   │   ├── 'l' (freq={"i love you": 5})
 │   │   │   ├── 'o' (freq={"i love you": 5})
 │   │   │   │   ├── 'v' (freq={"i love you": 5})
 │   │   │   │   │   ├── 'e' (freq={"i love you": 5})
 │   │   │   │   │   │   ├── ' ' (freq={"i love you": 5})
 │   │   │   │   │   │   │   ├── 'y' (freq={"i love you": 5})
 │   │   │   │   │   │   │   │   ├── 'o' (freq={"i love you": 5})
 │   │   │   │   │   │   │   │   │   ├── 'u' (freq={"i love you": 5})
Now, if the user types:

"i" → We look at freq in 'i''s TrieNode → Top sentences: ["i love you", "island", "ironman"].

"i l" → We look at freq in 'l''s TrieNode → Top sentences: ["i love you", "i love golang"].
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type TrieNode struct {
	children map[rune]*TrieNode // which are pointers to other TrieNodes
	freq     map[string]int     // keeps track of sentences and their frequency.
}

type AutoCompleteSystem struct {
	root    *TrieNode // The root node of the Trie (this is where the Trie starts).
	current *TrieNode // Tracks the current node as the user types.
	prefix  string    // Stores the current user input (prefix being typed).
}

// constructor, initializes the system and insert the given list of sentences and their frequency.
func NewAutoCompleteSystem(sentances []string, sentancesFreq []int) *AutoCompleteSystem {
	sys := &AutoCompleteSystem{
		root: &TrieNode{
			children: make(map[rune]*TrieNode),
			freq:     make(map[string]int),
		},
	}

	for i, sentance := range sentances {
		sys.insert(sentance, sentancesFreq[i])
	}

	// set the current node
	sys.current = sys.root
	return sys
}

// Insert a sentence into the trie with the given frequestion
func (sys *AutoCompleteSystem) insert(sentance string, freq int) {
	node := sys.root
	for _, char := range sentance {
		if node.children[char] == nil {
			node.children[char] = &TrieNode{
				make(map[rune]*TrieNode),
				make(map[string]int),
			}
		}
		node = node.children[char]  // point to this new node
		node.freq[sentance] += freq // set the frequency
	}
}

// process user input character by character
func (sys *AutoCompleteSystem) input(char byte) []string {
	if char == '\n' {
		return []string{}
	}
	if char == '#' { // end of the input
		// store this sentence into the trie
		sys.insert(sys.prefix, 1)
		fmt.Printf("\nSaved sentence: \"%s\"\n", sys.prefix)
		sys.prefix = "" // reset
		sys.current = sys.root
		return []string{}
	}

	sys.prefix += string(char) // append the current
	if sys.current.children[rune(char)] != nil {
		sys.current = sys.current.children[rune(char)]
	} else {
		return []string{}
	}

	return sys.getTop3Suggestions(sys.current.freq)
}

func (sys *AutoCompleteSystem) getTop3Suggestions(freqMap map[string]int) []string {
	sentencesList := make([]string, 0, len(freqMap))

	for key := range freqMap {
		sentencesList = append(sentencesList, key)
	}
	// sort the list based on santences frequency and then lexicographical order
	sort.Slice(sentencesList, func(i, j int) bool {
		if freqMap[sentencesList[i]] == freqMap[sentencesList[j]] {
			return sentencesList[i] < sentencesList[j]
		}
		return freqMap[sentencesList[i]] > freqMap[sentencesList[j]]
	})

	if len(sentencesList) > 3 {
		return sentencesList[:3]
	}
	return sentencesList
}

func main() {
	system := NewAutoCompleteSystem(
		[]string{"i love you", "island", "ironman", "i love golang"},
		[]int{5, 3, 2, 2},
	)

	fmt.Println("Start typing below. Press '#' to save and reset.")
	fmt.Println("-------------------------------------------------")

	reader := bufio.NewReader(os.Stdin)
	// continous reading
	for {
		// ch, err := reader.ReadByte()
		// if err != nil {
		// 	fmt.Println(fmt.Sprintf("error reading %v", err))
		// 	break
		// }
		ch, _, err := reader.ReadRune()
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}
		fmt.Println()

		suggestions := system.input(byte(ch))

		if ch != '#' && len(suggestions) > 0 {
			fmt.Printf("Suggestions: %v\n", suggestions)
		}
	}
}
