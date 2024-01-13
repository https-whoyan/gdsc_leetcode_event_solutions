class Solution {
    public int countTime(String time) {
        String myHour = time.substring(0, 2), myMinutes = time.substring(3, 5);
        int myHourAns = 1, myMinutesAns = 1;

        if (myHour.charAt(1) == '?') {
            if (myHour.charAt(0) == '?') myHourAns = 24;
            else if (myHour.charAt(0) == '2') myHourAns = 4;
            else myHourAns = 10;
        } else {
            if (myHour.charAt(0) == '?') {
                myHourAns = 2;
                if (myHour.charAt(1) <= '3') myHourAns += 1;
            }
        }

        if (myMinutes.charAt(0) == '?') myMinutesAns *= 6;
        if (myMinutes.charAt(1) == '?') myMinutesAns *= 10;

        return myHourAns * myMinutesAns;
    }
}