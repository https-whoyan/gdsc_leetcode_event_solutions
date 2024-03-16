class Solution {
    private boolean isMonotic(String word, boolean isUpper) {
        for (Character letter : word.toCharArray()) {
            if (isUpper && !('A' <= letter && letter <= 'Z')) {
                return false;
            }
            if (!isUpper && !('a' <= letter && letter <= 'z')) {
                return false;
            }
        }

        return true;
    }

    public boolean detectCapitalUse(String word) {
        if ('a' <= word.charAt(0) && word.charAt(0) <= 'z') {
            return isMonotic(word.substring(1), false);
        }
        return isMonotic(word.substring(1), false) || isMonotic(word.substring(1), true);
    }
}