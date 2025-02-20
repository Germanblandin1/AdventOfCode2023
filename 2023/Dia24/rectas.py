import matplotlib.pyplot as plt
from mpl_toolkits.mplot3d import Axes3D


def read_lines_from_file(filename):
    with open(filename, 'r') as file:
        lines = []
        for line in file.readlines():
            parts = line.split()
            if len(parts) != 6:
                raise ValueError(f"Invalid line: {line}")
            
            x1, y1, z1, x2, y2, z2 = map(float, parts)
            lines.append([(x1, y1, z1), (x2, y2, z2)])
        return lines

# Uso de la función
lines = read_lines_from_file('salida.out')


fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')

for line in lines:
    xs, ys, zs = zip(*line)  # Desempaquetar los puntos de la línea
    ax.plot(xs, ys, zs)

plt.show()