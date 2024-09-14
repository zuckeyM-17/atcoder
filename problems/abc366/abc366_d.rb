def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

a = []
n.times do |i|
  a << []
  n.times do |j|
    a[i] << getis
  end
end

# a.each do |ai|
#   p ai
# end

b = []
n.times do |i|
  b << {cols: [], rows: [], z: []}
  n.times do
    b[i][:cols] << []
    b[i][:rows] << []
    b[i][:z] << []
  end
  a[i].each_with_index do |ai, j|
    ai.each_with_index do |aij, k|
      if k == 0
        b[i][:rows][j] << aij
      else
        b[i][:rows][j] << b[i][:rows][j][k-1] + aij
      end

      if j == 0
        b[i][:cols][k] << aij
      else
        b[i][:cols][k] << b[i][:cols][k][j-1] + aij
      end
    end
  end
  b[i][:rows].each_with_index do |bi, j|
    b[i][:z] << []
    bi.each_with_index do |bij, k|
      if j == 0
        b[i][:z][j] << bij
      else
        b[i][:z][j] << b[i][:z][j-1][k] + bij
      end
    end
  end
end

q = geti

a.each do |ai|
  ai.each do |aij|
    p aij
  end
  puts "----"
end

q.times do
  lx, rx, ly, ry, lz, rz = getis
  # puts "lx: #{lx}, rx: #{rx}, ly: #{ly}, ry: #{ry}, lz: #{lz}, rz: #{rz}"
  ans = 0
  (lx..rx).each do |i|
    if ly < ry && lz < rz
      ans += b[i-1][:z][ry-1][rz-1] - b[i-1][:z][ly-1][rz-1] - b[i-1][:z][ry-1][lz-1] + b[i-1][:z][ly-1][lz-1]
    elsif ly < ry
      ans += b[i-1][:cols][ry-1][rz-1] - b[i-1][:cols][ry-1][lz-1]
    elsif lz < rz
      ans += b[i-1][:rows][ry-1][rz-1] - b[i-1][:rows][ly-1][rz-1]
    else
      ans += a[i-1][ry-1][rz-1]
    end
  end
  puts ans
end