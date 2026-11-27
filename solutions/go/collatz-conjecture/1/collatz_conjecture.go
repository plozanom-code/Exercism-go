package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("Zero or negative value")
    }

    counter := 0
    if n == 1 {
        return 0, nil
    }
    
    for n != 1 {
        if n % 2 == 0 {
            n = n / 2
        } else{
            n = 3 * n + 1
        }

        counter ++
    }
    
    return counter, nil
}
