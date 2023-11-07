# lab2-mv

Пример вывода программы:

Matrix A:
 10.000000   0.000812   0.000719   0.000660   0.000617   0.000584   0.000558   0.000536  ... |      30.041880
  0.001000  28.284271   0.000719   0.000660   0.000617   0.000584   0.000558   0.000536  ... |     113.178716
  0.001000   0.000812  51.961524   0.000660   0.000617   0.000584   0.000558   0.000536  ... |     259.848905
  0.001000   0.000812   0.000719  80.000000   0.000617   0.000584   0.000558   0.000536  ... |     480.040921
  0.001000   0.000812   0.000719   0.000660 111.803399   0.000584   0.000558   0.000536  ... |     782.664353
  0.001000   0.000812   0.000719   0.000660   0.000617 146.969385   0.000558   0.000536  ... |     1175.795283
  0.001000   0.000812   0.000719   0.000660   0.000617   0.000584 185.202592   0.000536  ... |     1666.863186
  0.001000   0.000812   0.000719   0.000660   0.000617   0.000584   0.000558 226.274170  ... |     2262.781221
    ...        ...        ...        ...        ...        ...        ...        ...     ... |         ...
The matrix is strictly diagonally dominant.


Jacobi method: the required accuracy per iteration has been achieved on iter 3
Gauss-Seidel method: the maximum number of iterations reached (3000)
Relaxation method (w = 0.50): the required accuracy per iteration has been achieved on iter 24
Relaxation method (w = 1.50): the required accuracy per iteration has been achieved on iter 25


Precise solution (x*):  30.041880 113.178716 259.848905 480.040921 782.664353 1175.795283 1666.863186 2262.781221 2970.039190 3794.772058
Approximate solution by the Jacobi method:   3.000000   4.000000   5.000000   6.000000   7.000000   8.000000   9.000000  10.000000  11.000000  12.000000
Approximate solution by the Gauss-Seidel method:   3.000000   4.000000   5.000000   6.000000   7.000000   8.000000   9.000000  10.000000  11.000000  12.000000
Approximate solution by the relaxation method (w = 0.5):   3.046293   4.043231   5.042198   6.041552   7.041045   8.040603   9.040200  10.039822  11.039462  12.039116
Approximate solution by the relaxation method (w = 1.5):   3.046293   4.043232   5.042198   6.041553   7.041045   8.040604   9.040201  10.039823  11.039463  12.039117


Jacobi method: 0s
Gauss-Seidel method: 1.5384ms


Cubic norm of difference (Jacobi method): 4613.748723
Relative error (Jacobi method):   0.996838
Cubic norm of difference (Gauss-Seidel method): 4613.748723
Relative error (Gauss-Seidel method):   0.996838
Cubic norm of difference (relaxation method, w = 0.5): 4613.687157
Relative error (relaxation method, w = 0.5):   0.996827
Cubic norm of difference (relaxation method, w = 1.5): 4613.687156
Relative error (relaxation method, w = 1.5):   0.996827
