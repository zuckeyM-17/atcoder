def geti
  gets.chomp.to_i
end

def getis
  gets.chomp.split(" ").map(&:to_i)
end

n = geti

def calculate_vn_mod(n)
  mod = 998244353
  n_str = n.to_s
  n_len = n_str.length

  power_of_10 = 10.pow(n_len, mod)

  total_power_of_10 = power_of_10.pow(n, mod)

  inv = mod_inv(power_of_10 - 1, mod)

  s = (total_power_of_10 - 1 + mod) % mod
  s = (s * inv) % mod

  vn = (n * s) % mod

  return vn
end

def mod_inv(a, m)
  m0, x0, x1 = m, 0, 1
  while a > 1
    q = a / m
    m, a = a % m, m
    x0, x1 = x1 - q * x0, x0
  end
  x1 += m0 if x1 < 0
  x1
end

puts calculate_vn_mod(n)