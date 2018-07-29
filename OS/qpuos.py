from pyquil import *
import init
import qpuroot
import qpustandard
import interface

def os():
    runroot = True
    runstand = True
    while runroot == True:
        qpuroot.root()

    while runstand == True:
        qpustandard.standard()
