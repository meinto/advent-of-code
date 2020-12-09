#!/usr/local/bin/bash

bv=`bash --version | grep version | cut -d' ' -f 4`
if [[ $bv != 5* ]]; then
  echo "Please upgrade to bash version >= 5.x"
  exit
fi

mapfile -t content < input.txt 
lineCount="${#content[@]}"

# exercise 1
for (( i=0; i<$lineCount; i++ )); do
  n1="${content[$i]}"
  for (( j=$((i+1)); j<$lineCount; j++ )); do
    n2="${content[$j]}"
    if test $(($n1+$n2)) -eq 2020; then
      echo "$n1+$n2=$(($n1+$n2))"
      echo "$n1*$n2=$(($n1*$n2))"
    fi
  done
done

# exercise 2
for (( i=0; i<$lineCount; i++ )); do
  n1="${content[$i]}"
  for (( j=$((i+1)); j<$lineCount; j++ )); do
    n2="${content[$j]}"
    for (( k=$((j+1)); k<$lineCount; k++ )); do
      n3="${content[$k]}"
      if test $(($n1+$n2+$n3)) -eq 2020; then
        echo "$n1+$n2+$n3=$(($n1+$n2+$n3))"
        echo "$n1*$n2*$n3=$(($n1*$n2*$n3))"
      fi
    done
  done
done