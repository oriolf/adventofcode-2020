for i in {1..25}; do
    cd $i
    go build -o main
    echo "problem $i"
    time cat input2 | ./main 1
    time cat input2 | ./main 2
    rm main
    cd ..
done
