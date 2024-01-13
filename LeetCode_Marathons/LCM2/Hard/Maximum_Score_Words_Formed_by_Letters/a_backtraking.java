class Solution {
    private List<String> globalWords;
    private List<Integer> globalScoreOfWords;

    private int countAns(int indexWord, Map<Character, Integer> lettersMap) {
        if (indexWord == globalWords.size()) return 0;

        Map<Character, Integer> originalState = new HashMap<>(lettersMap);

        int ansWithoutThisWord = countAns(indexWord + 1, lettersMap);

        for (char letter : globalWords.get(indexWord).toCharArray()) {
            if (!lettersMap.containsKey(letter) || lettersMap.get(letter) == 0) {
                lettersMap.putAll(originalState);
                return ansWithoutThisWord;
            }
            lettersMap.put(letter, lettersMap.get(letter) - 1);
        }

        int ansWithThisWord = countAns(indexWord + 1, lettersMap) + globalScoreOfWords.get(indexWord);
        lettersMap.putAll(originalState);

        return Math.max(ansWithThisWord, ansWithoutThisWord);
    }

    public int maxScoreWords(String[] words, char[] letters, int[] score) {
        globalWords = new ArrayList<>();
        for (String word : words) globalWords.add(word);

        Map<Character, Integer> mapOfLetters = new HashMap<>();
        for (char letter : letters) mapOfLetters.put(letter, mapOfLetters.getOrDefault(letter, 0) + 1);

        globalScoreOfWords = new ArrayList<>();
        for (String word : words) {
            int wordScore = 0;
            for (char letter : word.toCharArray()) wordScore += score[letter - 'a'];
            globalScoreOfWords.add(wordScore);
        }

        return countAns(0, mapOfLetters);
    }
}