
def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

xa, ya = getis
xb, yb = getis
xc, yc = getis

def dist(x1, y1, x2, y2)
  (x1-x2).pow(2) + (y1-y2).pow(2)
end

ac = dist(xa, ya, xc, yc)
ab = dist(xa, ya, xb, yb)
bc = dist(xb, yb, xc, yc)

if ac == ab + bc || ab == ac + bc || bc == ac + ab
  puts "Yes"
else
  puts "No"
end
