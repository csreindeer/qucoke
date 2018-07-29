from pyquil import *
import qpuos
import qpuroot
import qpustandard

#INIT
qvm = QVMConnection()
p = Program()

pswd = "root"; #CHANGE LE PASSWORD

running = True
while running == True:
    qpuos.os()
    
