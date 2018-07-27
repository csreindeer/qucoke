from qiskit import *
import qpuos
import qpuroot
import qpustandard

#INIT
qr = QuantumRegister(2)
qs = QuantumRegister(2)
qo = QuantumRegister(3)
c = ClassicalRegister(1)
memreg = QuantumRegister(2)
circuit = QuantumCircuit(qo, qr, qs, c)

pswd = "root"; #CHANGE LE PASSWORD

running = True
while running == True:
    qpuos.os()
    
