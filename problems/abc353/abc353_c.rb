def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti
a = getis

sum_a = a.sum % 100_000_000  # 全体の和を先に計算し、モジュロを適用
ans = 0
n.times do |i|
  sum_a = (sum_a - a[i] + 100_000_000) % 100_000_000  # 現在の要素を全体の和から引く（モジュロを適用して負の数にならないようにする）
  ans = (ans + a[i] * sum_a) % 100_000_000  # 新たなペアの和を計算し、モジュロを適用
end

puts ans