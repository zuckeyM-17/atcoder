def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

c = []
(1..1_000_000).each do |i|
  num_str = (i*i*i).to_s
  f = true
  num_str.each_char.with_index do |ch, i|
    if ch != num_str[num_str.size - i -1]
      f = false
      break
    end
  end

  c << i * i * i if f
end

c.reverse!
c.each do |v|
  if n >= v
    puts v
    break
  end
end