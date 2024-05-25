def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

a, b, c, d = getis

if a < 0 || c < 0
  a = a + 1000_000_000
  c = c + 1000_000_000
end

if b < 0 || d < 0
  b = b + 1000_000_000
  d = d + 1000_000_000
end

ar, cr = a % 4, c % 4
br, dr = b % 2, d % 2

W = (c - a) / 4
H = (d - b) / 2

# 中の四角形
ans = W * H * 4

# 四角形の左外側
if ar == 1
  ans += H * 2.5
elsif ar == 2
  ans += H * 1
elsif ar == 3
  ans += H * 0.5
end

# 四角形の右外側
if cr == 1
  ans += H * 1.5
elsif cr == 2
  ans += H * 3
elsif cr == 3
  ans += H * 3.5
end

if br == 1
  # 四角形の下外側
  ans += W * 2

  # 四角形の左下
  if ar == 1
    ans += 1.5
  elsif ar == 2
    ans += 0.5
  end

  # 四角形の右下
  if cr == 1
    ans += 0.5
  elsif cr == 2
    ans += 1.5
  elsif cr == 3
    ans += 2
  end
end

if dr == 1
  # 四角形の上外側
  ans += W * 2

  # 四角形の左上
  if ar == 1
    ans += 1
  elsif ar == 2 || ar == 3
    ans += 0.5
  end

  # 四角形の右上
  if cr == 1
    ans += 1
  elsif cr == 2 || cr == 3
    ans += 1.5
  end
end

puts (ans * 2).to_i