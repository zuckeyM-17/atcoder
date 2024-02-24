S = gets.chomp

h = {}
S.chars.each.with_index do |c, i|
  h[c] == nil ? h[c] = 1 : h[c] += 1
end

puts S.index(h.invert[1])+1