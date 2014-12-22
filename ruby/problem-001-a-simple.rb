ceiling = 1000

sum_of_series =
  (1...ceiling).select do |i|
    i % 3 == 0 || i % 5 == 0
  end.inject(0) do |sum, i|
    sum += i
  end

puts "The sum of multiples of 3 or 5 up to #{ceiling} is #{sum_of_series}"

