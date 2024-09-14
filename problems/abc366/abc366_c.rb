require 'set'

def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

class MySet
  def initialize
    @hash = Hash.new(0)
    @set = Set.new
  end

  def add(el)
    @hash[el] += 1
    @set.add(el)
  end

  def remove(el)
    @hash[el] -= 1
    @set.delete(el) if @hash[el] == 0
  end

  def size
    @set.size
  end
end


q = geti

s = MySet.new

q.times do
  query = getis
  if query[0] == 1
    s.add(query[1])
  elsif query[0] == 2
    s.remove(query[1])
  else
    puts s.size
  end
end
