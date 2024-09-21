
def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n, m = getis

h = Hash.new(0)

m.times do
  a, b = gets.chomp.split(" ")
  a = a.to_i

  if b == 'F'
    puts 'No'
  else
    if h[a] == 0
      puts 'Yes'
      h[a] = 1
    else
      puts 'No'
    end
  end
end