import java.util.*;

public class Solution {
    public String getMask(int number, int maskSize) {
        StringBuilder mask = new StringBuilder();
        while (number != 0) {
            mask.insert(0, (char) ('0' + (number % 2)));
            number /= 2;
        }
        int needToAppend = maskSize - mask.length();
        for (int jj = 0; jj < needToAppend; jj++) {
            mask.insert(0, '0');
        }
        return mask.toString();
    }

    public int maxScoreWords(String[] words, char[] letters, int[] score) {
        int n = words.length;
        int maskCount = (int) Math.pow(2, n);

        int ans = 0, scoreCounter = 0;
        Map<Character, Integer> mpLetters = new HashMap<>();

        for (char myChr : letters) {
            mpLetters.put(myChr, mpLetters.getOrDefault(myChr, 0) + 1);
        }

        for (int maskInt = 1; maskInt < maskCount; maskInt++) {
            String mask = getMask(maskInt, n);
            Map<Character, Integer> mpLettersOfMask = new HashMap<>();

            for (int indexOfMask = 0; indexOfMask < n; indexOfMask++) {
                if (mask.charAt(indexOfMask) == '1') {
                    for (char myChr : words[indexOfMask].toCharArray()) {
                        mpLettersOfMask.put(myChr, mpLettersOfMask.getOrDefault(myChr, 0) + 1);
                    }
                }
            }

            scoreCounter = 0;
            for (Map.Entry<Character, Integer> entry : mpLettersOfMask.entrySet()) {
                char key = entry.getKey();
                int value = entry.getValue();
                int value2 = mpLetters.getOrDefault(key, 0);
                if (value > value2) {
                    scoreCounter = 0;
                    break;
                } else {
                    scoreCounter += score[key - 'a'] * value;
                }
            }

            if (scoreCounter > ans) {
                ans = scoreCounter;
            }
            mpLettersOfMask.clear();
        }

        return ans;
    }
}