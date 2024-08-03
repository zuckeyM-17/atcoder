def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
s = gets.chomp.split("")

stone = 'R'
paper = 'P'
scissors = 'S'

def maximum(a, b)
  a > b ? a : b
end

win_hands = { stone => paper, paper => scissors, scissors => stone }
lose_hands = { stone => scissors, paper => stone, scissors => paper }
hands = [stone, paper, scissors]

dp = Array.new(n) { Array.new(3, 0) }

hands.each do |hand|
  if win_hands[s[0]] == hand
    dp[0][hands.index(hand)] = 1
  elsif lose_hands[s[0]] == hand
    dp[0][hands.index(hand)] = -1
  end
end

(1...n).each do |i|
  hands.each do |hand|
    if lose_hands[s[i]] == hand
      dp[i][hands.index(hand)] = -1
      next
    end

    others = hands - [hand]
    previous_others_max = maximum(dp[i-1][hands.index(others[0])], dp[i-1][hands.index(others[1])])
    if win_hands[s[i]] == hand
      dp[i][hands.index(hand)] = previous_others_max + 1
    else
      dp[i][hands.index(hand)] = previous_others_max
    end
  end
end

puts dp[n-1].max
