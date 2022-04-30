from PIL import Image
from bresenham import bresenham
from math import pi, cos, sin, floor

num_points = 200

def get_line(point_a, point_b):
	i = 0
	if(point_a < point_b):
		pass
	elif(point_a > point_b):
		point_a, point_b = point_b, point_a
	else:
		sys.exit()
	for pair in pairs:
		if pair[0] == point_a and pair[1] == point_b:
			i = pair[2]
	return lines[i]

with Image.open("selfie.jpg") as im:
	pixels = im.load()
	sum = 0
	for x in range(im.width):
		for y in range(im.height):
			sum += pixels[x, y][0]
	print(sum)

	if im.width != im.height:
		sys.exit()

	circle = []
	for i in range(num_points):
		angle = (pi*2/num_points) * i
		x = floor(cos(angle)*im.width/2 +im.width/2)
		y = floor(sin(angle)*im.height/2 +im.height/2)
		circle.append((x, y))

	pairs = []
	lines = []
	for i in range(num_points):
		for n in range(i+1, num_points):
			pairs.append([i, n, len(lines)])
			lines.append(list(bresenham(*circle[i], *circle[n])))
	print(len(pairs))


