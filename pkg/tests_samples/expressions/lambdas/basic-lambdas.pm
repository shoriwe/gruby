(lambda x, y: (x**2+y**2)**(1/2))(1, 1)
a = (lambda x, y: (x**2+y**2)**(1/2))
a(1, 1)
a(1, 1)
(lambda x, y: (x**2+y**2)**(1/2))(1, 1)
c = [1, 2, 3, (lambda x, y: (x**2+y**2)**(1/2))]
c[3](1, 1)
c[3](1, 1)
a = (lambda x, y: (x**2+y**2)**(1/2))
a.b = 1
a.c = 1
a(a.b, a.c)
(lambda: 1)()