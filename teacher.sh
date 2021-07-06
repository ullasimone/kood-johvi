interview=$(awk 'FNR==179 {print;exit}' streets/Buckingham_Place | cut -d "#" -f2)
echo $interview
cat */interviews/interview-$interview
grep -A 4 "L337" vehicles | grep -A 3 -B 1 "Honda" | grep -A 2 -B 2 "Blue" | grep -B 4 "Height: 6'*"
cat memberships/AAA memberships/Delta_SkyMiles memberships/Museum_of_Bash_History memberships/Terminal_City_Library | grep -c "$MAIN_SUSPECT"