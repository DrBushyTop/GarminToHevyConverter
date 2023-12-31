package main

import "fmt"

// convertExerciseNameToHevyFormat converts the exercise name to the Hevy format.
// It prioritizes the user's custom mappings over the default mappings.
func convertExerciseNameToHevyFormat(exerciseName string) string {
	if val, ok := customExerciseNameMap[exerciseName]; ok {
		return fmt.Sprintf("%s", val)
	}
	if val, ok := exerciseNameMap[exerciseName]; ok {
		return fmt.Sprintf("%s", val)
	}
	return exerciseName
}

// customExerciseNameMap maps the Garmin exercise name to a (possibly custom) exercise in Hevy.
// These are my personal mappings and may not be applicable to you, as I've used the limited Garmin exercise selection
// as placeholders in my own training.
var customExerciseNameMap = map[string]string{
	"BARBELL_HACK_SQUAT":                       "Hack Squat (Machine)",
	"BEHIND_THE_BACK_ONE_ARM_CABLE_CURL":       "Single Arm Behind the Back Bicep Curl (Cable)",
	"CHEST_SUPPORTED_DUMBBELL_ROW":             "Iso-Lateral Row (Machine)",
	"CABLE_CROSSOVER":                          "Butterfly (Pec Deck)",
	"DUMBBELL_FLYE":                            "Chest Fly (Machine)",
	"EZ_BAR_OVERHEAD_TRICEPS_EXTENSION":        "Triceps Extension (EZ Bar)",
	"INCLINE_DUMBBELL_SHRUG":                   "Incline Shrug (Dumbbell)",
	"INCLINE_REVERSE_FLYE":                     "Rear Delt Reverse Fly (Machine)",
	"INCLINE_Y_RAISE":                          "Lateral Raise (Cable)",
	"PULL_UP":                                  "Pull Up (Assisted)",
	"SEATED_EZ_BAR_OVERHEAD_TRICEPS_EXTENSION": "Triceps Extension (EZ Bar)",
	"STANDING_CALF_RAISE":                      "Standing Calf Raise (Barbell)",
	"SWISS_BALL_HIP_RAISE_AND_LEG_CURL":        "Nordic Hamstring Curls",
	"TRICEP_EXTENSION":                         "Triceps Extension (Cable)",
	"UPRIGHT_ROW":                              "Upright Row (Cable)",
	"WALKING_LUNGE":                            "Walking Lunge (Dumbbell)",
	"WEIGHTED_DIP":                             "Chest Dip (Assisted)",
}

// exerciseNameMap maps the Garmin exercise name to the Hevy exercise name.
// Note that this map is not complete.
// exerciseNameMap maps the Garmin exercise name to the Hevy exercise name.
// Note that this map is not complete.
var exerciseNameMap = map[string]string{
	"ALTERNATING_DUMBBELL_ROW":                   "Dumbbell Row",
	"ARNOLD_PRESS":                               "Arnold Press (Dumbbell)",
	"BARBELL_BACK_SQUAT":                         "Squat (Barbell)",
	"BARBELL_BENCH_PRESS":                        "Bench Press (Barbell)",
	"BARBELL_BICEPS_CURL":                        "Bicep Curl (Barbell)",
	"BARBELL_DEADLIFT":                           "Deadlift (Barbell)",
	"BARBELL_FRONT_SQUAT":                        "Front Squat",
	"BARBELL_HACK_SQUAT":                         "Hack Squat",
	"BARBELL_HIP_THRUST_WITH_BENCH":              "Hip Thrust (Barbell)",
	"BARBELL_LUNGE":                              "Lunge (Barbell)",
	"BARBELL_ROW":                                "Bent Over Row (Barbell)",
	"BARBELL_WRIST_CURL":                         "Behind the Back Bicep Wrist Curl (Barbell)",
	"BEHIND_THE_BACK_BARBELL_REVERSE_WRIST_CURL": "Behind the Back Bicep Wrist Curl (Barbell)",
	"BEHIND_THE_BACK_ONE_ARM_CABLE_CURL":         "Bicep Curl (Cable)",
	"BENT_OVER_ROW_WITH_BARBELL":                 "Bent Over Row (Barbell)",
	"BICEPS_STRETCH":                             "Stretching",
	"BICYCLE_CRUNCH":                             "Bicycle Crunch",
	"BODY_WEIGHT_DIP":                            "Chest Dip",
	"CABLE_BICEPS_CURL":                          "Bicep Curl (Cable)",
	"CABLE_CRUNCH":                               "Cable Crunch",
	"CABLE_CROSSOVER":                            "Cable Fly Crossovers",
	"CABLE_OVERHEAD_TRICEPS_EXTENSION":           "Triceps Extension (Cable)",
	"CHEST_SUPPORTED_DUMBBELL_ROW":               "Chest Supported Incline Row (Dumbbell)",
	"CHIN_UP":                                    "Chin Up (Weighted)",
	"CLOSE_GRIP_BARBELL_BENCH_PRESS":             "Bench Press - Close Grip (Barbell)",
	"CRUNCH":                                     "Crunch",
	"DIAMOND_PUSH_UP":                            "Diamond Push Up",
	"DUMBBELL_FLYE":                              "Chest Fly (Dumbbell)",
	"DUMBBELL_HAMMER_CURL":                       "Hammer Curl (Dumbbell)",
	"DUMBBELL_LATERAL_RAISE":                     "Lateral Raise (Dumbbell)",
	"DUMBBELL_ROW":                               "Dumbbell Row",
	"DUMBBELL_SHOULDER_PRESS":                    "Shoulder Press (Dumbbell)",
	"DUMBBELL_SHRUG":                             "Shrug (Dumbbell)",
	"FACE_PULL":                                  "Face Pull",
	"GOBLET_SQUAT":                               "Goblet Squat",
	"HANGING_LEG_RAISE":                          "Hanging Leg Raise",
	"INCLINE_DUMBBELL_BENCH_PRESS":               "Incline Bench Press (Dumbbell)",
	"INCLINE_DUMBBELL_BICEPS_CURL":               "Seated Incline Curl (Dumbbell)",
	"INCLINE_REVERSE_FLYE":                       "Rear Delt Reverse Fly (Dumbbell)",
	"INCLINE_Y_RAISE":                            "Chest Supported Y Raise (Dumbbell)",
	"INDOOR_ROW":                                 "Rowing Machine",
	"KNEELING_LAT_PULLDOWN":                      "Single Arm Lat Pulldown",
	"LAT_PULLDOWN":                               "Lat Pulldown (Cable)",
	"LAT_PULLOVER":                               "Straight Arm Lat Pulldown (Cable)",
	"LEG_CURL":                                   "Lying Leg Curl (Machine)",
	"LEG_EXTENSIONS":                             "Leg Extension (Machine)",
	"LEG_PRESS":                                  "Leg Press (Machine)",
	"LYING_EZ_BAR_TRICEPS_EXTENSION":             "Skullcrusher (Barbell)",
	"MEDICINE_BALL_ALTERNATING_V_UP":             "V Up",
	"MILITARY_PRESS":                             "Overhead Press (Barbell)",
	"MODIFIED_SIT_UP":                            "Sit Up (Weighted)",
	"NECK_TILTS":                                 "Lying Neck Extension (Weighted)",
	"ONE_ARM_CABLE_LATERAL_RAISE":                "Single Arm Lateral Raise (Cable)",
	"ONE_ARM_CONCENTRATION_CURL":                 "Concentration Curl",
	"ONE_ARM_PREACHER_CURL":                      "Preacher Curl (Dumbbell)",
	"OVERHEAD_BARBELL_PRESS":                     "Overhead Press (Barbell)",
	"PULL_APART":                                 "Band Pullaparts",
	"PULL_UP":                                    "Pull Up",
	"REVERSE_GRIP_PULLDOWN":                      "Reverse Grip Lat Pulldown (Cable)",
	"ROMANIAN_DEADLIFT":                          "Romanian Deadlift (Barbell)",
	"SEATED_ALTERNATING_DUMBBELL_BICEPS_CURL":    "Bicep Curl (Dumbbell)",
	"SEATED_CALF_RAISE":                          "Seated Calf Raise",
	"SEATED_CABLE_ROW":                           "Seated Cable Row - V Grip (Cable)",
	"SEATED_DUMBBELL_BICEPS_CURL":                "Bicep Curl (Dumbbell)",
	"SEATED_DUMBBELL_SHOULDER_PRESS":             "Shoulder Press (Dumbbell)",
	"SEATED_EZ_BAR_OVERHEAD_TRICEPS_EXTENSION":   "Triceps Extension (Barbell)",
	"SMITH_MACHINE_BENCH_PRESS":                  "Bench Press (Smith Machine)",
	"STANDING_ALTERNATING_DUMBBELL_CURLS":        "Bicep Curl (Dumbbell)",
	"STANDING_CALF_RAISE":                        "Standing Calf Raise",
	"STANDING_CABLE_PULLOVER":                    "Straight Arm Lat Pulldown (Cable)",
	"STANDING_DUMBBELL_BICEPS_CURL":              "Bicep Curl (Dumbbell)",
	"STANDING_EZ_BAR_BICEPS_CURL":                "EZ Bar Biceps Curl",
	"STANDING_HIP_ABDUCTION":                     "Hip Abduction (Machine)",
	"STATIC_BACK_EXTENSION":                      "Back Extension (Hyperextension)",
	"STRETCH_LAT":                                "Stretching",
	"STRETCH_PECTORAL":                           "Stretching",
	"STRETCH_SHOULDER":                           "Stretching",
	"STRETCH_WALL_CHEST_AND_SHOULDER":            "Stretching",
	"SUMO_DEADLIFT":                              "Sumo Deadlift",
	"TRICEP_EXTENSION":                           "Triceps Extension (Dumbbell)",
	"TRICEPS_EXTENSION_ON_FLOOR":                 "Skullcrusher (Barbell)",
	"TRICEPS_PRESSDOWN":                          "Triceps Pushdown",
	"UPRIGHT_ROW":                                "Upright Row (Barbell)",
	"WALKING_LUNGE":                              "Walking Lunge",
	"WEIGHT_PLATE_FRONT_RAISE":                   "Plate Front Raise",
	"WEIGHTED_CRUNCH":                            "Crunch (Weighted)",
	"WEIGHTED_DIP":                               "Chest Dip (Weighted)",
	"WEIGHTED_HANGING_LEG_RAISE":                 "Hanging Leg Raise",
	"WEIGHTED_LEG_CURL":                          "Lying Leg Curl (Machine)",
	"WEIGHTED_LEG_EXTENSIONS":                    "Leg Extension (Machine)",
	"WEIGHTED_PULL_UP":                           "Pull Up (Weighted)",
	"WEIGHTED_PUSH_UP":                           "Push Up (Weighted)",
	"WEIGHTED_SEATED_CALF_RAISE":                 "Seated Calf Raise",
	"WEIGHTED_STANDING_CALF_RAISE":               "Standing Calf Raise (Barbell)",
	"WEIGHTED_STANDING_HIP_ABDUCTION":            "Hip Abduction (Machine)",
	"WEIGHTED_WALKING_LUNGE":                     "Walking Lunge (Dumbbell)",
	"WIDE_GRIP_SEATED_CABLE_ROW":                 "Seated Cable Row - Bar Wide Grip",
}

// repOnlyExercises is a map of exercises that only have a rep count and no weight in Hevy.
var repOnlyExercises = map[string]bool{
	"Nordic Hamstring Curls": true,
}