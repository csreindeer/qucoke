from pyquil.quil import Program
from pyquil.api import QVMConnection
from pyquil.gates import CNOT, H, X

# each application will have its own method
#bell stae is gonna be defualt position

#init the computerish?
p = Program()

#what to run in root

def main():
    test()

def test():
    qvm = QVMConnection()
    p = Program(X(1), H(0), CNOT(0, 1))
    wf = qvm.wavefunction(p)
    wfo = str(wf)

    #writes to output
    o = open ("output.txt", "a")
    o.write(wfo)
def test2():
    print("eh")