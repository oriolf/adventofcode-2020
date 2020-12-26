cd $1
cat input1 | go run main.go 1
cat input2 | go run main.go 1
if [ "$1" = "14" ]; then
    cat input3 | go run main.go 2 # input1 runs for too long
else
    cat input1 | go run main.go 2
fi
cat input2 | go run main.go 2
