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

execute() {
    executable_suffix=`echo ${executable##*.}`
    if [[ ${executable} = "jar" ]]; then
        echo `python ./${executable}`
    elif [[ ${executable} = "py" ]]; then
        echo `java -jar `
    else
        echo `./${executable} $1 | tr -d '\r'`
    fi
}

i=0
while read line || [[ -n ${line} ]]
do
    if [[ ${line} != "" ]]; then
        line=`echo ${line} | tr -d '\r'`
        your_result=`execute ${i}${suffix}`
        if [[ "$your_result" = "$line" ]]; then
            echo "test ${i}${suffix} correct"
        else
            echo "test ${i}${suffix} wrong answer"
        fi
    fi
    
    i=$(( $i + 1 ))
done < ${result_file_name}