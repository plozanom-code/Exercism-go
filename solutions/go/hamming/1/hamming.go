package hamming

import "errors"

func Distance(a, b string) (int, error) {
    counter := 0
    
	if len(a) != len(b) {
        return 0, errors.New("Different size of strands")
    }

    for i := 0; i < len(a); i++ {
    	if a[i] != b[i] {
    		counter += 1
    	}
    }
    return counter, nil
}

