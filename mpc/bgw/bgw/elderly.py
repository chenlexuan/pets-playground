from mpyc.runtime import mpc

async def main():
    secint = mpc.SecInt(16)

    await mpc.start()

    my_age = int(input('Enter your age: '))
    our_ages = mpc.input(secint(my_age))

    total_age = sum(our_ages)
    max_age = mpc.max(our_ages)
    m = len(mpc.parties)
    above_avg = mpc.sum(age * m > total_age for age in our_ages)

    print('Average age:', await mpc.output(total_age) / m)
    print('Maximum age:', await mpc.output(max_age))
    print('Number of "elderly":', await mpc.output(above_avg))

    await mpc.shutdown()

mpc.run(main())