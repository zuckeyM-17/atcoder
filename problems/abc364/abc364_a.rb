def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

flag = false
n.times do |i|
  if i + 1 == n
    break
  end
  cur = gets.chomp
  if flag == true
    if cur == "sweet"
      puts "No"
      exit
    else
      flag = false
    end
  else
    if cur == "sweet"
      flag = true
    end
  end
end

puts "Yes"
