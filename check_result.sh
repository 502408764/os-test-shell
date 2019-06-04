#!/bin/bash

result_file_name="result.txt"

if [ ! -f "$result_file_name" ]; then
    echo "missing ${result_file_name}, please put ${result_file_name} in this working directory"
    exit
fi
#please pass your executable file as the first parameter and the max index of file count as the second parameter
if [ $# -eq 2 ]; then
    echo "parameter count mismatch, use this by $0 <executable file name> <test counts>"
    exit
fi

suffix=".txt"
executable=$1
test_count=$2

for (( i = 0; i < $test_count; ++i )); do
    your_result=`${executable} ${i}${suffix}`
    read line < ${result_file_name}
    your_result_array=(${your_result})
    actual_result_array=(${line})
    your_count=${#your_result_array[@]}
    actual_count=${#actual_result_array[@]}
    if [${your_count} -eq ${actual_count}]; then
        for (( j = 0; j < ${#your_result_array}; ++j )); do
            if [${your_result_array[j]}  -ne ${actual_result_array[j]}]; then
                echo "test ${i}${suffix} wrong answer"
                break
            fi
        done
    else
        echo "test ${i}${suffix} wrong answer"
    fi
done


