package main

// ReviveResult base structure of revive json result
type ReviveResult struct {
	Failure  string
	RuleName string
	Position RevivePosition
}

// RevivePosition revive json position
type RevivePosition struct {
	Start RevivePositionResult
	End   RevivePositionResult
}

// RevivePositionResult revive json position result
type RevivePositionResult struct {
	Filename string
	Line     int
	Offset   int
	Column   int
}
