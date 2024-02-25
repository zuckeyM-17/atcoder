s = gets.chomp


stack = []
s.chars.each.with_index(1) do |c, i|
  if c == "("
    stack.push(i)
  else
    puts "#{stack.pop} #{i}"
  end
end