
def main():
    def adder(x):
        def add(y):
            x = x + 1
            return x + y
        return add
    
    ret = adder(100)
    print(ret(200))


if __name__ == "__main__":
    main()
