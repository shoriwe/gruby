yield 1

def my_function()
    yield 2
end

gen my_generator()
    yield None
end

for a in range(100)
    yield 3
end

for value in range(2000)
    gen __anonymous()
        yield value
    end
end

if False
    yield 4
end