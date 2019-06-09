#!/bin/bash

result_file_name="result.txt"

if [[ ! -f "$result_file_name" ]]; then
    echo "missing ${result_file_name}, please put ${result_file_name} in this working directory"
    exit
fi
#please pass your executable file as the first parameter and the max index of file count as the second parameter
if [[ ! $# -eq 1 ]]; then
    echo "parameter count mismatch, use this by $0 <executable file name>"
    exit
elif [[ ! -f "$1" ]]; then
    echo "no executable file $1 exist in working directory"
    exit
fi

suffix=".txt"
executable=$1

i=0
while read line || [[ -n ${line} ]]
do
    your_result=`./${executable} ${i}${suffix}`
    your_result_array=(${your_result})
    actual_result_array=(${line})
    your_count=${#your_result_array[@]}
    actual_count=${#actual_result_array[@]}
    is_correct=1
    if [[ ${your_count} -eq ${actual_count} ]]; then
        for (( j = 0; j < ${#your_result_array}; ++j )); do
            if [[ ${your_result_array[j]} != ${actual_result_array[j]} ]]; then
                is_correct=0
                break
            fi
        done
        if [[ ${is_correct} -eq 1 ]]; then
            echo "test ${i}${suffix} correct"
        else
            echo "test ${i}${suffix} wrong answer"
        fi
    else
        echo "test ${i}${suffix} wrong answer"
    fi
    i=$(( $i + 1 ))
done < ${result_file_name}


