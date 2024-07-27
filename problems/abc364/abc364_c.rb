def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end
n, x, y = getis
a = getis
b = getis

dishes = (0..n-1)
  .map { |i| [a[i], b[i]] }
  .sort_by { |a, b| [a, b] }

suma, sumb, cnt = 0, 0, 0

dishes.each do |a, b|
  break if suma + a > x || sumb + b > y

  suma += a
  sumb += b
  cnt += 1
end

puts cnt
