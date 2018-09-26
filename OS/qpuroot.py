from pyquil.quil import Program
from pyquil.api import QVMConnection
from pyquil.gates import CNOT, H, X

p = Program()

def main():
    test()

def test():
    qvm = QVMConnection()
    p = Program(X(1), H(0), CNOT(0, 1))
    wf = qvm.wavefunction(p)
    wfo = str(wf)
    output.writetooutput()