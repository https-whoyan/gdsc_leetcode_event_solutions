class Solution {
    public String convertToBase7(int num) {
        if (num == 0) {
            return "0";
        }

        Boolean needAddMinusOne = num < 0;
        if (needAddMinusOne) num = -1 * num;

        StringBuilder base7s = new StringBuilder();
        while (num != 0) {
            base7s.insert(0, Character.toChars((int)'0' + num % 7));
            num /= 7;
        }
        
        if (needAddMinusOne) {
            base7s.insert(0, '-');
        }
        return base7s.toString();
    }
}