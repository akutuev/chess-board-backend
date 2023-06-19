package board

import "fmt"

type piece string

func (p piece) ToString() string {
	return string(p)
}

const (
	empty piece = "o"

	PawnB   piece = "bP"
	KnightB piece = "bKn"
	BishopB piece = "bB"
	QueenB  piece = "bQ"
	KingB   piece = "bK"
	RockB   piece = "bR"

	PawnW   piece = "wP"
	KnightW piece = "wKn"
	BishopW piece = "wB"
	QueenW  piece = "wQ"
	KingW   piece = "wK"
	RockW   piece = "wR"
)

type Cell struct {
	activePiece piece
}

func InitBoard() [8][8]Cell {
	boardCells := [8][8]Cell{}

	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			boardCells[file][rank] = Cell{
				empty,
			}
		}
	}

	boardCells[0][0] = Cell{activePiece: RockW}
	boardCells[0][1] = Cell{activePiece: KnightW}
	boardCells[0][2] = Cell{activePiece: BishopW}
	boardCells[0][3] = Cell{activePiece: QueenW}
	boardCells[0][4] = Cell{activePiece: KingW}
	boardCells[0][5] = Cell{activePiece: BishopW}
	boardCells[0][6] = Cell{activePiece: KnightW}
	boardCells[0][7] = Cell{activePiece: RockW}
	boardCells[1][0] = Cell{activePiece: PawnW}
	boardCells[1][1] = Cell{activePiece: PawnW}
	boardCells[1][2] = Cell{activePiece: PawnW}
	boardCells[1][3] = Cell{activePiece: PawnW}
	boardCells[1][4] = Cell{activePiece: PawnW}
	boardCells[1][5] = Cell{activePiece: PawnW}
	boardCells[1][6] = Cell{activePiece: PawnW}
	boardCells[1][7] = Cell{activePiece: PawnW}

	boardCells[7][0] = Cell{activePiece: RockB}
	boardCells[7][1] = Cell{activePiece: KnightB}
	boardCells[7][2] = Cell{activePiece: BishopB}
	boardCells[7][3] = Cell{activePiece: QueenB}
	boardCells[7][4] = Cell{activePiece: KingB}
	boardCells[7][5] = Cell{activePiece: BishopB}
	boardCells[7][6] = Cell{activePiece: KnightB}
	boardCells[7][7] = Cell{activePiece: RockB}
	boardCells[6][0] = Cell{activePiece: PawnB}
	boardCells[6][1] = Cell{activePiece: PawnB}
	boardCells[6][2] = Cell{activePiece: PawnB}
	boardCells[6][3] = Cell{activePiece: PawnB}
	boardCells[6][4] = Cell{activePiece: PawnB}
	boardCells[6][5] = Cell{activePiece: PawnB}
	boardCells[6][6] = Cell{activePiece: PawnB}
	boardCells[6][7] = Cell{activePiece: PawnB}

	return boardCells
}

func PrintBoard(boardCells [8][8]Cell) {
	for file := 0; file < 8; file++ {
		for rank := 0; rank < 8; rank++ {
			fmt.Printf(boardCells[file][rank].activePiece.ToString() + "\t")
		}
		fmt.Printf("\n")
	}
}

var initiaPosition = "a1:wR;b1:wKn;c1:wB;d1:wQ;e1:wK;f1:wB;g1:wKn;h1:wR;a2:wP;b2:wwP;c2:wP;d2:wP;e2:wP;f2:wP;g2:wP;h2:wP;a8:wR;b8:wKn;c8:wB;d8:wQ;e8:wK;f8:wB;g8:wKn;h8:wR;a7:wP;b7:wP;c7:wP;d7:wP;e7:wP;f7:wP;g7:wP;h7:wP;"
