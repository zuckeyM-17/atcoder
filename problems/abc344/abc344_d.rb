def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

ss = []
t = gets.chomp
n = geti

n.times do |i|
  a, *sa = gets.chomp.split(" ")
  sa << ""
  ss << sa
end

c = {
  "" => 0
}
ss.each do |s|
  h = Hash.new(0)
  s.each do |a|
    c.keys.each do |b|
      if b+a == t[0..(b+a).size-1]
        plus = a == "" ? 0 : 1
        if h[b+a] != 0
          h[b+a] = [h[b+a], c[b]+plus].min
        else
          h[b+a] = c[b]+plus
        end
      end
    end
  end
  c = h
end

if c[t] != 0
  puts c[t]
else
  puts -1
end