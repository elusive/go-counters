package main

import (
    "sync"
    "testing"
)


func TestNotSafeCounterCorrectness(t *testing.T) {
    // arrange
    var counters [200]Counter
    for i := 0; i < 200; i++ {
        counters[i] = NewNotSafeCounter() 
    }

    // act
    correct := 0
    incorrect := 0
    for _, ctr := range counters {
        isCorrect := testCorrectness(t, ctr)    
        if isCorrect {
            correct++ 
        } else {
            incorrect++
        }
    }

    t.Logf("Correct = %d, Incorrect = %d", correct, incorrect)
    if incorrect > 1 {
        t.Errorf("Incorrect count should be 1 or less, but was %d", incorrect)
    }
}



func testCorrectness(t *testing.T, counter Counter) bool {
    // arrange
    wg := &sync.WaitGroup{}

    // act
    for i := 0; i < 100; i++ {
        wg.Add(1)
        // 2/3 adds one to counter 1/3 does not
        if i % 3 == 0 {
            go func(counter Counter) {
                counter.Read()
                wg.Done()
            }(counter)
        } else if i % 3 == 1 {
            go func(counter Counter) {
                counter.Add(1)
                wg.Done()
            }(counter)
        } else {
            go func(counter Counter) {
                counter.Add(1)
                counter.Read()
                wg.Done()
            }(counter)
        }
    }

    wg.Wait()

    if counter.Read() != 66 {
        t.Errorf("Counter should be %d but is")
        return false
    } else {
        return true
    }
}

