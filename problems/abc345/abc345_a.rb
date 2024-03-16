def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

s = gets.chomp

s.each_char.with_index do |c, idx|
  if idx == 0 && c != '<'
    puts 'No'
    exit
  end

  if idx == s.length - 1 && c != '>'
    puts 'No'
    exit
  end

  next if idx == 0 || idx == s.length - 1

  if c != '='
    puts 'No'
    exit
  end
end

puts 'Yes'