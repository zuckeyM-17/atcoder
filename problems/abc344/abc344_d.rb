def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

ss = []
t = gets.chomp
tl = t.size
n = geti

SUPER_BIG = 10**9 + 7

dp = []
(n+1).times do |i|
  dp[i] = []
  (tl + 1).times do |j|
    dp[i][j] = SUPER_BIG
  end
end

dp[0][0] = 0

n.times do |i|
  (tl + 1).times { |j| dp[i+1][j] = dp[i][j] }

  a, *ss = gets.chomp.split(" ")
  ss.each do |s|
    sl = s.size
    tl.times do |j|
      break if j + sl > tl

      if t[j, sl] == s
        dp[i+1][j+sl] = [dp[i+1][j+sl], dp[i][j] + 1].min
      end
    end
  end
end

if dp[n][tl] == SUPER_BIG
  dp[n][tl] = -1
end

puts dp[n][tl]