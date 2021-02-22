#!/usr/bin/python
# -*- coding:utf-8 -*-
import time
import turtle
def round_rectangle(length,high,cor_angle,cor_rad):
	for i in range(2):
		turtle.fd(high)
		turtle.circle(cor_rad,cor_angle)
		turtle.fd(length)
		turtle.circle(cor_rad,cor_angle)
		
class iPhone:
	def __init__(self):
		self.name='iPhone'
		self.brand='Apple'
		self.system='iOS'
	def show(self):
    		print(self.name+'\n'+self.brand+'\n'+self.system)
		
class iPhone_bezel_less(iPhone):
	def __init__(self,a):
		super(iPhone_bezel_less,self).__init__()
		self.name+=' '+a['name']
		self.screensize=a['screensize']
		self.color=a['color']
		self.year=a['year']
		self.processor=a['processor']
	def show(self):
	#画布
		turtle.setup(500,800)
		pythonsize=2
		turtle.pensize(pythonsize)#画笔宽度
		turtle.speed(20)
		turtle.seth(90)#启动时运行角度
	#最外边框
		turtle.pencolor("black")
		turtle.penup()
		turtle.goto(151,-231)
		turtle.pendown()
		round_rectangle(242,532,90,30)
	#填充
		turtle.penup()
		turtle.goto(150,-230)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("black")
		round_rectangle(240,530,90,30)
		turtle.end_fill()
	#屏幕
		turtle.pencolor("black")
		turtle.penup()
		turtle.goto(140,-225)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("#ffffff")
		round_rectangle(230,520,90,25)
		turtle.end_fill()
	#听筒
		turtle.penup()
		turtle.goto(75,310)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("black")
		round_rectangle(130,10,90,10)
		turtle.end_fill()
		turtle.hideturtle()
	#输出文字
		printer = turtle.Turtle()
		printer.hideturtle()
		printer.goto(0,0)
		printer.penup()
		printer.color("black")
		printer.write(self.name+"\n\n\n\n\n",align="center", font=("Consolas",16,"bold"))
		printer.write(self.brand+'\n\n\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write(self.system+'\n\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write('in '+str(self.year)+'\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write(str(self.screensize)+' inches Display\n', align="center", font=("Consolas", 16, "bold"))
		printer.write(self.processor+' Processor',align="center", font=("Consolas",16,"bold"))
		printer.goto(0,-20)
		printer.write(self.color,align="center", font=("Consolas",16,"bold"))
		time.sleep(10)

		
class iPhone_classic(iPhone):
	
	def __init__(self,a):
		super(iPhone_classic,self).__init__()
		self.name+=' '+a['name']
		self.screensize=a['screensize']
		self.color=a['color']
		self.year=a['year']
		self.processor=a['processor']
		
	def show(self):
	#画布
		turtle.setup(500,800)
		pythonsize=2
		turtle.pensize(pythonsize)#画笔宽度
		turtle.speed(20)
		turtle.seth(90)#启动时运行角度
	#最外边框
		turtle.pencolor("#8E8e8e")
		turtle.penup()
		turtle.goto(152,-202)
		turtle.pendown()
		round_rectangle(244,484,90,30)
	#填充
		turtle.penup()
		turtle.goto(150,-200)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("#ffffff")
		round_rectangle(240,480,90,30)
		turtle.end_fill()
	#手机屏
		turtle.pencolor("black")
		turtle.penup()
		turtle.goto(135,-150)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("black")
		round_rectangle(270,380,90,0)
		turtle.end_fill()
	#听筒
		turtle.penup()
		turtle.goto(30,265)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("#474747")
		round_rectangle(60,4,90,1)
		turtle.end_fill()
	#感光器
		turtle.penup()
		turtle.goto(5,290)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("#3c3c3c")
		turtle.circle(4,360)
		turtle.end_fill()
	#摄像头
		turtle.penup()
		turtle.goto(-50,265)
		turtle.pendown()
		turtle.begin_fill()
		turtle.color("#3c3c3c")
		turtle.circle(6,360)
		turtle.end_fill()
	#Home健
		turtle.pencolor("#9d9d9d")#（#3000440） 
		turtle.penup()
		turtle.goto(25,-185)
		turtle.pendown()
		turtle.circle(25,360)
	#Home方框
		turtle.pencolor("#9d9d9d")#（#3000440） 
		turtle.penup()
		turtle.goto(10,-190)
		turtle.pendown()
		round_rectangle(9, 9, 90, 5)
		turtle.hideturtle()
	#输出文字
		printer = turtle.Turtle()
		printer.hideturtle()
		printer.goto(0,0)
		printer.penup()
		printer.color("#ffffff")
		printer.write(self.name+"\n\n\n\n\n",align="center", font=("Consolas",16,"bold"))
		printer.write(self.brand+'\n\n\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write(self.system+'\n\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write('in '+str(self.year)+'\n\n',align="center", font=("Consolas",16,"bold"))
		printer.write(str(self.screensize)+' inches Display\n', align="center", font=("Consolas", 16, "bold"))
		printer.write(self.processor+' Processor',align="center", font=("Consolas",16,"bold"))
		printer.goto(0,-20)
		printer.write(self.color,align="center", font=("Consolas",16,"bold"))
		time.sleep(10)


iPhone = {'name':'','screensize':3.5,'color':'Silver','year':2007,'processor':'A1'}
iPhone3g = {'name':'3G','screensize':3.5,'color':'Black','year':2008,'processor':'A2'}
iPhone3gs = {'name':'3GS','screensize':3.5,'color':'Black','year':2009,'processor':'A3'}
iPhone4 = {'name':'4','screensize':3.5,'color':'Black','year':2010,'processor':'A4'}
iPhone4s = {'name':'4S','screensize':3.5,'color':'White','year':2011,'processor':'A5'}
iPhone5 = {'name':'5','screensize':4,'color':'Silver','year':2012,'processor':'A6'}
iPhone5s = {'name':'5S','screensize':4,'color':'Gold','year':2013,'processor':'A7'}
iPhone5c = {'name':'5C','screensize':4,'color':'Blue','year':2013,'processor':'A6'}
iPhone6 = {'name':'6','screensize':4.7,'color':'Space Grey','year':2014,'processor':'A8'}
iPhone6p = {'name':'6 Plus','screensize':5.5,'color':'Gold','year':2014,'processor':'A8'}
iPhone6s = {'name':'6S','screensize':4.7,'color':'Rose Gold','year':2015,'processor':'A9'}
iPhone6sp = {'name':'6S Plus','screensize':5.5,'color':'Rose Gold','year':2015,'processor':'A9'}
iPhonese = {'name':'SE','screensize':4,'color':'Pink Gold','year':2016,'processor':'A9'}
iPhone7 = {'name':'7','screensize':4.7,'color':'Bright Black','year':2016,'processor':'A10'}
iPhone7p = {'name':'7 Plus','screensize':5.5,'color':'Bright Black','year':2016,'processor':'A10'}
iPhone8 = {'name':'8','screensize':4.7,'color':'Glass Black','year':2017,'processor':'A11'}
iPhone8p = {'name':'8 Plus','screensize':5.5,'color':'Glass Black','year':2017,'processor':'A11'}
iPhonex = {'name':'X','screensize':5.8,'color':'White','year':2017,'processor':'A11'}
iPhonexr = {'name':'Xr','screensize':6.1,'color':'Coral','year':2018,'processor':'A12'}
iPhonexs = {'name':'XS','screensize':5.8,'color':'Gold','year':2018,'processor':'A12'}
iPhonexsm = {'name':'XS Max','screensize':6.5,'color':'Gold','year':2018,'processor':'A12'}

L={'iPhone':iPhone,'3g':iPhone3g,'3gs':iPhone3gs,'4':iPhone4,'4s':iPhone4s,'5':iPhone5,'5s':iPhone5s,'5c':iPhone5c,'6':iPhone6,'6p':iPhone6p,'se':iPhonese,'7':iPhone7,'7p':iPhone7p,'8':iPhone8p,'x':iPhonex,'xr':iPhonexr,'xs':iPhonexs,'xsm':iPhonexsm,'6s':iPhone6s,'6sp':iPhone6sp}
i=input("Please input the iPhone's name:")
if i=='x' or i=='xr' or i=='xs' or i=='xsm':
	a=iPhone_bezel_less(L[i])
	a.show()
else:
	a=iPhone_classic(L[i])
	a.show()