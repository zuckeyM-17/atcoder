def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

q = geti

map = {}

q.times do
  command, name, *args = gets.chomp.split(" ")

  case command
  when '1'
    map[name] = args.first.to_i
  when '2'
    puts map[name]
  end
end