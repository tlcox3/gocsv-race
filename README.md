# gocsv-race
This project is meant as a demo of a possible race condtion in the [gocsv](https://github.com/gocarina/gocsv) library. The effects of the race condtion can be seen when parsing multiple CSVs with different delimiters at the same time. It occurs because the SetCSVReader function sets a variable function globally within the package.

## Executing demo
Clone the repo:
```
git clone https://github.com/tlcox3/gocsv-race.git
cd gocsv-race
```

Execute the program:
```
go run .
```

The ouput of the program will most likely resemble:
```
CSV: COLUMN1|COLUMN2
test3|test4
Data: [{Column1: Column2:}]

CSV: COLUMN1,COLUMN2
test1,test2
Data: [{Column1:test1 Column2:test2}]

Data1 Col1: test1 Data2 Col2:
```

It may differ in that the first csv may be the one that has no data. The thing to notice is that either the first or the second CSV's data will not parse. It seems to be the second one that is usually failing due to the Go scheduler allowing it to set it's CSVReader first, then the first call set's it's CSVReader with a different delimiter.

The correct results would be:
```
CSV: COLUMN1,COLUMN2
test1,test2
Data: [{Column1:test1 Column2:test2}]

CSV: COLUMN1|COLUMN2
test3|test4
Data: [{Column1:test3 Column2:test4}]

Data1 Col1: test1 Data2 Col2: test4
```

In this example both CSVs are parsed correctly and that can be seen reflected in the printed data.

## Executing tests with race detection
Run the tests with race detection:
```
go test . --race
```
The test should fail due to both bad data returned and the race detection.

## "Fixing" the results
The results can be fixed by adding a pause to ensure the CSVs are not parsed at the same time. This can be done by uncommenting the sleep line on line 35 of main.go

If this is done the results of the function will be correct, but if you run the tests again they will still trigger race detection.