import sys
import requests
from bs4 import BeautifulSoup

fn = open('output.txt', 'w')
class ShortDesc:

	def __init__(self):
		
		#Gets webite and parses 
		page = requests.get(\
			"https://forecast.weather.gov/MapClick.php?lat=28.601870000000076&lon=-81.19759999999997#.W51fAN1KhhE")
		soup = BeautifulSoup(page.content, 'html.parser')
		self.seven_day = soup.find_all(class_="forecast-tombstone")

		#Replaces breaks with spaces
		for br in soup.find_all("br"):
			br.replace_with(" ")

	#Enter high or low 	
	def today(self):
		#Gets description
		desc = self.seven_day[0].find(class_="short-desc").getText()

		#Checks morning or night
		high_temp = self.seven_day[0].find(class_="temp temp-high")	
		if(high_temp == None):
			#Gets all information and writes to output
			low_temp = self.seven_day[0].find(class_="temp temp-low").getText()
			fn.write('Tonight\n')
			fn.write(desc + '\n')
			fn.write(low_temp + '\n\n')	
		else:
			#Gets all information and writes to output
			low_temp = self.seven_day[1].find(class_="temp temp-low").getText()
			fn.write('Today\n')
			fn.write(desc + '\n')	
			fn.write(high_temp.getText() + '\n')
			fn.write(low_temp + '\n\n')

	def tomorrow(self):
		high_temp = self.seven_day[1].find(class_='temp temp-high')

		#Checks if it's morning or night
		if high_temp == None:
			#Gets all information 
			desc = self.seven_day[2].find(class_='short-desc').getText()
			high_temp = self.seven_day[2].find(class_='temp temp-high').getText()
			low_temp = self.seven_day[3].find(class_='temp temp-low').getText(	)
			
			#Writes to file
			fn.write('Tomorrow\n')
			fn.write(desc + '\n')
			fn.write(high_temp + '\n')
			fn.write(low_temp + '\n\n')
		else:
			#Gets all information
			desc = self.seven_day[1].find(class_='short-desc').getText()
			low_temp = self.seven_day[2].find(class_='temp temp-low').getText()
			
			#Write to files
			fn.write('Tomorrow\n')
			fn.write(desc + '\n')
			fn.write(high_temp.getText() + '\n')
			fn.write(low_temp + '\n\n')
			
	def four_day(self):

		high_temp = self.seven_day[0].find(class_='temp temp-high')

		#Checks morning or night
		if high_temp == None:
			#Gets information
			tonight_desc = self.seven_day[0].find(class_='short-desc').getText()
			low_temp = self.seven_day[0].find(class_='temp temp-low').getText()
			
			#Writes to file
			fn.write('Tomorrow\n')
			fn.write(tonight_desc + '\n')
			fn.write(low_temp + '\n\n')

			days = [1,3,5,7]
			for i in days:
				#Gets information for each day
				period = self.seven_day[i].find(class_='period-name').getText()
				desc = self.seven_day[i].find(class_='short-desc').getText()
				high_temp = self.seven_day[i].find(class_='temp temp-high').getText()
				low_temp = self.seven_day[i + 1].find(class_='temp temp-low').getText()
				
				#Writes to file
				fn.write(period + '\n')
				fn.write(desc + '\n')
				fn.write(high_temp + '\n')
				fn.write(low_temp + '\n\n')

		else:
			days = [0,2,4,6]
			for i in days:
				#Gets information for each day
				period = self.seven_day[i].find(class_='period-name').getText()
				desc = self.seven_day[i].find(class_='short-desc').getText()
				high_temp = self.seven_day[i].find(class_='temp temp-high').getText()
				low_temp = self.seven_day[i + 1].find(class_='temp temp-low').getText()
				
				#Writes to file
				fn.write(period + '\n')
				fn.write(desc + '\n')
				fn.write(high_temp + '\n')
				fn.write(low_temp + '\n\n')


_string = sys.argv[1]
weather = ShortDesc()
if _string == 'today':
	weather.today()
elif _string == 'tomorrow':
	weather.tomorrow()
elif _string == 'week':
	weather.four_day()
else:
	fn.write('Invalid argument.')				
#Clean up
fn.close()