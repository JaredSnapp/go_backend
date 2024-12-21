package models

type Workout struct {
	// excercises
}

type WeightedExcercise struct {
	Name string         `json:"name"`
	Reps []WeightedReps `json:"weight_reps"`
}

type BodyExcercise struct {
	Name string `json:"name"`
	Reps []int  `json:"reps"`
}

type WeightedReps struct {
	Weight int `json:"weight"`
	Reps   int `json:"reps"`
}

type TimeExcercise struct {
	Name    string `json:"name"`
	Minutes string `json:"minutes"`
}
