def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a11, a12, a13 = getis
a21, a22, a23 = getis
a31, a32, a33 = getis

A = [
  [a11, a12, a13],
  [a21, a22, a23],
  [a31, a32, a33],
]

patterns = [
  {t: [[1, 1], [1, 2], [2, 2], [2, 3], [3, 1]], a: [[1, 3], [2, 1], [3, 2], [3, 3]]},
  {t: [[1, 1], [1, 3], [2, 2], [2, 3], [3, 2]], a: [[1, 2], [2, 1], [3, 1], [3, 3]]},
  {t: [[1, 3], [2, 1], [2, 2], [3, 2], [3, 3]], a: [[1, 1], [1, 2], [2, 3], [3, 1]]},
  {t: [[1, 2], [2, 1], [2, 2], [3, 1], [3, 3]], a: [[1, 1], [1, 3], [2, 3], [3, 2]]},
  {t: [[1, 2], [1, 3], [2, 1], [2, 2], [3, 3]], a: [[1, 1], [2, 3], [3, 1], [3, 2]]},
  {t: [[1, 1], [1, 3], [2, 1], [2, 2], [3, 2]], a: [[1, 2], [2, 3], [3, 1], [3, 3]]},
  {t: [[1, 1], [2, 2], [2, 3], [3, 1], [3, 2]], a: [[1, 2], [1, 3], [2, 1], [3, 3]]},
  {t: [[1, 2], [2, 2], [2, 3], [3, 1], [3, 3]], a: [[1, 1], [1, 3], [2, 1], [3, 2]]},
]

patterns.each do |pattern|
  t = pattern[:t]
  a = pattern[:a]

  takahashi = A[t[0][0] - 1][t[0][1] - 1] + A[t[1][0] - 1][t[1][1] - 1] + A[t[2][0] - 1][t[2][1] - 1] + A[t[3][0] - 1][t[3][1] - 1] + A[t[4][0] - 1][t[4][1] - 1]
  aoki = A[a[0][0] - 1][a[0][1] - 1] + A[a[1][0] - 1][a[1][1] - 1] + A[a[2][0] - 1][a[2][1] - 1] + A[a[3][0] - 1][a[3][1] - 1]

  if takahashi < aoki
    puts "Aoki"
    exit
  end
end

puts "Takahashi"