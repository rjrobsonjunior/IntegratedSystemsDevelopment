import numpy as np

# create two random matrices
A = np.random.rand(1000, 1000)
B = np.random.rand(1000, 1000)

# multiply them
C = A @ B

print("Result shape:", C.shape)
