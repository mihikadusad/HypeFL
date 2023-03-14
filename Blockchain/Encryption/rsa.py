# Python for RSA asymmetric cryptographic algorithm.
# For demonstration, values are
# relatively small compared to practical application
import math
import numpy as np

def getRandomPrimeInteger(bounds):

    for i in range(bounds.__len__()-1):
        if bounds[i + 1] > bounds[i]:
            x = bounds[i] + np.random.randint(bounds[i+1]-bounds[i])
            if isPrime(x):
                return x

        else:
            if isPrime(bounds[i]):
                return bounds[i]

        if isPrime(bounds[i + 1]):
            return bounds[i + 1]

    newBounds = [0 for i in range(2*bounds.__len__() - 1)]
    newBounds[0] = bounds[0]
    for i in range(1, bounds.__len__()):
        newBounds[2*i-1] = int((bounds[i-1] + bounds[i])/2)
        newBounds[2*i] = bounds[i]

    return getRandomPrimeInteger(newBounds)

def isPrime(x):
    count = 0
    for i in range(int(x/2)):
        if x % (i+1) == 0:
            count = count+1
    return count == 1
 
 
def gcd(a, h):
    temp = 0
    while(1):
        temp = a % h
        if (temp == 0):
            return h
        a = h
        h = temp
 
bounds = [100, 10000000000000]
p = getRandomPrimeInteger(bounds)
q = getRandomPrimeInteger(bounds)
n = p*q
e = 2
phi = (p-1)*(q-1)
 
while (e < phi):
 
    # e must be coprime to phi and smaller than phi.
    if(gcd(e, phi) == 1):
        break
    else:
        e = e+1
 
# Private key (d stands for decrypt)
# choosing d such that it satisfies
# d*e = 1 + k * totient
 
k = 2
d = (1 + (k*phi))/e
 
# Message to be encrypted
msg = AVD.dataTransmission
 
print("Message data = ", msg)
 
# Encryption c = (msg ^ e) % n
c = pow(msg, e)
c = math.fmod(c, n)
print("Encrypted data = ", c)
 
# Decryption m = (c ^ d) % n
m = pow(c, d)
m = math.fmod(m, n)
print("Original Message Sent = ", m)
