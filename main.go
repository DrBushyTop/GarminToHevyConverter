package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

// Exercise represents a single guess by Garmin for the exercise performed
type Exercise struct {
	Category    string  `json:"category"`
	Name        string  `json:"name"`
	Probability float64 `json:"probability"`
}

// GarminData represents the relevant fields of the Garmin workout data
type GarminData struct {
	ActivityId   int `json:"activityId"`
	ActivityType struct {
		TypeKey string `json:"typeKey"`
	} `json:"activityType"`
	StartTimeLocal   string  `json:"startTimeLocal"`
	ActivityName     string  `json:"activityName"`
	Description      string  `json:"description"`
	Duration         float64 `json:"duration"`
	FullExerciseSets []struct {
		Category        string     `json:"category"`
		SetType         string     `json:"setType"`
		Exercises       []Exercise `json:"exercises"`
		RepetitionCount int        `json:"repetitionCount"`
		Duration        float64    `json:"duration"`
		Weight          int        `json:"weight"`
	} `json:"fullExerciseSets"`
}

func main() {
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Usage: ./GarminToHevy <path-to-json-file>")
		return
	}

	// Load the custom exercise names
	LoadCustomExerciseNames()

	jsonFile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(io.Reader(jsonFile))

	var data []GarminData
	json.Unmarshal(byteValue, &data)

	// Process the data
	processedData := processGarminData(data)

	// Write the processed data to a CSV file
	fileName := "Converted_Garmin_Workouts"
	if len(os.Args) == 3 {
		fileName = os.Args[2]
	}
	writeCSV(processedData, fileName)
}

func processGarminData(data []GarminData) [][]string {
	// Process the data according to the rules
	// This function will be lengthy due to the data processing logic
	var processedData [][]string
	workoutsWithUnknownExercises := make(map[int]string)
	workoutsNotFoundInMapping := make(map[string]bool)

	for _, activity := range data {
		if activity.ActivityType.TypeKey != "strength_training" {
			continue
		}

		date := parseDate(activity.StartTimeLocal)
		workoutName := fmt.Sprintf("%s-%s", activity.ActivityName, strings.Split(activity.StartTimeLocal, " ")[0])
		workoutNotes := fmt.Sprintf("\"%s\"", activity.Description)
		workoutDuration := formatWorkoutDuration(activity.Duration)

		var lastExerciseName string
		setOrder := 0

		for _, set := range activity.FullExerciseSets {
			// Skip warmup sets, sets with no weight, and sets with no repetitions (avoids custom exercises with no ability to track weight)
			if set.SetType != "ACTIVE" || set.Category == "WARM_UP" || set.RepetitionCount == 0 || (set.Weight == 0 && set.RepetitionCount == 1) {
				continue
			}

			exerciseName := getExerciseNameWithHighestProbability(set.Exercises)
			if exerciseName == "UNKNOWN" {
				workoutsWithUnknownExercises[activity.ActivityId] = fmt.Sprintf("https://connect.garmin.com/modern/activity/%d", activity.ActivityId)
			}

			if exerciseName != lastExerciseName {
				setOrder = 1
				lastExerciseName = exerciseName
			} else {
				setOrder++
			}

			convertedExcerciseName := convertExerciseNameToHevyFormat(exerciseName)
			if (strings.Contains(convertedExcerciseName, "_") || convertedExcerciseName == strings.ToUpper(convertedExcerciseName)) && convertedExcerciseName != "UNKNOWN" {
				workoutsNotFoundInMapping[convertedExcerciseName] = true
			}

			// Weight should be gotten from the data, but hevy has some rep-only exercises, so we might need to set it to 0 to avoid creating a custom exercise
			// It's arguable whether this should be generally done, but I'm doing it for my own data
			weight := formatWeight(set.Weight)
			if _, ok := repOnlyExercises[convertedExcerciseName]; ok {
				weight = "0"
			}

			reps := formatReps(set.RepetitionCount)
			// Seconds with real values do not seem to import correctly to Hevy
			//seconds := strconv.Itoa(int(set.Duration))
			seconds := "0"

			processedRow := []string{
				date, workoutName, convertedExcerciseName, strconv.Itoa(setOrder), weight, "kg", reps, "", "", "km", seconds, "", workoutNotes, workoutDuration,
			}
			processedData = append(processedData, processedRow)
		}
	}

	// Print out the workouts with exercises that were not found in the mapping
	if len(workoutsNotFoundInMapping) > 0 {
		fmt.Println("Workouts not mapped to Hevy equivalents. These need a PR to the repo to add the mapping, or a custom exercise in Hevy:")
		for key := range workoutsNotFoundInMapping {
			fmt.Printf("%s\n", key)
		}
	}

	// Print out the workouts with unknown exercises
	if len(workoutsWithUnknownExercises) > 0 {
		fmt.Println("Workouts with unknown exercises. You might want to go through these, add the correct exercise and then redownload the Garmin data:")
		for _, value := range workoutsWithUnknownExercises {
			fmt.Printf("%s\n", value)
		}
	}

	return processedData
}

func writeCSV(data [][]string, fileName string) {
	file, err := os.Create(fmt.Sprintf("%s.csv", fileName))
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the headers
	writer.Write(
		[]string{
			"Date",
			"Workout Name",
			"Exercise Name",
			"Set Order",
			"Weight",
			"Weight Unit",
			"Reps",
			"RPE",
			"Distance",
			"Distance Unit",
			"Seconds",
			"Notes",
			"Workout Notes",
			"Workout Duration",
		},
	)

	// Write the data
	for _, value := range data {
		writer.Write(value)
	}
}
func parseDate(dateStr string) string {
	parsedDate, err := time.Parse("2006-01-02T15:04:05.0", dateStr)
	if err != nil {
		parsedDate, err = time.Parse("2006-01-02 15:04:05", dateStr)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return ""
		}
	}
	return parsedDate.Format("2006-01-02 15:04:05")
}

// formatWorkoutDuration converts duration from seconds to a string format like "30m" or "45s".
func formatWorkoutDuration(duration float64) string {
	roundedDuration := int(duration) // Round to nearest whole number
	if roundedDuration >= 60 {
		return fmt.Sprintf("%dm", roundedDuration/60)
	}
	return fmt.Sprintf("%ds", roundedDuration)
}

// getExerciseNameWithHighestProbability selects the exercise name with the highest probability.
func getExerciseNameWithHighestProbability(exercises []Exercise) string {
	if len(exercises) == 0 {
		return ""
	}

	maxProb := 0.0
	var selectedExercise Exercise
	for _, exercise := range exercises {
		if exercise.Probability > maxProb {
			maxProb = exercise.Probability
			selectedExercise = exercise
		}
	}

	if selectedExercise.Name == "" {
		return selectedExercise.Category
	}
	return selectedExercise.Name
}

// formatWeight converts weight from Garmin format to kg and returns it as a string.
func formatWeight(weight int) string {
	if weight > 0 {
		return strconv.Itoa(weight / 1000)
	}
	return "0"
}

// formatReps ensures that the repetitions are set to '1' if null or missing.
func formatReps(reps int) string {
	if reps > 0 {
		return strconv.Itoa(reps)
	}
	return "1"
}