package exercise

type GetAllResponse struct {
	Exercises []Exercise `json:"exercises"`
}

type GetByMuscleGroupResponse struct {
	Exercises []Exercise `json:"exercises"`
}
