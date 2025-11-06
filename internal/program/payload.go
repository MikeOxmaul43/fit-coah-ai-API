package program

type GetByIdRequest struct {
	Program Program `json:"program"`
}

type CreateProgramRequest struct {
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description"`
	Level         string `json:"level" validate:"required"`
	DurationWeeks int    `json:"duration_weeks"`
}
