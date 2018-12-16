#Input comes from https://adventofcode.com/2018/day/3/input
#copied to text file day3_input.txt for ease of solving problem.

require_relative "GuardEvent"
require 'pp'

events = []
guardSleepTime = {}
guardSleepTime.default = 0
guardStartSleep = {}
guardStartSleep.default = []

File.open("day4_input.txt", "r") do |file|
	file.each_line do |claim_data|
		data = claim_data.split(" ")
		events.push(GuardEvent.new(data[0].tr("[",""), data[1].tr("]",""), data[2..-1]))
	end
end

events = events.sort_by { |event| [event.date, event.time]  }

currentGuardId = ""
startTime = 0

#assuming that they only fallasleep and wakeup at 12-1 am
events.each do |event|
	if /#/.match(event.event[1])
		currentGuardId = event.event[1]
	elsif /falls/.match(event.event[0])
		startTime = event.time.split(":")[1].to_i
		guardStartSleep[currentGuardId] = guardStartSleep[currentGuardId].push(startTime)
	elsif /wakes/.match(event.event[0])
		endTime = event.time.split(":")[1].to_i
		sleepTime = endTime - startTime
		guardSleepTime[currentGuardId] = guardSleepTime[currentGuardId] + sleepTime
	end
end

p guardSleepTime.sort_by{|_key, value| value}
p guardStartSleep