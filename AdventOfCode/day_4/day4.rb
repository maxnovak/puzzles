#Input comes from https://adventofcode.com/2018/day/3/input
#copied to text file day3_input.txt for ease of solving problem.

require_relative "GuardEvent"
require 'pp'

events = []

File.open("day4_input.txt", "r") do |file|
	file.each_line do |claim_data|
		data = claim_data.split(" ")
		events.push(GuardEvent.new(data[0].tr("[",""), data[1].tr("]",""), data[2..-1]))
	end
end

events = events.sort_by { |event| [event.date, event.time]  }

pp events