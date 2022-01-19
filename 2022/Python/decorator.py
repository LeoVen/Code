import functools

def example_decorator(func):

    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print('Start of decorator')
        result = func(*args, **kwargs)
        print('End of decorator')

        return result

    return wrapper

@example_decorator
def my_function(name):
    print('Going through the function')
    return f'Hello, {name}!'

# The same as above
# my_function = example_decorator(my_function)

msg = my_function('Isaac')
print(msg)

print(my_function.__name__)
