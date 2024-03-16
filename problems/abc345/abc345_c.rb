s = gets.chomp

h = Hash.new(0)

s.each_char do |c|
  h[c] += 1
end

ans = s.length * (s.length - 1) / 2

dup = 0

h.values.each do |v|
  if v > 1
    dup += 1
  end
  if v > 2
    ans -= v * (v - 1) / 2 - 1
  end
end

if dup <= 1
  puts ans
else
  puts ans - (dup -1)
end