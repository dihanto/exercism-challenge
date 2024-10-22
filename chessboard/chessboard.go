package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	if _, exists := cb[file]; !exists {
		return 0
	}

	occupiedCount := 0
	for _, occupied := range cb[file] {
		if occupied {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	occupiedCount := 0

	if rank < 1 || rank > 8 {
		return 0
	}
	for _, board := range cb {
		if board[rank-1] {
			occupiedCount++
		}
	}

	return occupiedCount
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0

	for _, board := range cb {
		count += len(board)
	}

	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	occupied := 0

	for _, board := range cb {
		for _, square := range board {
			if square {
				occupied++
			}
		}
	}

	return occupied
}
