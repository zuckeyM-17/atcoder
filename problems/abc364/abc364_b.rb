def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

h, w = getis
si, sj = getis
si -= 1
sj -= 1
c = []
h.times do
  c << gets.chomp.split('')
end
x = gets.chomp.split('')

x.each do |char|
  if char == 'U'
    if si > 0 && c[si - 1][sj] != '#'
      si -= 1
    end
  elsif char == 'D'
    if si < h - 1 && c[si + 1][sj] != '#'
      si += 1
    end
  elsif char == 'L'
    if sj > 0 && c[si][sj - 1] != '#'
      sj -= 1
    end
  elsif char == 'R'
    if sj < w - 1 && c[si][sj + 1] != '#'
      sj += 1
    end
  end
end

puts "#{si + 1} #{sj + 1}"