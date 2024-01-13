class Solution:
    def countTime(self, time: str) -> int:
        myHour, myMinutes = time.split(":")
        myHourAns, myMinutesAns = 1, 1

        if myHour[1] == "?":
            if myHour[0] == "?": myHourAns = 24
            elif myHour[0] == "2": myHourAns = 4
            else: myHourAns = 10
        else:
            if myHour[0] == "?": 
                myHourAns = 2
                if myHour[1] <= "3": myHourAns += 1
        
        if myMinutes[0] == "?": myMinutesAns *= 6
        if myMinutes[1] == "?": myMinutesAns *= 10

        return myMinutesAns * myHourAns
        