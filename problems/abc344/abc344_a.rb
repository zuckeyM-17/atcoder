def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp

skip = false
s.each_char do |c|
  if c == "|" && !skip
    skip = true
    next
  end

  if c == "|" && skip
    skip = false
    next
  end

  if skip
    next
  end

  print c
end

puts