from pyquil.quil import Program
from pyquil.api import QVMConnection
from pyquil.gates import CNOT, H

#yes for right now I am using the bell state example on the docs
qvm = QVMConnection()
p = Program(H(0), CNOT(0, 1))

wf = qvm.wavefunction(p)
wfo = str(wf)

#writes to output
o = open ("output.txt", "a")
o.write(wfo)