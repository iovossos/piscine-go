interview_num=$(head -n 179 streets/"Buckingham_Place" | tail -n 1 | grep -o '[0-9]\+')
export interview_num

echo $interview_num

cat interviews/interview-$interview_num

echo "$MAIN_SUSPECT"