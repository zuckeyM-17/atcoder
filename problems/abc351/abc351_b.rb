def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

a = []
b = []

n.times do
  a << gets.chomp
end

n.times do
  b << gets.chomp
end

a.each_with_index do |ai, i|
  ai.each_char.each_with_index do |aai, j|
    if aai != b[i][j]
      puts "#{i + 1} #{j + 1}"
      exit
    end
  end
end