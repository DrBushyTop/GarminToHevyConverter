# Garmin to Hevy data converter

This is a simple tool to convert Garmin Strength workout data in JSON format to CSV files that can be imported into [Hevy](https://www.hevyapp.com/).
The actual format is the CSV export format from the [Strong app](https://www.strong.app/) (at the time of writing, the only importable data type in Hevy).

## Usage

1. Get the JSON data from Garmin. I used this [Chrome extension](https://chromewebstore.google.com/detail/garmin-workout-downloader/hpimimpdkghmejbcldfccdbaebjifnkk).
2. Run the tool:
```sh 
go run . <path-to-json-file>
e.g. 
go run . ./data.json

# Optionally you can specify the output file name as a second argument. 
# By default it will be named "Converted_Garmin_Workouts.csv"
go run . <path-to-json-file> <name-of-csv-file-without-extension>
e.g.
go run . ./data.json myImportFile
```
3. Use Hevy's import feature to import the CSV file.

## Notes

- The tool has been tested once, and probably won't be updated unless I need to use it again.
- Only English language is supported for the Garmin exercise names.
- The tool will skip any workouts that are not strength workouts.
- The tool will skip any warmup or rest steps from the Garmin data. Any data without rep values will be set as 1 rep.
- Exercise names picked from Garmin data will prioritize the name with the highest "probability" value. It seems that any manually set exercises have this set to 100 and will thus be picked.
- Due to Garmin's limited selection of exercises, I had to use some placeholders for my own workouts. Yours might not match. To change these, edit the values in the conversions.go customExercises map.
- The tool tries to avoid creating custom exercises in Hevy, but I ended up with some from my custom setup.